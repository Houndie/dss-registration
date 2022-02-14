package common

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"strconv"

	"github.com/Houndie/dss-registration/dynamic/storage"
	"github.com/Houndie/dss-registration/dynamic/utility"
	"github.com/Houndie/square-go"
	"github.com/Houndie/square-go/checkout"
	"github.com/Houndie/square-go/inventory"
	"github.com/Houndie/square-go/objects"
	"github.com/Houndie/square-go/orders"
)

type PurchaseItem struct {
	ID    string `json:"id"`
	Price int    `json:"price"`
}

type DiscountAmount interface {
	isDiscountAmount()
}

type PercentDiscount string

type DollarDiscount int

func (DollarDiscount) isDiscountAmount()  {}
func (PercentDiscount) isDiscountAmount() {}

type Discount struct {
	ID        string               `json:"id"`
	Amount    DiscountAmount       `json:"-"`
	AppliedTo storage.PurchaseItem `json:"-"`
}

type discountAlias Discount

type jsonDiscount struct {
	*discountAlias
	Amount       int    `json:"amount"`
	Percentage   string `json:"percentage"`
	DiscountType string `json:"discount_type"`
	AppliedTo    string `json:"applied_to"`
}

var (
	appliedToToJSON = map[storage.PurchaseItem]string{
		storage.FullWeekendPurchaseItem:     "Full Weekend",
		storage.DanceOnlyPurchaseItem:       "Dance Only",
		storage.MixAndMatchPurchaseItem:     "Mix And Match",
		storage.SoloJazzPurchaseItem:        "Solo Jazz",
		storage.TeamCompetitionPurchaseItem: "Team Competition",
		storage.TShirtPurchaseItem:          "TShirt",
	}

	appliedToFromJSON = map[string]storage.PurchaseItem{
		"Full Weekend":     storage.FullWeekendPurchaseItem,
		"Dance Only":       storage.DanceOnlyPurchaseItem,
		"Mix And Match":    storage.MixAndMatchPurchaseItem,
		"Solo Jazz":        storage.SoloJazzPurchaseItem,
		"Team Competition": storage.TeamCompetitionPurchaseItem,
		"TShirt":           storage.TShirtPurchaseItem,
	}
)

func (d *Discount) MarshalJSON() ([]byte, error) {
	j := jsonDiscount{
		discountAlias: (*discountAlias)(d),
	}

	appliedTo, ok := appliedToToJSON[d.AppliedTo]
	if !ok {
		return nil, fmt.Errorf("unknown applied to found: %v", d.AppliedTo)
	}

	j.AppliedTo = appliedTo

	switch a := d.Amount.(type) {
	case DollarDiscount:
		j.Amount = (int)(a)
		j.DiscountType = "FIXED_AMOUNT"
	case PercentDiscount:
		j.Percentage = (string)(a)
		j.DiscountType = "FIXED_PERCENTAGE"
	default:
		return nil, errors.New("unknown discount type found")
	}

	b, err := json.Marshal(j)
	if err != nil {
		return nil, fmt.Errorf("error marshaling discount type: %w", err)
	}

	return b, nil
}

func (d *Discount) UnmarshalJSON(b []byte) error {
	j := jsonDiscount{
		discountAlias: (*discountAlias)(d),
	}

	if err := json.Unmarshal(b, &j); err != nil {
		return fmt.Errorf("error unmarhsaling discount type: %w", err)
	}

	appliedTo, ok := appliedToFromJSON[j.AppliedTo]
	if !ok {
		return fmt.Errorf("unknown applied to found: %v", j.AppliedTo)
	}
	d.AppliedTo = appliedTo

	switch j.DiscountType {
	case "FIXED_AMOUNT", "VARIABLE_AMOUNT":
		d.Amount = DollarDiscount(j.Amount)
	case "FIXED_PERCENTAGE", "VARIABLE_PERCENTAGE":
		d.Amount = PercentDiscount(j.Percentage)
	default:
		return fmt.Errorf("unknown discount type found: %v", j.DiscountType)
	}

	return nil
}

