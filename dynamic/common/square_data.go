package common

import (
	"context"
	"errors"
	"fmt"
	"strconv"

	"github.com/Houndie/dss-registration/dynamic/storage"
	"github.com/Houndie/dss-registration/dynamic/utility"
	"github.com/Houndie/square-go"
	"github.com/Houndie/square-go/objects"
)

type PurchaseItem struct {
	VariationID string
	Cost        int
}

type DiscountAmount interface {
	isDiscountAmount()
}

type PercentDiscount string

type DollarDiscount int

func (DollarDiscount) isDiscountAmount()  {}
func (PercentDiscount) isDiscountAmount() {}

type Discount struct {
	ID     string
	Amount DiscountAmount
}

type SquareData struct {
	FullWeekend     map[storage.WeekendPassTier]*PurchaseItem
	DanceOnly       *PurchaseItem
	SoloJazz        *PurchaseItem
	MixAndMatch     *PurchaseItem
	TeamCompetition *PurchaseItem
	TShirt          *PurchaseItem

	StudentDiscount *Discount
	Discounts       map[string]*Discount
}

func singleVariationItem(o *objects.CatalogItem) (*PurchaseItem, error) {
	if len(o.Variations) != 1 {
		return nil, fmt.Errorf("expected 1 variation, found %d", len(o.Variations))
	}
	variation, ok := o.Variations[0].CatalogObjectType.(*objects.CatalogItemVariation)
	if !ok {
		return nil, errors.New("item variation isn't variation")
	}
	return &PurchaseItem{
		VariationID: o.Variations[0].ID,
		Cost:        variation.PriceMoney.Amount,
	}, nil
}

func GetSquareCatalog(ctx context.Context, client *square.Client) (*SquareData, error) {
	result := &SquareData{}
	result.FullWeekend = map[storage.WeekendPassTier]*PurchaseItem{}
	result.Discounts = map[string]*Discount{}
	catalogObjects := client.Catalog.List(ctx, []objects.CatalogObjectType{objects.CatalogObjectTypeItem, objects.CatalogObjectTypeDiscount})

	var err error
	for catalogObjects.Next() {
		switch o := catalogObjects.Value().CatalogObjectType.(type) {
		case *objects.CatalogItem:
			switch o.Name {
			case utility.MixAndMatchItem:
				result.MixAndMatch, err = singleVariationItem(o)
				if err != nil {
					return nil, fmt.Errorf("error fetching mix and match data: %w", err)
				}
			case utility.SoloJazzItem:
				result.SoloJazz, err = singleVariationItem(o)
				if err != nil {
					return nil, fmt.Errorf("error fetching solo jazz data: %w", err)
				}
			case utility.TeamCompItem:
				result.TeamCompetition, err = singleVariationItem(o)
				if err != nil {
					return nil, fmt.Errorf("error fetching team competition data: %w", err)
				}
			case utility.TShirtItem:
				result.TShirt, err = singleVariationItem(o)
				if err != nil {
					return nil, fmt.Errorf("error fetching t-shirt data: %w", err)
				}
			case utility.DancePassItem:
				for _, variation := range o.Variations {
					v, ok := variation.CatalogObjectType.(*objects.CatalogItemVariation)
					if !ok {
						// Should never happen, but just move on
						continue
					}

					if v.Name != utility.DancePassPresaleName {
						continue
					}
					result.DanceOnly = &PurchaseItem{
						VariationID: variation.ID,
						Cost:        v.PriceMoney.Amount,
					}
				}
				if result.DanceOnly == nil {
					return nil, errors.New("dance only item found, but no presale variation found")
				}
			case utility.WeekendPassItem:
				for _, variation := range o.Variations {
					v, ok := variation.CatalogObjectType.(*objects.CatalogItemVariation)
					if !ok {
						// Should never happen, but just move on
						continue
					}

					for tier, name := range utility.WeekendPassName {
						if v.Name == name {
							result.FullWeekend[tier] = &PurchaseItem{
								VariationID: variation.ID,
								Cost:        v.PriceMoney.Amount,
							}
							break
						}
					}
					//if not found, continue
				}
			}
		case *objects.CatalogDiscount:
			var amount DiscountAmount
			switch t := o.DiscountType.(type) {
			case *objects.CatalogDiscountFixedAmount:
				amount = DollarDiscount(t.AmountMoney.Amount)
			case *objects.CatalogDiscountVariableAmount:
				amount = DollarDiscount(t.AmountMoney.Amount)
			case *objects.CatalogDiscountFixedPercentage:
				amount = PercentDiscount(t.Percentage)
			case *objects.CatalogDiscountVariablePercentage:
				amount = PercentDiscount(t.Percentage)
			default:
				return nil, errors.New("found unknown catalog discount type")

			}
			if o.Name == utility.StudentDiscountItem {
				result.StudentDiscount = &Discount{
					ID:     catalogObjects.Value().ID,
					Amount: amount,
				}
				continue
			}

			result.Discounts[o.Name] = &Discount{
				ID:     catalogObjects.Value().ID,
				Amount: amount,
			}
		}
	}
	if catalogObjects.Error() != nil {
		return nil, fmt.Errorf("error fetching all items from square: %w", catalogObjects.Error())
	}

	if result.FullWeekend[storage.Tier1] == nil {
		return nil, errors.New("no full weekend tier 1 data found")
	}
	if result.FullWeekend[storage.Tier2] == nil {
		return nil, errors.New("no full weekend tier 2 data found")
	}
	if result.FullWeekend[storage.Tier3] == nil {
		return nil, errors.New("no full weekend tier 3 data found")
	}
	if result.FullWeekend[storage.Tier4] == nil {
		return nil, errors.New("no full weekend tier 4 data found")
	}
	if result.FullWeekend[storage.Tier5] == nil {
		return nil, errors.New("no full weekend tier 5 data found")
	}
	if result.DanceOnly == nil {
		return nil, errors.New("no dance only data found")
	}
	if result.MixAndMatch == nil {
		return nil, errors.New("no mix and match data found")
	}
	if result.TeamCompetition == nil {
		return nil, errors.New("no team competition data found")
	}
	if result.SoloJazz == nil {
		return nil, errors.New("no solo jazz data found")
	}
	if result.TShirt == nil {
		return nil, errors.New("no tshirt data found")
	}
	if result.StudentDiscount == nil {
		return nil, errors.New("no student discount found")
	}
	return result, nil
}

