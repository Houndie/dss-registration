package registration

import (
	"context"
	"errors"
	"fmt"
	"strconv"

	"github.com/Houndie/dss-registration/dynamic/square"
	"github.com/Houndie/dss-registration/dynamic/storage"
	"github.com/Houndie/dss-registration/dynamic/utility"
)

type purchaseItem struct {
	variationID string
	cost        int
}

type DiscountAmount interface {
	isDiscountAmount()
}

type PercentDiscount string

type DollarDiscount int

func (DollarDiscount) isDiscountAmount()  {}
func (PercentDiscount) isDiscountAmount() {}

type discount struct {
	id     string
	amount DiscountAmount
}

type squareData struct {
	fullWeekend     map[storage.WeekendPassTier]*purchaseItem
	danceOnly       *purchaseItem
	soloJazz        *purchaseItem
	mixAndMatch     *purchaseItem
	teamCompetition *purchaseItem
	tShirt          *purchaseItem

	studentDiscount *discount
	discounts       map[string]*discount
}

func singleVariationItem(o *square.CatalogItem) (*purchaseItem, error) {
	if len(o.Variations) != 1 {
		return nil, fmt.Errorf("expected 1 variation, found %d", len(o.Variations))
	}
	variation, ok := o.Variations[0].CatalogObjectType.(*square.CatalogItemVariation)
	if !ok {
		return nil, errors.New("item variation isn't variation")
	}
	return &purchaseItem{
		variationID: o.Variations[0].Id,
		cost:        variation.PriceMoney.Amount,
	}, nil
}

func getSquareCatalog(ctx context.Context, client SquareClient) (*squareData, error) {
	result := &squareData{}
	result.fullWeekend = map[storage.WeekendPassTier]*purchaseItem{}
	result.discounts = map[string]*discount{}
	objects := client.ListCatalog(ctx, []square.CatalogObjectType{square.CatalogObjectTypeItem, square.CatalogObjectTypeDiscount})

	var err error
	for objects.Next() {
		switch o := objects.Value().CatalogObjectType.(type) {
		case *square.CatalogItem:
			switch o.Name {
			case utility.MixAndMatchItem:
				result.mixAndMatch, err = singleVariationItem(o)
				if err != nil {
					return nil, fmt.Errorf("error fetching mix and match data: %w", err)
				}
			case utility.SoloJazzItem:
				result.soloJazz, err = singleVariationItem(o)
				if err != nil {
					return nil, fmt.Errorf("error fetching solo jazz data: %w", err)
				}
			case utility.TeamCompItem:
				result.teamCompetition, err = singleVariationItem(o)
				if err != nil {
					return nil, fmt.Errorf("error fetching team competition data: %w", err)
				}
			case utility.TShirtItem:
				result.tShirt, err = singleVariationItem(o)
				if err != nil {
					return nil, fmt.Errorf("error fetching t-shirt data: %w", err)
				}
			case utility.DancePassItem:
				for _, variation := range o.Variations {
					v, ok := variation.CatalogObjectType.(*square.CatalogItemVariation)
					if !ok {
						// Should never happen, but just move on
						continue
					}

					if v.Name != utility.DancePassPresaleName {
						continue
					}
					result.danceOnly = &purchaseItem{
						variationID: variation.Id,
						cost:        v.PriceMoney.Amount,
					}
				}
				if result.danceOnly == nil {
					return nil, errors.New("dance only item found, but no presale variation found")
				}
			case utility.WeekendPassItem:
				for _, variation := range o.Variations {
					v, ok := variation.CatalogObjectType.(*square.CatalogItemVariation)
					if !ok {
						// Should never happen, but just move on
						continue
					}

					for tier, name := range utility.WeekendPassName {
						if v.Name == name {
							result.fullWeekend[tier] = &purchaseItem{
								variationID: variation.Id,
								cost:        v.PriceMoney.Amount,
							}
							break
						}
					}
					//if not found, continue
				}
			}
		case *square.CatalogDiscount:
			var amount DiscountAmount
			switch t := o.DiscountType.(type) {
			case *square.CatalogDiscountFixedAmount:
				amount = DollarDiscount(t.AmountMoney.Amount)
			case *square.CatalogDiscountVariableAmount:
				amount = DollarDiscount(t.AmountMoney.Amount)
			case *square.CatalogDiscountFixedPercentage:
				amount = PercentDiscount(t.Percentage)
			case *square.CatalogDiscountVariablePercentage:
				amount = PercentDiscount(t.Percentage)
			default:
				return nil, errors.New("found unknown catalog discount type")

			}
			if o.Name == utility.StudentDiscountItem {
				result.studentDiscount = &discount{
					id:     objects.Value().Id,
					amount: amount,
				}
				continue
			}

			result.discounts[o.Name] = &discount{
				id:     objects.Value().Id,
				amount: amount,
			}
		}
	}
	if objects.Error() != nil {
		return nil, fmt.Errorf("error fetching all items from square: %w", objects.Error())
	}

	if result.fullWeekend[storage.Tier1] == nil {
		return nil, errors.New("no full weekend tier 1 data found")
	}
	if result.fullWeekend[storage.Tier2] == nil {
		return nil, errors.New("no full weekend tier 2 data found")
	}
	if result.fullWeekend[storage.Tier3] == nil {
		return nil, errors.New("no full weekend tier 3 data found")
	}
	if result.fullWeekend[storage.Tier4] == nil {
		return nil, errors.New("no full weekend tier 4 data found")
	}
	if result.fullWeekend[storage.Tier5] == nil {
		return nil, errors.New("no full weekend tier 5 data found")
	}
	if result.danceOnly == nil {
		return nil, errors.New("no dance only data found")
	}
	if result.mixAndMatch == nil {
		return nil, errors.New("no mix and match data found")
	}
	if result.teamCompetition == nil {
		return nil, errors.New("no team competition data found")
	}
	if result.soloJazz == nil {
		return nil, errors.New("no solo jazz data found")
	}
	if result.tShirt == nil {
		return nil, errors.New("no tshirt data found")
	}
	if result.studentDiscount == nil {
		return nil, errors.New("no student discount found")
	}
	return result, nil
}

