package registration

import (
	"context"
	"errors"
	"fmt"

	"github.com/Houndie/dss-registration/dynamic/common"
	"github.com/Houndie/dss-registration/dynamic/discount"
	"github.com/Houndie/dss-registration/dynamic/storage"
	"github.com/Houndie/dss-registration/dynamic/utility"
)

type FormData struct {
	WeekendPassCost int
	WeekendPassTier storage.WeekendPassTier
	DancePassCost   int
	MixAndMatchCost int
	SoloJazzCost    int
	TeamCompCost    int
	TShirtCost      int
	StudentDiscount discount.DiscountAmount
}

func (s *Service) Populate(ctx context.Context) (*FormData, error) {
	s.logger.Trace("Fetching all items from square")
	catalogData, err := common.GetSquareCatalog(ctx, s.client)
	if err != nil {
		utility.LogSquareError(s.logger, err, "error fetching square catalog data")
		return nil, fmt.Errorf("error fetching square catalog data: %w", err)
	}

	s.logger.Tracef("Finished parsing square catalog list response")

	bestTier, bestCost, err := common.LowestInStockTier(ctx, catalogData, s.client)
	if err != nil {
		utility.LogSquareError(s.logger, err, "error finding best tier and cost")
		return nil, fmt.Errorf("error finding best tier and cost: %w", err)
	}

	var studentDiscount discount.DiscountAmount
	switch sd := catalogData.StudentDiscount.Amount.(type) {
	case common.DollarDiscount:
		studentDiscount = discount.DollarDiscount(int(sd))
	case common.PercentDiscount:
		studentDiscount = discount.PercentDiscount(string(sd))
	default:
		return nil, errors.New("unknown discount type from square data")
	}

	return &FormData{
		WeekendPassTier: bestTier,
		WeekendPassCost: bestCost,
		DancePassCost:   catalogData.DanceOnly.Cost,
		SoloJazzCost:    catalogData.SoloJazz.Cost,
		MixAndMatchCost: catalogData.MixAndMatch.Cost,
		TeamCompCost:    catalogData.TeamCompetition.Cost,
		TShirtCost:      catalogData.TShirt.Cost,
		StudentDiscount: studentDiscount,
	}, nil
}
