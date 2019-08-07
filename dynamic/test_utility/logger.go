package test_utility

import (
	"strings"
	"testing"

	"github.com/sirupsen/logrus"
)

type ErrorWriter struct {
	T *testing.T
}

func (e *ErrorWriter) Write(b []byte) (int, error) {
	e.T.Log(strings.TrimSuffix(string(b), "\n"))
	return len(b), nil
}

type ErrorHook struct {
	T *testing.T
}

func (*ErrorHook) Levels() []logrus.Level {
	return []logrus.Level{
		logrus.PanicLevel,
		logrus.FatalLevel,
		logrus.ErrorLevel,
		logrus.WarnLevel,
	}
}

func (h *ErrorHook) Fire(e *logrus.Entry) error {
	msg, err := e.Logger.Formatter.Format(e)
	if err != nil {
		h.T.Errorf("Found logged messsage, unable to format message")
	}
	h.T.Errorf("Found logged message: %s", string(msg))
	return nil
}
