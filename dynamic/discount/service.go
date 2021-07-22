package discount

import (
	"github.com/Houndie/dss-registration/dynamic/common"
	"github.com/sirupsen/logrus"
)

func NewService(logger *logrus.Logger, squareData *common.SquareData) *Service {
	return &Service{
		logger:     logger,
		squareData: squareData,
	}
}

type Service struct {
	logger     *logrus.Logger
	squareData *common.SquareData
}
