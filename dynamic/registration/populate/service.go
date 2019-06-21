package populate

import (
	"fmt"
	"time"

	"github.com/Houndie/dss-registration/dynamic/square"
	"github.com/Houndie/dss-registration/dynamic/utility"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
)

const (
	mixAndMatchItem = "Mix And Match"
	teamCompItem    = "Team Competition"
	soloJazzItem    = "Solo"
	dancePassItem   = "Dance Only"
	weekendPassItem = "Full Weekend Pass"
	tShirtItem      = "2020 T-Shirt"

	weekendPassTier1Name = "Tier 1"
	weekendPassTier2Name = "Tier 2"
	weekendPassTier3Name = "Tier 3"
	weekendPassTier4Name = "Tier 4"
	weekendPassTier5Name = "Tier 5"
)

type tierData struct {
	tier int
	cost int
}

type SquareClient interface {
	ListCatalog([]string) ([]*square.CatalogObject, error)
	BatchRetrieveInventoryCounts([]string, []string, *time.Time) ([]*square.InventoryCount, error)
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

func (s *Service) Populate() (*FormData, error) {
	s.logger.Trace("Fetching all items from square")
	objects, err := s.client.ListCatalog(nil)
	if err != nil {
		wrap := "error fetching all items from square"
		utility.LogSquareError(s.logger, err, wrap)
		return nil, errors.Wrap(err, wrap)
	}
	s.logger.Trace("Iterating over square list catalog responses")
	res := &FormData{}
	tiers := map[string]tierData{}
	for _, object := range objects {
		s.logger.Trace("Found square object")
		item, ok := object.CatalogObjectType.(*square.CatalogItem)
		if !ok {
			s.logger.Trace("Square object was not of type catalog item")
			continue
		}
		if item == nil {
			s.logger.Tracef("here")
		}
		s.logger.Tracef("Comparing item name %s to legend", item.Name)
		switch item.Name {
		case mixAndMatchItem, teamCompItem, soloJazzItem, tShirtItem:
			s.logger.Trace("Found competition item")
			if len(item.Variations) != 1 {
				err := fmt.Errorf("Found unexpected number of variations: %v", len(item.Variations))
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
			case mixAndMatchItem:
				res.MixAndMatchCost = variation.PriceMoney.Amount
			case teamCompItem:
				res.TeamCompCost = variation.PriceMoney.Amount
			case soloJazzItem:
				res.SoloJazzCost = variation.PriceMoney.Amount
			case tShirtItem:
				res.TShirtCost = variation.PriceMoney.Amount
			default:
				err := errors.New("Impossible code path...how did I get here")
				s.logger.Error(err)
				return nil, err
			}
		case dancePassItem:
			s.logger.Trace("Found dance pass item")
			for _, v := range item.Variations {
				variation, ok := v.CatalogObjectType.(*square.CatalogItemVariation)
				if !ok {
					err := "Invalid response from square...item variation isn't a variation?"
					s.logger.Error(err)
					return nil, errors.New(err)
				}
				if variation.Name == "Presale" {
					s.logger.Trace("Found dance pass variant Presale")
					res.DancePassCost = variation.PriceMoney.Amount
					break
				}
				s.logger.Tracef("Did not find dance pass variant Presale (found %s), moving on", variation.Name)
			}
		case weekendPassItem:
			s.logger.Trace("Found weekend pass object")
			for _, v := range item.Variations {
				variation, ok := v.CatalogObjectType.(*square.CatalogItemVariation)
				if !ok {
					err := "Invalid response from square...item variation isn't a variation?"
					s.logger.Error(err)
					return nil, errors.New(err)
				}

				s.logger.Tracef("Found variation with name %s and id %s", variation.Name, v.Id)
				switch variation.Name {
				case weekendPassTier1Name:
					s.logger.Trace("Variation matched tier 1")
					tiers[v.Id] = tierData{1, variation.PriceMoney.Amount}
				case weekendPassTier2Name:
					s.logger.Trace("Variation matched tier 2")
					tiers[v.Id] = tierData{2, variation.PriceMoney.Amount}
				case weekendPassTier3Name:
					s.logger.Trace("Variation matched tier 3")
					tiers[v.Id] = tierData{3, variation.PriceMoney.Amount}
				case weekendPassTier4Name:
					s.logger.Trace("Variation matched tier 4")
					tiers[v.Id] = tierData{4, variation.PriceMoney.Amount}
				case weekendPassTier5Name:
					s.logger.Trace("Variation matched tier 5")
					tiers[v.Id] = tierData{5, variation.PriceMoney.Amount}
				default: // Do nothing, we have other names that are allowable
					s.logger.Trace("Variation did not match tier list, moving on")
				}
			}
		}
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
	counts, err := s.client.BatchRetrieveInventoryCounts(weekendPassIds, nil, nil)
	if err != nil {
		wrap := "error retrieving inventory counts from square"
		utility.LogSquareError(s.logger, err, wrap)
		return nil, errors.Wrap(err, wrap)
	}
	for _, count := range counts {
		// argh, comparing float with equals
		if count.Quantity != 0 {
			thisTier := tiers[count.CatalogObjectId]
			if thisTier.tier < res.WeekendPassTier {
				res.WeekendPassTier = thisTier.tier
				res.WeekendPassCost = thisTier.cost
			}
		}
	}
	return res, nil
}
