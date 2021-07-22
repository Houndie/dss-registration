package registration

import (
	"context"
	"fmt"

	"github.com/Houndie/dss-registration/dynamic/common"
	"github.com/Houndie/dss-registration/dynamic/storage"
	"github.com/Houndie/dss-registration/dynamic/utility"
)

func (s *Service) Populate(ctx context.Context) (storage.WeekendPassTier, error) {
	bestTier, _, err := common.LowestInStockTier(ctx, s.squareData.PurchaseItems.FullWeekend, s.client)
	if err != nil {
		utility.LogSquareError(s.logger, err, "error finding best tier and cost")
		return 0, fmt.Errorf("error finding best tier and cost: %w", err)
	}

	return bestTier, nil
}