type tierData struct {
	tier storage.WeekendPassTier
	cost int
}

func lowestInStockTier(ctx context.Context, s *squareData, squareClient SquareClient) (storage.WeekendPassTier, int, error) {
	weekendPassIds := make([]string, len(s.fullWeekend))
	tierMap := map[string]*tierData{}
	idx := 0
	for tier, weekendItem := range s.fullWeekend {
		weekendPassIds[idx] = weekendItem.variationID
		tierMap[weekendItem.variationID] = &tierData{
			tier: tier,
			cost: weekendItem.cost,
		}
		idx++
	}

	bestTier, bestCost := storage.Tier5, s.fullWeekend[storage.Tier5].cost
	counts := squareClient.BatchRetrieveInventoryCounts(ctx, weekendPassIds, nil, nil)
	for counts.Next() {
		quantity, err := strconv.ParseFloat(counts.Value().Quantity, 64)
		if err != nil {
			return 0, 0, fmt.Errorf("could not convert quantity %s to float: %w", counts.Value().Quantity, err)
		}
		if quantity > 0 {
			thisTier := tierMap[counts.Value().CatalogObjectId]
			if thisTier.tier < bestTier {
				bestTier = thisTier.tier
				bestCost = thisTier.cost
			}
		}
	}
	if counts.Error() != nil {
		return 0, 0, fmt.Errorf("error retrieving inventory counts from square: %w", counts.Error())
	}

	return bestTier, bestCost, nil

}

type paymentData struct {
	weekendPassPaid     bool
	danceOnlyPaid       bool
	mixAndMatchPaid     bool
	soloJazzPaid        bool
	teamCompetitionPaid bool
	tShirtPaid          bool
}

func getSquarePayments(ctx context.Context, client SquareClient, squareData *squareData, locationId string, orderIDs []string) (*paymentData, error) {
	orders, err := client.BatchRetrieveOrders(ctx, locationId, orderIDs)
	if err != nil {
		return nil, fmt.Errorf("error fetching orders from square")
	}

	pd := &paymentData{}
	for _, order := range orders {
		if order.State != square.OrderStateCompleted {
			continue
		}
		for _, lineItem := range order.LineItems {
			for _, purchaseItem := range squareData.fullWeekend {
				if purchaseItem.variationID == lineItem.CatalogObjectId {
					pd.weekendPassPaid = true
					break
				}
			}
			switch lineItem.CatalogObjectId {
			case squareData.danceOnly.variationID:
				pd.danceOnlyPaid = true
			case squareData.mixAndMatch.variationID:
				pd.mixAndMatchPaid = true
			case squareData.soloJazz.variationID:
				pd.soloJazzPaid = true
			case squareData.teamCompetition.variationID:
				pd.teamCompetitionPaid = true
			case squareData.tShirt.variationID:
				pd.tShirtPaid = true
			}
		}
	}
	return pd, nil
}

func createCheckout(ctx context.Context, client SquareClient, locationID, idempotencyKey string, order *square.CreateOrderRequest, userEmail string, redirectUrl string) (string, string, error) {
	checkout, err := client.CreateCheckout(ctx, locationID, idempotencyKey, order, false, utility.SmackdownEmail, userEmail, nil, redirectUrl, nil, "")
	if err != nil {
		errorList, ok := err.(*square.ErrorList)

		// If this error is anything other than "can't create checkouts worth less than a dollar"
		if !ok || len(errorList.Errors) > 1 || errorList.Errors[0].Category != square.ErrorCategoryInvalidRequestError || errorList.Errors[0].Code != square.ErrorCodeValueTooLow || errorList.Errors[0].Field != "order.total_money.amount" {
			return "", "", fmt.Errorf("error creating square checkout: %w", err)
		}
		return redirectUrl, "", nil
	}
	return checkout.CheckoutPageUrl, checkout.Order.Id, nil
}
