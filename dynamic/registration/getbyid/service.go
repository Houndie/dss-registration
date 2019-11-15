package getbyid

import (
	"context"
	"fmt"
	"strconv"
	"time"

	"github.com/Houndie/dss-registration/dynamic/authorizer"
	"github.com/Houndie/dss-registration/dynamic/registration/common"
	"github.com/Houndie/dss-registration/dynamic/square"
	"github.com/Houndie/dss-registration/dynamic/storage"
	"github.com/Houndie/dss-registration/dynamic/utility"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
)

type tierData struct {
	tier int
	cost int
}

type Authorizer interface {
	Userinfo(ctx context.Context, accessToken string) (*authorizer.Userinfo, error)
}

type Store interface {
	GetRegistrationById(ctx context.Context, id string) (*StoreRegistration, error)
	UpdateRegistrationTier(ctx context.Context, id string, newTier common.WeekendPassTier) error
}

type SquareClient interface {
	ListLocations(ctx context.Context) ([]*square.Location, error)
	BatchRetrieveOrders(ctx context.Context, locationId string, orderIds []string) ([]*square.Order, error)
	BatchRetrieveInventoryCounts(ctx context.Context, catalogObjectIds, locationIds []string, updatedAfter *time.Time) square.BatchRetrieveInventoryCountsIterator
	ListCatalog(ctx context.Context, types []square.CatalogObjectType) square.ListCatalogIterator
}

type Service struct {
	logger     *logrus.Logger
	authorizer Authorizer
	store      Store
	client     SquareClient
}

func NewService(logger *logrus.Logger, authorizer Authorizer, store Store, client SquareClient) *Service {
	return &Service{
		logger:     logger,
		authorizer: authorizer,
		store:      store,
		client:     client,
	}
}

