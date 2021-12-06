package vaccine

import (
	"context"
	"fmt"

	"github.com/Houndie/dss-registration/dynamic/api"
	pb "github.com/Houndie/dss-registration/dynamic/rpc/dss"
	"github.com/Houndie/dss-registration/dynamic/vaccine"
	"github.com/twitchtv/twirp"
)

func (s *Server) Get(ctx context.Context, req *pb.VaccineGetReq) (*pb.VaccineGetRes, error) {
	auth, ok := api.GetAuth(ctx)
	if !ok {
		return nil, twirp.NewError(twirp.Unauthenticated, "unauthenticated")
	}

	info, err := s.service.Get(ctx, auth, req.Id)
	if err != nil {
		return nil, err
	}

	res := &pb.VaccineGetRes{}
	switch t := info.(type) {
	case *vaccine.VaxApproved:
		res.Info = &pb.VaccineGetRes_VaxApproved{VaxApproved: &pb.VaxApproved{}}
	case *vaccine.VaxApprovalPending:
		res.Info = &pb.VaccineGetRes_VaxApprovalPending{
			VaxApprovalPending: &pb.VaxApprovalPending{
				Url: t.URL,
			},
		}
	case *vaccine.NoVaxProofSupplied:
		res.Info = &pb.VaccineGetRes_NoVaxProofSupplied{NoVaxProofSupplied: &pb.NoVaxProofSupplied{}}
	default:
		return nil, fmt.Errorf("unknown vaccine info supplied")
	}

	return res, nil
}
