package golang_logging

import (
	"testing"

	"github.com/sirupsen/logrus"
)

func TestField(t *testing.T) {
	logger := logrus.New()
	logger.SetFormatter(&logrus.JSONFormatter{})
	logger.WithField("username", "syauqi").Info("Hello World")

	logger.WithField("username", "syauqi").WithField("password", "syauqi").Info("Hello World")
}

func TestFields(t *testing.T) {
	logger := logrus.New()
	logger.SetFormatter(&logrus.JSONFormatter{})
	logger.WithFields(logrus.Fields{
		"username":"syauqi",
		"name": "Syauqi Damario Djohan",
	}).Infof("Hello World")
}