var (
	tierToJson = map[storage.WeekendPassTier]string{
		storage.Tier1: "Tier 1",
		storage.Tier2: "Tier 2",
		storage.Tier3: "Tier 3",
		storage.Tier4: "Tier 4",
		storage.Tier5: "Tier 5",
	}
	tierFromJson = map[string]storage.WeekendPassTier{
		"Tier 1": storage.Tier1,
		"Tier 2": storage.Tier2,
		"Tier 3": storage.Tier3,
		"Tier 4": storage.Tier4,
		"Tier 5": storage.Tier5,
	}

	roleToJson = map[storage.MixAndMatchRole]string{
		storage.MixAndMatchRoleLeader:   "Leader",
		storage.MixAndMatchRoleFollower: "Follower",
	}
	roleFromJson = map[string]storage.MixAndMatchRole{
		"Leader":   storage.MixAndMatchRoleLeader,
		"Follower": storage.MixAndMatchRoleFollower,
	}

	styleToJson = map[storage.TShirtStyle]string{
		storage.TShirtStyleUnisexS:   "Unisex Small",
		storage.TShirtStyleUnisexM:   "Unisex Medium",
		storage.TShirtStyleUnisexL:   "Unisex Large",
		storage.TShirtStyleUnisexXL:  "Unisex XL",
		storage.TShirtStyleUnisex2XL: "Unisex 2XL",
		storage.TShirtStyleUnisex3XL: "Unisex 3XL",
		storage.TShirtStyleBellaS:    "Bella Small",
		storage.TShirtStyleBellaM:    "Bella Medium",
		storage.TShirtStyleBellaL:    "Bella Large",
		storage.TShirtStyleBellaXL:   "Bella XL",
		storage.TShirtStyleBella2XL:  "Bella 2XL",
	}

	styleFromJson = map[string]storage.TShirtStyle{
		"Unisex Small":  storage.TShirtStyleUnisexS,
		"Unisex Medium": storage.TShirtStyleUnisexM,
		"Unisex Large":  storage.TShirtStyleUnisexL,
		"Unisex XL":     storage.TShirtStyleUnisexXL,
		"Unisex 2XL":    storage.TShirtStyleUnisex2XL,
		"Unisex 3XL":    storage.TShirtStyleUnisex3XL,
		"Bella Small":   storage.TShirtStyleBellaS,
		"Bella Medium":  storage.TShirtStyleBellaM,
		"Bella Large":   storage.TShirtStyleBellaL,
		"Bella XL":      storage.TShirtStyleBellaXL,
		"Bella 2XL":     storage.TShirtStyleBella2XL,
	}
)

type PurchaseItems struct {
	FullWeekend     map[storage.WeekendPassTier]*PurchaseItem `json:"-"`
	DanceOnly       *PurchaseItem                             `json:"dance_only_pass"`
	SoloJazz        *PurchaseItem                             `json:"solo_jazz"`
	MixAndMatch     map[storage.MixAndMatchRole]*PurchaseItem `json:"-"`
	TeamCompetition *PurchaseItem                             `json:"team_competition"`
	TShirt          map[storage.TShirtStyle]*PurchaseItem     `json:"-"`
}

type purchaseItemsAlias PurchaseItems

type jsonPurchaseItems struct {
	*purchaseItemsAlias
	FullWeekend map[string]*PurchaseItem `json:"full_weekend_pass"`
	MixAndMatch map[string]*PurchaseItem `json:"mix_and_match"`
	TShirt      map[string]*PurchaseItem `json:"t_shirt"`
}

func (p *PurchaseItems) MarshalJSON() ([]byte, error) {
	j := jsonPurchaseItems{
		purchaseItemsAlias: (*purchaseItemsAlias)(p),
	}

	if len(p.FullWeekend) > 0 {
		j.FullWeekend = map[string]*PurchaseItem{}
		for tier, item := range p.FullWeekend {
			jTier, ok := tierToJson[tier]
			if !ok {
				return nil, fmt.Errorf("unknown tier found: %v", tier)
			}

			j.FullWeekend[jTier] = item
		}
	}

	if len(p.MixAndMatch) > 0 {
		j.MixAndMatch = map[string]*PurchaseItem{}
		for role, item := range p.MixAndMatch {
			jRole, ok := roleToJson[role]
			if !ok {
				return nil, fmt.Errorf("unknown role found: %v", role)
			}

			j.MixAndMatch[jRole] = item
		}
	}

	if len(p.TShirt) > 0 {
		j.TShirt = map[string]*PurchaseItem{}
		for style, item := range p.TShirt {
			jStyle, ok := styleToJson[style]
			if !ok {
				return nil, fmt.Errorf("unknown style found: %v", style)
			}

			j.TShirt[jStyle] = item
		}
	}

	b, err := json.Marshal(j)
	if err != nil {
		return nil, fmt.Errorf("error marshalling purchase items type: %w", err)
	}

	return b, nil
}

