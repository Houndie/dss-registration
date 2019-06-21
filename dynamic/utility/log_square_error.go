package utility

import (
	"github.com/Houndie/dss-registration/dynamic/square"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
)

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