func (s *Service) GetById(ctx context.Context, token, registrationId string) (*Registration, error) {
	s.logger.Trace("In get by id service")
	s.logger.Tracef("fetching user-info for token %s", token)
	userinfo, err := s.authorizer.Userinfo(ctx, token)
	if err != nil {
		msg := "could not authorize user"
		s.logger.WithError(err).Debug(msg)
		return nil, errors.Wrap(err, msg)
	}
	s.logger.Tracef("found user %s", userinfo.UserId)

	s.logger.Tracef("fetching registrations for user %s", userinfo.UserId)
	r, err := s.store.GetRegistrationById(ctx, registrationId)
	if err != nil {
		switch errors.Cause(err).(type) {
		case storage.ErrNotFound:
			newErr := ErrBadRegistrationId{registrationId}
			s.logger.WithError(err).Debug(newErr)
			return nil, newErr
		default:
			msg := "error fetching registrations from store"
			s.logger.WithError(err).Error(msg)
			return nil, errors.Wrap(err, msg)
		}
	}
	s.logger.Trace("found registration")

	if r.UserId != userinfo.UserId {
		err := ErrBadRegistrationId{registrationId}
		s.logger.WithError(err).Debug("user id does not match that of found registration")
		s.logger.WithError(err).Tracef("registration provided user id %s, user provided %s", r.UserId, userinfo.UserId)
		return nil, err
	}

	s.logger.Trace("fetching locations from square")
	locations, err := s.client.ListLocations(ctx)
	if err != nil {
		msg := "error listing locations from square"
		s.logger.WithError(err).Error(msg)
		return nil, errors.Wrap(err, msg)
	}
	s.logger.Tracef("found %d locations", len(locations))

	if len(locations) != 1 {
		msg := fmt.Errorf("found unexpected number of locations %d", len(locations))
		s.logger.Error(msg)
		return nil, msg
	}
	s.logger.Tracef("found location %s", locations[0].Id)

	var unpaidItems *UnpaidItems
	unpaidFullWeekend := false
	if len(r.OrderIds) > 0 {
		s.logger.Trace("retrieving orders from square")
		squareOrders, err := s.client.BatchRetrieveOrders(ctx, locations[0].Id, r.OrderIds)
		if err != nil {
			msg := "error retrieving orders matching ids"
			s.logger.WithError(err).Error(msg)
			return nil, errors.Wrap(err, msg)
		}
		unpaidItems = &UnpaidItems{
			OrderIds: []string{},
			Items:    []string{},
			Cost:     0,
		}
		for _, squareOrder := range squareOrders {
			if squareOrder.State != square.OrderStateOpen {
				continue
			}
			for _, squareOrderItem := range squareOrder.LineItems {
				if squareOrderItem.Name == utility.WeekendPassItem {
					unpaidFullWeekend = true
				} else {
					unpaidItems.Items = append(unpaidItems.Items, squareOrderItem.Name)
					unpaidItems.Cost += squareOrderItem.TotalMoney.Amount
				}
			}
			unpaidItems.OrderIds = append(unpaidItems.OrderIds, squareOrder.Id)
		}
		if len(unpaidItems.OrderIds) == 0 {
			unpaidItems = nil
		}
	}

	squareDiscounts := map[string]common.ItemDiscount{}
	for _, d := range r.Discounts {
		for _, sd := range d.Discounts {
			squareDiscounts[sd.Name] = nil
		}
	}

	fetchTypes := []square.CatalogObjectType{square.CatalogObjectTypeDiscount}
	tiers := map[string]tierData{}
	myTierId := ""
	myTierName := ""
	if unpaidFullWeekend {
		fetchTypes = append(fetchTypes, square.CatalogObjectTypeItem)
		weekendPass, ok := r.PassType.(*common.WeekendPass)
		if !ok {
			s.logger.Error("weekend pass found in unpaid order, but not in registration")
			return nil, errors.New("weekend pass found in unpaid order, but not in registration")
		}
		switch weekendPass.Tier {
		case common.WeekendPassTier1:
			myTierName = utility.WeekendPassTier1Name
		case common.WeekendPassTier2:
			myTierName = utility.WeekendPassTier2Name
		case common.WeekendPassTier3:
			myTierName = utility.WeekendPassTier3Name
		case common.WeekendPassTier4:
			myTierName = utility.WeekendPassTier4Name
		case common.WeekendPassTier5:
			myTierName = utility.WeekendPassTier5Name
		}
	}
	objects := s.client.ListCatalog(ctx, fetchTypes)
	for objects.Next() {
		switch object := objects.Value().CatalogObjectType.(type) {
		case *square.CatalogDiscount:
			_, ok := squareDiscounts[object.Name]
			if !ok {
				continue
			}

			var itemDiscount common.ItemDiscount
			switch t := object.DiscountType.(type) {
			case *square.CatalogDiscountFixedAmount:
				itemDiscount = &common.DollarDiscount{
					Amount: t.AmountMoney.Amount,
				}
			case *square.CatalogDiscountVariableAmount:
				itemDiscount = &common.DollarDiscount{
					Amount: t.AmountMoney.Amount,
				}
			case *square.CatalogDiscountFixedPercentage:
				itemDiscount = &common.PercentDiscount{
					Amount: t.Percentage,
				}
			case *square.CatalogDiscountVariablePercentage:
				itemDiscount = &common.PercentDiscount{
					Amount: t.Percentage,
				}
			default:
				err := errors.New("unknown item discount type found from square")
				s.logger.Error(err)
				return nil, err
			}

			squareDiscounts[object.Name] = itemDiscount
		case *square.CatalogItem:
			if object.Name != utility.WeekendPassItem {
				continue
			}
			for _, v := range object.Variations {
				variation, ok := v.CatalogObjectType.(*square.CatalogItemVariation)
				if !ok {
					err := "Invalid response from square...item variation isn't a variation?"
					s.logger.Error(err)
					return nil, errors.New(err)
				}
				switch variation.Name {
				case utility.WeekendPassTier1Name:
					tiers[v.Id] = tierData{1, variation.PriceMoney.Amount}
				case utility.WeekendPassTier2Name:
					tiers[v.Id] = tierData{2, variation.PriceMoney.Amount}
				case utility.WeekendPassTier3Name:
					tiers[v.Id] = tierData{3, variation.PriceMoney.Amount}
				case utility.WeekendPassTier4Name:
					tiers[v.Id] = tierData{4, variation.PriceMoney.Amount}
				case utility.WeekendPassTier5Name:
					tiers[v.Id] = tierData{5, variation.PriceMoney.Amount}
				default: // Do nothing, we have other names that are allowable
				}
				if variation.Name == myTierName {
					s.logger.Trace("Found weekend pass")
					myTierId = v.Id
				}
			}
		default:
			s.logger.Error("found non discount or catalog object when discount and item was requested")
		}
	}
	if objects.Error() != nil {
		wrap := "error fetching catalog objects from square"
		s.logger.WithError(err).Error(wrap)
		return nil, errors.Wrap(err, wrap)
	}
	discounts := []*Discount{}
	s.logger.Tracef("parsing %d discounts", len(r.Discounts))
	for _, discount := range r.Discounts {
		singleDiscounts := make([]*SingleDiscount, len(discount.Discounts))
		allDiscountsFound := true
		for i, sd := range discount.Discounts {
			itemDiscount, ok := squareDiscounts[sd.Name]
			if !ok {
				err := errors.New("impossible code path, somehow a square discount was not added or removed from square discount map")
				s.logger.Error(err)
				return nil, err
			}
			if itemDiscount == nil {
				s.logger.Errorf("Discount %s was applied but is no longer found in square store.  Omitting this code in result", sd.Name)
				allDiscountsFound = false
				break
			}
			singleDiscounts[i] = &SingleDiscount{
				ItemDiscount: itemDiscount,
				AppliedTo:    sd.AppliedTo,
			}
		}
		if !allDiscountsFound {
			continue
		}
		discounts = append(discounts, &Discount{
			Code:      discount.Code,
			Discounts: singleDiscounts,
		})
	}
	s.logger.Tracef("%d discounts parsed and applied to registration", len(discounts))

	// Figure out if the current tier is still good
	passType := r.PassType
	updatedTier := false
	if unpaidFullWeekend {
		lowestTier := 999999
		lowestTierCost := 0
		outOfStock := false
		weekendPassIds := []string{}
		for weekendPassId, _ := range tiers {
			weekendPassIds = append(weekendPassIds, weekendPassId)
		}
		counts := s.client.BatchRetrieveInventoryCounts(ctx, weekendPassIds, nil, nil)
		for counts.Next() {
			quantity, err := strconv.ParseFloat(counts.Value().Quantity, 64)
			if err != nil {
				s.logger.WithField("quantity", counts.Value().Quantity).Error("could not convert quantity to float")
				return nil, errors.Wrapf(err, "could not convert quantity %s to float", counts.Value().Quantity)
			}
			if counts.Value().CatalogObjectId == myTierId {
				if quantity < 1 {
					outOfStock = true
				}
			}
			if quantity > 0 {
				thisTier := tiers[counts.Value().CatalogObjectId]
				s.logger.Tracef("tier %d", thisTier.tier)
				if thisTier.tier < lowestTier {
					s.logger.Tracef("new lowest tier found")
					lowestTier = thisTier.tier
					lowestTierCost = thisTier.cost
				}
			}
		}
		unpaidItems.Items = append(unpaidItems.Items, utility.WeekendPassItem)
		unpaidItems.Cost += lowestTierCost
		passType = &common.WeekendPass{
			Tier:  common.WeekendPassTier(lowestTier),
			Level: r.PassType.(*common.WeekendPass).Level,
		}
		if outOfStock {
			s.store.UpdateRegistrationTier(ctx, registrationId, common.WeekendPassTier(lowestTier))
			updatedTier = true
		}
	}

	return &Registration{
		FirstName:       r.FirstName,
		LastName:        r.LastName,
		StreetAddress:   r.StreetAddress,
		City:            r.City,
		State:           r.State,
		ZipCode:         r.ZipCode,
		Email:           r.Email,
		HomeScene:       r.HomeScene,
		IsStudent:       r.IsStudent,
		PassType:        passType,
		MixAndMatch:     r.MixAndMatch,
		SoloJazz:        r.SoloJazz,
		TeamCompetition: r.TeamCompetition,
		TShirt:          r.TShirt,
		Housing:         r.Housing,
		UnpaidItems:     unpaidItems,
		Discounts:       discounts,
		UpdatedTier:     updatedTier,
	}, nil
}
