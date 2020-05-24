package discount

import (
	"fmt"

	"github.com/Houndie/dss-registration/dynamic/common"
	"github.com/Houndie/dss-registration/dynamic/discount"
	pb "github.com/Houndie/dss-registration/dynamic/rpc/dss"
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

type ErrUnknownDiscountType struct {
	T string
}

func (e ErrUnknownDiscountType) Error() string {
	return fmt.Sprintf("unknown discount type %s", e.T)
}

func amountFromProto(d *pb.DiscountAmount) (common.DiscountAmount, error) {
	switch t := d.Amount.(type) {
	case *pb.DiscountAmount_Dollar:
		return common.DollarDiscount(t.Dollar), nil
	case *pb.DiscountAmount_Percent:
		return common.PercentDiscount(t.Percent), nil
	default:
		return nil, ErrUnknownDiscountType{T: fmt.Sprintf("%T", t)}
	}
}

func bundleToProto(d *discount.Bundle) (*pb.DiscountBundle, error) {
	singleDiscounts := make([]*pb.SingleDiscount, len(d.Discounts))
	for i, discount := range d.Discounts {
		amt, err := AmountToProto(discount.Amount)
		if err != nil {
			return nil, err
		}
		singleDiscounts[i] = &pb.SingleDiscount{
			Name:   discount.Name,
			Amount: amt,
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
		singleDiscounts[i] = &discount.Single{
			Name:   sd.Name,
			Amount: amt,
		}
	}
	return &discount.Bundle{
		Code:      d.Code,
		Discounts: singleDiscounts,
	}, nil

}