type tierData struct {
	tier storage.WeekendPassTier
	cost int
}

func LowestInStockTier(ctx context.Context, s *SquareData, squareClient *square.Client) (storage.WeekendPassTier, int, error) {
	weekendPassIDs := make([]string, len(s.FullWeekend))
	tierMap := map[string]*tierData{}
	idx := 0
	for tier, weekendItem := range s.FullWeekend {
		weekendPassIDs[idx] = weekendItem.VariationID
		tierMap[weekendItem.VariationID] = &tierData{
			tier: tier,
			cost: weekendItem.Cost,
		}
		idx++
	}

	bestTier, bestCost := storage.Tier5, s.FullWeekend[storage.Tier5].Cost
	counts := squareClient.Inventory.BatchRetrieveCounts(ctx, weekendPassIDs, nil, nil)
	for counts.Next() {
		quantity, err := strconv.ParseFloat(counts.Value().Quantity, 64)
		if err != nil {
			return 0, 0, fmt.Errorf("could not convert quantity %s to float: %w", counts.Value().Quantity, err)
		}
		if quantity > 0 {
			thisTier := tierMap[counts.Value().CatalogObjectID]
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

type PaymentData struct {
	WeekendPassPaid     bool
	DanceOnlyPaid       bool
	MixAndMatchPaid     bool
	SoloJazzPaid        bool
	TeamCompetitionPaid bool
	TShirtPaid          bool
}

func GetSquarePayments(ctx context.Context, client *square.Client, squareData *SquareData, locationID string, orderIDs []string) (*PaymentData, error) {
	orders, err := client.Orders.BatchRetrieve(ctx, locationID, orderIDs)
	if err != nil {
		return nil, fmt.Errorf("error fetching orders from square")
	}

	pd := &PaymentData{}
	for _, order := range orders {
		if order.State != objects.OrderStateCompleted {
			continue
		}
		for _, lineItem := range order.LineItems {
			for _, purchaseItem := range squareData.FullWeekend {
				if purchaseItem.VariationID == lineItem.CatalogObjectID {
					pd.WeekendPassPaid = true
					break
				}
			}
			switch lineItem.CatalogObjectID {
			case squareData.DanceOnly.VariationID:
				pd.DanceOnlyPaid = true
			case squareData.MixAndMatch.VariationID:
				pd.MixAndMatchPaid = true
			case squareData.SoloJazz.VariationID:
				pd.SoloJazzPaid = true
			case squareData.TeamCompetition.VariationID:
				pd.TeamCompetitionPaid = true
			case squareData.TShirt.VariationID:
				pd.TShirtPaid = true
			}
		}
	}
	return pd, nil
}

func CreateCheckout(ctx context.Context, client *square.Client, locationID, idempotencyKey string, order *objects.CreateOrderRequest, userEmail string, redirectUrl string) (string, string, error) {
	checkout, err := client.Checkout.Create(ctx, locationID, idempotencyKey, order, false, utility.SmackdownEmail, userEmail, nil, redirectUrl, nil, "")
	if err != nil {
		errorList, ok := err.(*objects.ErrorList)

		// If this error is anything other than "can't create checkouts worth less than a dollar"
		if !ok || len(errorList.Errors) > 1 || errorList.Errors[0].Category != objects.ErrorCategoryInvalidRequestError || errorList.Errors[0].Code != objects.ErrorCodeValueTooLow || errorList.Errors[0].Field != "order.total_money.amount" {
			return "", "", fmt.Errorf("error creating square checkout: %w", err)
		}
		return redirectUrl, "", nil
	}
	return checkout.CheckoutPageURL, checkout.Order.ID, nil
}
