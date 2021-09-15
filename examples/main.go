package main

import (
	"github.com/teamseodo/logrusen"
)

func main() {
	logger := logrusen.New()
	err := logger.Setup()

	if err != nil {
		logger.Error("Logger initialization error", err, nil)
	}
	logger.Info("Info testing", nil)
	logger.Debug("Debug testing", logrusen.Fields{"test": "testvalue"})
	logger.Warn("Warn logging", err, nil)
	logger.Fatal("Fatal logging", err, nil)
}