func (p *PurchaseItems) UnmarshalJSON(b []byte) error {
	j := jsonPurchaseItems{
		purchaseItemsAlias: (*purchaseItemsAlias)(p),
	}

	if err := json.Unmarshal(b, &j); err != nil {
		return fmt.Errorf("error unmarshaling purchase items type: %w", err)
	}

	if len(j.FullWeekend) > 0 {
		p.FullWeekend = map[storage.WeekendPassTier]*PurchaseItem{}
		for jTier, item := range j.FullWeekend {
			tier, ok := tierFromJson[jTier]
			if !ok {
				return fmt.Errorf("unknown tier found: %v", jTier)
			}

			p.FullWeekend[tier] = item
		}
	}

	if len(j.MixAndMatch) > 0 {
		p.MixAndMatch = map[storage.MixAndMatchRole]*PurchaseItem{}
		for jRole, item := range j.MixAndMatch {
			role, ok := roleFromJson[jRole]
			if !ok {
				return fmt.Errorf("unknown role found: %v", jRole)
			}

			p.MixAndMatch[role] = item
		}
	}

	if len(j.TShirt) > 0 {
		p.TShirt = map[storage.TShirtStyle]*PurchaseItem{}
		for jStyle, item := range j.TShirt {
			style, ok := styleFromJson[jStyle]
			if !ok {
				return fmt.Errorf("unknown style found: %v", jStyle)
			}

			p.TShirt[style] = item
		}
	}

	return nil
}

type Discounts struct {
	StudentDiscount *Discount              `json:"student"`
	CodeDiscounts   map[string][]*Discount `json:"code"`
}

type SquareData struct {
	PurchaseItems *PurchaseItems `json:"purchase_items"`
	Discounts     *Discounts     `json:"discounts"`
}

type tierData struct {
	tier  storage.WeekendPassTier
	price int
}

