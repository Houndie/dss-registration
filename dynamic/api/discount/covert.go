package discount

import (
	"fmt"

	"github.com/Houndie/dss-registration/dynamic/discount"
	pb "github.com/Houndie/dss-registration/dynamic/rpc/dss"
	"github.com/Houndie/dss-registration/dynamic/storage"
	"github.com/golang/protobuf/ptypes/empty"
)

func AmountToProto(d discount.DiscountAmount) (*pb.DiscountAmount, error) {
	switch t := d.(type) {
	case discount.PercentDiscount:
		return &pb.DiscountAmount{
			Amount: &pb.DiscountAmount_Percent{
				Percent: string(t),
			},
		}, nil
	case discount.DollarDiscount:
		return &pb.DiscountAmount{
			Amount: &pb.DiscountAmount_Dollar{
				Dollar: int64(t),
			},
		}, nil
	case discount.SquareNotFound:
		return &pb.DiscountAmount{
			Amount: &pb.DiscountAmount_SquareNotFound{
				SquareNotFound: &empty.Empty{},
			},
		}, nil
	default:
		return nil, fmt.Errorf("unknown discount type %T", t)
	}
}

type ErrUnknownDiscountType struct {
	T string
}

func (e ErrUnknownDiscountType) Error() string {
	return fmt.Sprintf("unknown discount type %s", e.T)
}

func amountFromProto(d *pb.DiscountAmount) (discount.DiscountAmount, error) {
	if d == nil {
		return nil, nil
	}
	switch t := d.Amount.(type) {
	case *pb.DiscountAmount_Dollar:
		return discount.DollarDiscount(t.Dollar), nil
	case *pb.DiscountAmount_Percent:
		return discount.PercentDiscount(t.Percent), nil
	case *pb.DiscountAmount_SquareNotFound:
		return discount.SquareNotFound{}, nil
	default:
		return nil, ErrUnknownDiscountType{T: fmt.Sprintf("%T", t)}
	}
}

type ErrUnknownPurchaseItem struct {
	AppliedTo pb.PurchaseItem
}

func (e ErrUnknownPurchaseItem) Error() string {
	return fmt.Sprintf("unknown purchase item %s", e.AppliedTo)
}

func appliedToFromProto(appliedTo pb.PurchaseItem) (storage.PurchaseItem, error) {
	switch appliedTo {
	case pb.PurchaseItem_FullWeekendPassPurchaseItem:
		return storage.FullWeekendPurchaseItem, nil
	case pb.PurchaseItem_DanceOnlyPassPurchaseItem:
		return storage.DanceOnlyPurchaseItem, nil
	case pb.PurchaseItem_MixAndMatchPurchaseItem:
		return storage.MixAndMatchPurchaseItem, nil
	case pb.PurchaseItem_SoloJazzPurchaseItem:
		return storage.SoloJazzPurchaseItem, nil
	case pb.PurchaseItem_TeamCompetitionPurchaseItem:
		return storage.TeamCompetitionPurchaseItem, nil
	case pb.PurchaseItem_TShirtPurchaseItem:
		return storage.TShirtPurchaseItem, nil
	default:
		return "", ErrUnknownPurchaseItem{AppliedTo: appliedTo}
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

func bundleToProto(d *discount.Bundle) (*pb.DiscountBundle, error) {
	singleDiscounts := make([]*pb.SingleDiscount, len(d.Discounts))
	for i, discount := range d.Discounts {
		amt, err := AmountToProto(discount.Amount)
		if err != nil {
			return nil, err
		}
		appliedTo, err := appliedToToProto(discount.AppliedTo)
		if err != nil {
			return nil, err
		}
		singleDiscounts[i] = &pb.SingleDiscount{
			Name:      discount.Name,
			Amount:    amt,
			AppliedTo: appliedTo,
		}
	}
	return &pb.DiscountBundle{
		Code:      d.Code,
		Discounts: singleDiscounts,
	}, nil
}

func bundleFromProto(d *pb.DiscountBundle) (*discount.Bundle, error) {
	singleDiscounts := make([]*discount.Single, len(d.Discounts))
	for i, sd := range d.Discounts {
		amt, err := amountFromProto(sd.Amount)
		if err != nil {
			return nil, err
		}
		appliedTo, err := appliedToFromProto(sd.AppliedTo)
		if err != nil {
			return nil, err
		}
		singleDiscounts[i] = &discount.Single{
			Name:      sd.Name,
			Amount:    amt,
			AppliedTo: appliedTo,
		}
	}
	return &discount.Bundle{
		Code:      d.Code,
		Discounts: singleDiscounts,
	}, nil

}
