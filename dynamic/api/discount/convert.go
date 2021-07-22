package discount

import (
	"fmt"

	"github.com/Houndie/dss-registration/dynamic/common"
	pb "github.com/Houndie/dss-registration/dynamic/rpc/dss"
	"github.com/Houndie/dss-registration/dynamic/storage"
)

func AmountToProto(d common.DiscountAmount) (*pb.DiscountAmount, error) {
	switch t := d.(type) {
	case common.PercentDiscount:
		return &pb.DiscountAmount{
			Amount: &pb.DiscountAmount_Percent{
				Percent: string(t),
			},
		}, nil
	case common.DollarDiscount:
		return &pb.DiscountAmount{
			Amount: &pb.DiscountAmount_Dollar{
				Dollar: int64(t),
			},
		}, nil
	default:
		return nil, fmt.Errorf("unknown discount type %T", t)
	}
}

func appliedToToProto(appliedTo storage.PurchaseItem) (pb.PurchaseItem, error) {
	switch appliedTo {
	case storage.FullWeekendPurchaseItem:
		return pb.PurchaseItem_FullWeekendPassPurchaseItem, nil
	case storage.DanceOnlyPurchaseItem:
		return pb.PurchaseItem_DanceOnlyPassPurchaseItem, nil
	case storage.MixAndMatchPurchaseItem:
		return pb.PurchaseItem_MixAndMatchPurchaseItem, nil
	case storage.SoloJazzPurchaseItem:
		return pb.PurchaseItem_SoloJazzPurchaseItem, nil
	case storage.TeamCompetitionPurchaseItem:
		return pb.PurchaseItem_TeamCompetitionPurchaseItem, nil
	case storage.TShirtPurchaseItem:
		return pb.PurchaseItem_TShirtPurchaseItem, nil
	default:
		return 0, fmt.Errorf("Unknown purchase item: %v", appliedTo)
	}
}

func bundleToProto(code string, d []*common.Discount) (*pb.DiscountBundle, error) {
	singleDiscounts := make([]*pb.SingleDiscount, len(d))
	for i, discount := range d {
		amt, err := AmountToProto(discount.Amount)
		if err != nil {
			return nil, err
		}
		appliedTo, err := appliedToToProto(discount.AppliedTo)
		if err != nil {
			return nil, err
		}
		singleDiscounts[i] = &pb.SingleDiscount{
			Amount:    amt,
			AppliedTo: appliedTo,
		}
	}
	return &pb.DiscountBundle{
		Code:      code,
		Discounts: singleDiscounts,
	}, nil
}
