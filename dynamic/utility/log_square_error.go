package utility

import (
	"errors"

	"github.com/Houndie/dss-registration/dynamic/square"
	"github.com/sirupsen/logrus"
)

func LogSquareError(logger *logrus.Logger, err error, message string) {
	serr := &square.Error{}
	slerr := &square.ErrorList{}
	if errors.As(err, &serr) {
		logger.WithFields(logrus.Fields{
			"Category": serr.Category,
			"Code":     serr.Code,
			"Detail":   serr.Detail,
			"Field":    serr.Field,
		}).Error(message)
	} else if errors.As(err, &slerr) {
		for _, squareError := range slerr.Errors {
			logger.WithFields(logrus.Fields{
				"Category": squareError.Category,
				"Code":     squareError.Code,
				"Detail":   squareError.Detail,
				"Field":    squareError.Field,
			}).Error(message)
		}
	} else {
		logger.WithError(err).Error(message)
	}
}