func LowestInStockTier(ctx context.Context, s map[storage.WeekendPassTier]*PurchaseItem, squareClient *square.Client) (storage.WeekendPassTier, int, error) {
	weekendPassIDs := make([]string, len(s))
	tierMap := map[string]*tierData{}
	idx := 0
	for tier, weekendItem := range s {
		weekendPassIDs[idx] = weekendItem.ID
		tierMap[weekendItem.ID] = &tierData{
			tier:  tier,
			price: weekendItem.Price,
		}
		idx++
	}

	bestTier, bestCost := storage.Tier5, s[storage.Tier5].Price
	res, err := squareClient.Inventory.BatchRetrieveCounts(ctx, &inventory.BatchRetrieveCountsRequest{
		CatalogObjectIDs: weekendPassIDs,
	})
	if err != nil {
		return 0, 0, fmt.Errorf("error fetching inventory counts from square: %w", err)
	}

	for res.Counts.Next() {
		count := res.Counts.Value().Count

		if count.State != objects.InventoryStateInStock {
			continue
		}

		quantity, err := strconv.ParseFloat(count.Quantity, 64)
		if err != nil {
			return 0, 0, fmt.Errorf("could not convert quantity %s to float: %w", count.Quantity, err)
		}
		if quantity > 0 {
			thisTier := tierMap[count.CatalogObjectID]
			if thisTier.tier < bestTier {
				bestTier = thisTier.tier
				bestCost = thisTier.price
			}
		}
	}
	if res.Counts.Error() != nil {
		return 0, 0, fmt.Errorf("error retrieving inventory counts from square: %w", res.Counts.Error())
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

func GetSquarePayments(ctx context.Context, client *square.Client, purchaseItems *PurchaseItems, locationID string, orderIDs map[string][]string) (map[string]*PaymentData, error) {
	orderIDList := []string{}
	orderIDMap := map[string]string{}
	for regID, regOrderIDs := range orderIDs {
		orderIDList = append(orderIDList, regOrderIDs...)
		for _, orderID := range regOrderIDs {
			orderIDMap[orderID] = regID
		}
	}

	ordersList := make([]*objects.Order, 0, len(orderIDList))
	for i := 0; i < len(orderIDList); i += 100 {
		chunkEnd := i + 100
		if chunkEnd > len(orderIDList) {
			chunkEnd = len(orderIDList)
		}

		res, err := client.Orders.BatchRetrieve(ctx, &orders.BatchRetrieveRequest{
			LocationID: locationID,
			OrderIDs:   orderIDList[i:chunkEnd],
		})
		if err != nil {
			return nil, fmt.Errorf("error fetching orders from square: %w", err)
		}

		ordersList = append(ordersList, res.Orders...)
	}

	pd := map[string]*PaymentData{}
	for regID, _ := range orderIDs {
		pd[regID] = &PaymentData{}
	}

	for _, order := range ordersList {
		if order.State != objects.OrderStateCompleted {
			continue
		}
		for _, lineItem := range order.LineItems {
			for _, purchaseItem := range purchaseItems.FullWeekend {
				if purchaseItem.ID == lineItem.CatalogObjectID {
					pd[orderIDMap[order.ID]].WeekendPassPaid = true
					break
				}
			}
			switch lineItem.CatalogObjectID {
			case purchaseItems.DanceOnly.ID:
				pd[orderIDMap[order.ID]].DanceOnlyPaid = true
			case purchaseItems.MixAndMatch[storage.MixAndMatchRoleLeader].ID, purchaseItems.MixAndMatch[storage.MixAndMatchRoleFollower].ID:
				pd[orderIDMap[order.ID]].MixAndMatchPaid = true
			case purchaseItems.SoloJazz.ID:
				pd[orderIDMap[order.ID]].SoloJazzPaid = true
			case purchaseItems.TeamCompetition.ID:
				pd[orderIDMap[order.ID]].TeamCompetitionPaid = true
			case purchaseItems.TShirt[storage.TShirtStyleUnisexS].ID, purchaseItems.TShirt[storage.TShirtStyleUnisexM].ID, purchaseItems.TShirt[storage.TShirtStyleUnisexL].ID, purchaseItems.TShirt[storage.TShirtStyleUnisexXL].ID, purchaseItems.TShirt[storage.TShirtStyleUnisex2XL].ID, purchaseItems.TShirt[storage.TShirtStyleUnisex3XL].ID, purchaseItems.TShirt[storage.TShirtStyleBellaS].ID, purchaseItems.TShirt[storage.TShirtStyleBellaM].ID, purchaseItems.TShirt[storage.TShirtStyleBellaL].ID, purchaseItems.TShirt[storage.TShirtStyleBellaXL].ID, purchaseItems.TShirt[storage.TShirtStyleBella2XL].ID:
				pd[orderIDMap[order.ID]].TShirtPaid = true
			}
		}
	}
	return pd, nil
}

func CreateCheckout(ctx context.Context, client *square.Client, locationID, idempotencyKey string, order *objects.CreateOrderRequest, userEmail string, redirectUrl string) (string, string, error) {
	res, err := client.Checkout.Create(ctx, &checkout.CreateRequest{
		LocationID:            locationID,
		IdempotencyKey:        idempotencyKey,
		Order:                 order,
		MerchantSupportEmail:  utility.SmackdownEmail,
		PrePopulateBuyerEmail: userEmail,
		RedirectURL:           redirectUrl,
	})
	if err != nil {
		errorList, ok := err.(*objects.ErrorList)

		// If this error is anything other than "can't create checkouts worth less than a dollar"
		if !ok || len(errorList.Errors) > 1 || errorList.Errors[0].Category != objects.ErrorCategoryInvalidRequestError || errorList.Errors[0].Code != objects.ErrorCodeValueTooLow || errorList.Errors[0].Field != "order.total_money.amount" {
			return "", "", fmt.Errorf("error creating square checkout: %w", err)
		}
		return redirectUrl, "", nil
	}
	return res.Checkout.CheckoutPageURL, res.Checkout.Order.ID, nil
}
