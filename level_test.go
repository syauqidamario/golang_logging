package golang_logging

import (
	"testing"

	"github.com/sirupsen/logrus"
)

func TestLevel(t *testing.T) {
	logger := logrus.New()
	logger.Trace("Info")
	logger.Debug("Info")
	logger.Info("Info")
	logger.Warn("Info")
	logger.Error("Info")
}

func TestLoggerLevel(t *testing.T){
	logger := logrus.New()
	logger.SetLevel(logrus.TraceLevel)
	logger.Trace("Info")
	logger.Debug("Info")
	logger.Info("Info")
	logger.Warn("Info")
	logger.Error("Info")
}