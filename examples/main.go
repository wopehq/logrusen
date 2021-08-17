package main

import (
	"github.com/teamseodo/logrusen"

	"github.com/sirupsen/logrus"
)

func main() {
	logger := logrusen.New()
	err := logger.Setup()

	if err != nil {
		logger.Error("Logger initialization error", err, nil)
	}
	logger.Info("Info testing", nil)
	logger.Debug("Debug testing", logrus.Fields{"test": "testvalue"})
	logger.Warn("Warn logging", err, nil)
	logger.Fatal("Fatal logging", err, nil)
}
