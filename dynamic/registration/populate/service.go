package populate

import (
	"context"
	"fmt"
	"strconv"
	"time"

	"github.com/Houndie/dss-registration/dynamic/square"
	"github.com/Houndie/dss-registration/dynamic/utility"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
)

type tierData struct {
	tier int
	cost int
}

type SquareClient interface {
	ListCatalog(ctx context.Context, types []square.CatalogObjectType) square.ListCatalogIterator
	BatchRetrieveInventoryCounts(ctx context.Context, catalogObjectIds, locationIds []string, updatedAfter *time.Time) square.BatchRetrieveInventoryCountsIterator
}

type Service struct {
	client SquareClient
	logger *logrus.Logger
}

func NewService(logger *logrus.Logger, client SquareClient) *Service {
	return &Service{
		client: client,
		logger: logger,
	}
}

func (s *Service) Populate(ctx context.Context) (*FormData, error) {
	s.logger.Trace("Fetching all items from square")
	objects := s.client.ListCatalog(ctx, []square.CatalogObjectType{square.CatalogObjectTypeItem, square.CatalogObjectTypeDiscount})
	s.logger.Trace("Iterating over square list catalog responses")
	res := &FormData{}
	tiers := map[string]tierData{}

	mixAndMatchFound := false
	teamCompFound := false
	soloJazzFound := false
	tShirtFound := false
	dancePassFound := false
	weekendPassFound := false
	studentDiscountFound := false
	for objects.Next() {
		s.logger.Trace("Found square object")
		switch item := objects.Value().CatalogObjectType.(type) {
		case *square.CatalogItem:
			s.logger.Tracef("Comparing item name %s to legend", item.Name)
			switch item.Name {
			case utility.MixAndMatchItem, utility.TeamCompItem, utility.SoloJazzItem, utility.TShirtItem:
				if len(item.Variations) != 1 {
					err := &ErrUnxpectedVariationCount{item.Name, len(item.Variations)}
					s.logger.Error(err)
					return nil, err
				}
				variation, ok := item.Variations[0].CatalogObjectType.(*square.CatalogItemVariation)
				if !ok {
					err := "Invalid response from square...item variation isn't a variation?"
					s.logger.Error(err)
					return nil, errors.New(err)
				}
				switch item.Name {
				case utility.MixAndMatchItem:
					mixAndMatchFound = true
					res.MixAndMatchCost = variation.PriceMoney.Amount
				case utility.TeamCompItem:
					teamCompFound = true
					res.TeamCompCost = variation.PriceMoney.Amount
				case utility.SoloJazzItem:
					soloJazzFound = true
					res.SoloJazzCost = variation.PriceMoney.Amount
				case utility.TShirtItem:
					tShirtFound = true
					res.TShirtCost = variation.PriceMoney.Amount
				default:
					err := errors.New("Impossible code path...how did I get here")
					s.logger.Error(err)
					return nil, err
				}
			case utility.DancePassItem:
				s.logger.Trace("Found dance pass item")
				for _, v := range item.Variations {
					variation, ok := v.CatalogObjectType.(*square.CatalogItemVariation)
					if !ok {
						err := "Invalid response from square...item variation isn't a variation?"
						s.logger.Error(err)
						return nil, errors.New(err)
					}
					if variation.Name == utility.DancePassPresaleName {
						s.logger.Trace("Found dance pass variant Presale")
						dancePassFound = true
						res.DancePassCost = variation.PriceMoney.Amount
						break
					}
				}
				if !dancePassFound {
					err := &ErrMissingVariation{item.Name, utility.DancePassPresaleName}
					s.logger.Error(err)
					return nil, err
				}
			case utility.WeekendPassItem:
				s.logger.Trace("Found weekend pass object")
				weekendPassFound = true

				weekendPassTier1Found := false
				weekendPassTier2Found := false
				weekendPassTier3Found := false
				weekendPassTier4Found := false
				weekendPassTier5Found := false
				for _, v := range item.Variations {
					variation, ok := v.CatalogObjectType.(*square.CatalogItemVariation)
					if !ok {
						err := "Invalid response from square...item variation isn't a variation?"
						s.logger.Error(err)
						return nil, errors.New(err)
					}

					s.logger.Tracef("Found variation with name %s and id %s", variation.Name, v.Id)
					switch variation.Name {
					case utility.WeekendPassTier1Name:
						weekendPassTier1Found = true
						s.logger.Trace("Variation matched tier 1")
						tiers[v.Id] = tierData{1, variation.PriceMoney.Amount}
					case utility.WeekendPassTier2Name:
						weekendPassTier2Found = true
						s.logger.Trace("Variation matched tier 2")
						tiers[v.Id] = tierData{2, variation.PriceMoney.Amount}
					case utility.WeekendPassTier3Name:
						weekendPassTier3Found = true
						s.logger.Trace("Variation matched tier 3")
						tiers[v.Id] = tierData{3, variation.PriceMoney.Amount}
					case utility.WeekendPassTier4Name:
						weekendPassTier4Found = true
						s.logger.Trace("Variation matched tier 4")
						tiers[v.Id] = tierData{4, variation.PriceMoney.Amount}
					case utility.WeekendPassTier5Name:
						weekendPassTier5Found = true
						s.logger.Trace("Variation matched tier 5")
						tiers[v.Id] = tierData{5, variation.PriceMoney.Amount}
					default: // Do nothing, we have other names that are allowable
						s.logger.Trace("Variation did not match tier list, moving on")
					}
				}
				var err error
				if !weekendPassTier1Found {
					err = &ErrMissingVariation{item.Name, utility.WeekendPassTier1Name}
				} else if !weekendPassTier2Found {
					err = &ErrMissingVariation{item.Name, utility.WeekendPassTier2Name}
				} else if !weekendPassTier3Found {
					err = &ErrMissingVariation{item.Name, utility.WeekendPassTier3Name}
				} else if !weekendPassTier4Found {
					err = &ErrMissingVariation{item.Name, utility.WeekendPassTier4Name}
				} else if !weekendPassTier5Found {
					err = &ErrMissingVariation{item.Name, utility.WeekendPassTier5Name}
				}
				if err != nil {
					s.logger.Error(err)
					return nil, err
				}
			}
		case *square.CatalogDiscount:
			if item.Name != utility.StudentDiscountItem {
				continue
			}
			studentDiscountFound = true
			switch dt := item.DiscountType.(type) {
			case *square.CatalogDiscountFixedAmount:
				res.StudentDiscount = dt.AmountMoney.Amount
			case *square.CatalogDiscountVariableAmount:
				res.StudentDiscount = dt.AmountMoney.Amount
			default:
				s.logger.Error("cannot handle a percentage student discount")
				return nil, errors.New("cannot handle a percentage student discount")
			}
		default:
			s.logger.Error("Unknown catalog item type found")
		}
	}
	if objects.Error() != nil {
		wrap := "error fetching all items from square"
		utility.LogSquareError(s.logger, objects.Error(), wrap)
		return nil, errors.Wrap(objects.Error(), wrap)
	}

	var err error
	if !mixAndMatchFound {
		err = &ErrMissingCatalogItem{utility.MixAndMatchItem}
	} else if !teamCompFound {
		err = &ErrMissingCatalogItem{utility.TeamCompItem}
	} else if !soloJazzFound {
		err = &ErrMissingCatalogItem{utility.SoloJazzItem}
	} else if !dancePassFound {
		err = &ErrMissingCatalogItem{utility.DancePassItem}
	} else if !weekendPassFound {
		err = &ErrMissingCatalogItem{utility.WeekendPassItem}
	} else if !tShirtFound {
		err = &ErrMissingCatalogItem{utility.TShirtItem}
	} else if !studentDiscountFound {
		err = &ErrMissingCatalogItem{utility.StudentDiscountItem}
	}
	if err != nil {
		s.logger.Error(err)
		return nil, err
	}

	s.logger.Tracef("Finished parsing square catalog list response")

	if len(tiers) != 5 {
		err := fmt.Errorf("Found unexpected number of full weekend tiers %v", len(tiers))
		s.logger.Error(err)
		return nil, err
	}

	weekendPassIds := []string{}
	for key, value := range tiers {
		weekendPassIds = append(weekendPassIds, key)
		if value.tier == 5 {
			res.WeekendPassTier = 5
			res.WeekendPassCost = value.cost
		}
	}
	counts := s.client.BatchRetrieveInventoryCounts(ctx, weekendPassIds, nil, nil)
	for counts.Next() {
		quantity, err := strconv.ParseFloat(counts.Value().Quantity, 64)
		if err != nil {
			s.logger.WithField("quantity", counts.Value().Quantity).Error("could not convert quantity to float")
			return nil, errors.Wrapf(err, "could not convert quantity %s to float", counts.Value().Quantity)
		}
		if quantity > 0 {
			thisTier := tiers[counts.Value().CatalogObjectId]
			if thisTier.tier < res.WeekendPassTier {
				res.WeekendPassTier = thisTier.tier
				res.WeekendPassCost = thisTier.cost
			}
		}
	}
	if counts.Error() != nil {
		wrap := "error retrieving inventory counts from square"
		utility.LogSquareError(s.logger, counts.Error(), wrap)
		return nil, errors.Wrap(counts.Error(), wrap)
	}
	return res, nil
}
