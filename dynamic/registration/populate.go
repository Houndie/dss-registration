package registration

import (
	"fmt"

	"github.com/Houndie/dss-registration/dynamic/square"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
)

type PopulateData struct {
	WeekendPassCost int
	WeekendPassTier int
	DancePassCost   int
	MixAndMatchCost int
	SoloJazzCost    int
	TeamCompCost    int
	TShirtCost      int
}

func LogSquareError(logger *logrus.Logger, err error, message string) {
	switch e := errors.Cause(err).(type) {
	case *square.Error:
		logger.WithFields(logrus.Fields{
			"Category": e.Category,
			"Code":     e.Code,
			"Detail":   e.Detail,
			"Field":    e.Field,
		}).Error(message)
	case *square.ErrorList:
		for _, squareError := range e.Errors {
			logger.WithFields(logrus.Fields{
				"Category": squareError.Category,
				"Code":     squareError.Code,
				"Detail":   squareError.Detail,
				"Field":    squareError.Field,
			}).Error(message)
		}
	default:
		logger.WithError(err).Error(message)
	}
}

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

func Populate(client *square.Client, logger *logrus.Logger) (PopulateData, error) {
	logger.Trace("Fetching all items from square")
	objects, err := client.ListCatalog(nil)
	if err != nil {
		wrap := "error fetching all items from square"
		LogSquareError(logger, err, wrap)
		return PopulateData{}, errors.Wrap(err, wrap)
	}
	logger.Trace("Iterating over square list catalog responses")
	res := PopulateData{}
	tiers := map[string]tierData{}
	for _, object := range objects {
		logger.Trace("Found square object")
		item, ok := object.CatalogObjectType.(*square.CatalogItem)
		if !ok {
			logger.Trace("Square object was not of type catalog item")
			continue
		}
		if item == nil {
			logger.Tracef("here")
		}
		logger.Tracef("Comparing item name %s to legend", item.Name)
		switch item.Name {
		case mixAndMatchItem, teamCompItem, soloJazzItem, tShirtItem:
			logger.Trace("Found competition item")
			if len(item.Variations) != 1 {
				err := fmt.Errorf("Found unexpected number of variations: %v", len(item.Variations))
				logger.Error(err)
				return PopulateData{}, err
			}
			variation, ok := item.Variations[0].CatalogObjectType.(*square.CatalogItemVariation)
			if !ok {
				err := "Invalid response from square...item variation isn't a variation?"
				logger.Error(err)
				return PopulateData{}, errors.New(err)
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
				logger.Error(err)
				return PopulateData{}, err
			}
		case dancePassItem:
			logger.Trace("Found dance pass item")
			for _, v := range item.Variations {
				variation, ok := v.CatalogObjectType.(*square.CatalogItemVariation)
				if !ok {
					err := "Invalid response from square...item variation isn't a variation?"
					logger.Error(err)
					return PopulateData{}, errors.New(err)
				}
				if variation.Name == "Presale" {
					logger.Trace("Found dance pass variant Presale")
					res.DancePassCost = variation.PriceMoney.Amount
					break
				}
				logger.Tracef("Did not find dance pass variant Presale (found %s), moving on", variation.Name)
			}
		case weekendPassItem:
			logger.Trace("Found weekend pass object")
			for _, v := range item.Variations {
				variation, ok := v.CatalogObjectType.(*square.CatalogItemVariation)
				if !ok {
					err := "Invalid response from square...item variation isn't a variation?"
					logger.Error(err)
					return PopulateData{}, errors.New(err)
				}

				logger.Tracef("Found variation with name %s and id %s", variation.Name, v.Id)
				switch variation.Name {
				case weekendPassTier1Name:
					logger.Trace("Variation matched tier 1")
					tiers[v.Id] = tierData{1, variation.PriceMoney.Amount}
				case weekendPassTier2Name:
					logger.Trace("Variation matched tier 2")
					tiers[v.Id] = tierData{2, variation.PriceMoney.Amount}
				case weekendPassTier3Name:
					logger.Trace("Variation matched tier 3")
					tiers[v.Id] = tierData{3, variation.PriceMoney.Amount}
				case weekendPassTier4Name:
					logger.Trace("Variation matched tier 4")
					tiers[v.Id] = tierData{4, variation.PriceMoney.Amount}
				case weekendPassTier5Name:
					logger.Trace("Variation matched tier 5")
					tiers[v.Id] = tierData{5, variation.PriceMoney.Amount}
				default: // Do nothing, we have other names that are allowable
					logger.Trace("Variation did not match tier list, moving on")
				}
			}
		}
	}
	logger.Tracef("Finished parsing square catalog list response")

	if len(tiers) != 5 {
		err := fmt.Errorf("Found unexpected number of full weekend tiers %v", len(tiers))
		logger.Error(err)
		return PopulateData{}, err
	}

	weekendPassIds := []string{}
	for key, value := range tiers {
		weekendPassIds = append(weekendPassIds, key)
		if value.tier == 5 {
			res.WeekendPassTier = 5
			res.WeekendPassCost = value.cost
		}
	}
	counts, err := client.BatchRetrieveInventoryCounts(weekendPassIds, nil, nil)
	if err != nil {
		wrap := "error retrieving inventory counts from square"
		LogSquareError(logger, err, wrap)
		return PopulateData{}, errors.Wrap(err, wrap)
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
