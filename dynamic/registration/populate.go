package registration

import (
	"context"
	"fmt"

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
	StudentDiscount DiscountAmount
}

func (s *Service) Populate(ctx context.Context) (*FormData, error) {
	s.logger.Trace("Fetching all items from square")
	catalogData, err := getSquareCatalog(ctx, s.client)
	if err != nil {
		utility.LogSquareError(s.logger, err, "error fetching square catalog data")
		return nil, fmt.Errorf("error fetching square catalog data: %w", err)
	}

	s.logger.Tracef("Finished parsing square catalog list response")

	bestTier, bestCost, err := lowestInStockTier(ctx, catalogData, s.client)
	if err != nil {
		utility.LogSquareError(s.logger, err, "error finding best tier and cost")
		return nil, fmt.Errorf("error finding best tier and cost: %w", err)
	}

	return &FormData{
		WeekendPassTier: bestTier,
		WeekendPassCost: bestCost,
		DancePassCost:   catalogData.danceOnly.cost,
		SoloJazzCost:    catalogData.soloJazz.cost,
		MixAndMatchCost: catalogData.mixAndMatch.cost,
		TeamCompCost:    catalogData.teamCompetition.cost,
		TShirtCost:      catalogData.tShirt.cost,
		StudentDiscount: catalogData.studentDiscount.amount,
	}, nil
}
