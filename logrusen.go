package logrusen

import (
	"errors"
	"fmt"
	"os"
	"runtime"

	"github.com/evalphobia/logrus_sentry"
	"github.com/rifflock/lfshook"
	"github.com/sirupsen/logrus"
	log "github.com/sirupsen/logrus"
)

type StandardLogger interface {
	Trace(event, topic, description string)
	Debug(event, topic, description string)
	Info(event, topic, description string)
	Warn(event, topic, description string)
	Error(event, topic, description, errorMessage string)
	Fatal(event, topic, description, errorMessage string)
	Panic(event, topic, description, errorMessage string)

	Setup(env, dsn string) (*standardLogger, error)
}

type standardLogger struct {
	*logrus.Logger
}

func New() StandardLogger {
	baseLogger := logrus.New()
	standardLogger := &standardLogger{baseLogger}
	return standardLogger
}

func (l *standardLogger) Setup(env, dsn string) (*standardLogger, error) {
	switch env {
	case "prod":
		err := setProduction(dsn)
		if err != nil {
			return nil, err
		}
	case "dev":
		setDevelopment()
	case "":
		return nil, errors.New("env variable is nil (must: dev/prod)")
	default:
		return nil, errors.New("invalid env variable (must: dev/prod)")
	}
	return nil, nil
}

func setDevelopment() {
	log.SetOutput(os.Stdout)

	log.SetFormatter(&log.TextFormatter{
		FullTimestamp: true,
		ForceColors:   true,
	})
}

func setProduction(dsn string) error {
	log.SetOutput(os.Stdout)

	log.SetFormatter(&log.TextFormatter{
		FullTimestamp: true,
		ForceColors:   true,
	})

	log.AddHook(lfshook.NewHook(
		lfshook.PathMap{
			logrus.InfoLevel:  "logs.log",
			logrus.ErrorLevel: "logs.log",
			logrus.DebugLevel: "logs.log",
			logrus.FatalLevel: "logs.log",
			logrus.PanicLevel: "logs.log",
			logrus.TraceLevel: "logs.log",
			logrus.WarnLevel:  "logs.log",
		},
		&log.JSONFormatter{},
	))

	sentryHook, err := logrus_sentry.NewSentryHook(dsn, []log.Level{
		logrus.PanicLevel,
		logrus.FatalLevel,
		logrus.ErrorLevel,
	})
	if err != nil {
		return err
	}

	log.AddHook(sentryHook)
	return nil
}

func stdFields(event, topic string) *log.Fields {
	return &log.Fields{
		"event": event,
		"topic": topic,
	}
}

func errFields(event, topic, description string) *log.Fields {
	alloc, totalAlloc, sys, numGC := PrintMemUsage()
	pc, _, line, _ := runtime.Caller(2)
	return &log.Fields{
		"description":   description,
		"event":         event,
		"topic":         topic,
		"caller":        fmt.Sprintf("%s:%d", runtime.FuncForPC(pc).Name(), line),
		"memAlloc":      alloc,
		"totalMemAlloc": totalAlloc,
		"sysMem":        sys,
		"numGC":         numGC,
		"numCPU":        runtime.NumCPU(),
		"numGoroutine":  runtime.NumGoroutine(),
	}
}

func (l *standardLogger) Trace(event, topic, description string) {
	fields := stdFields(event, topic)

	log.WithFields(*fields).Trace(description)
}

func (l *standardLogger) Debug(event, topic, description string) {
	fields := stdFields(event, topic)

	log.WithFields(*fields).Debug(description)
}

func (l *standardLogger) Info(event, topic, description string) {
	fields := stdFields(event, topic)

	log.WithFields(*fields).Info(description)
}

func (l *standardLogger) Warn(event, topic, description string) {
	fields := stdFields(event, topic)

	log.WithFields(*fields).Warn(description)
}

func (l *standardLogger) Error(event, topic, description, errorMessage string) {
	fields := errFields(event, topic, description)

	log.WithFields(*fields).Error(errorMessage)
}

func (l *standardLogger) Fatal(event, topic, description, errorMessage string) {
	fields := errFields(event, topic, description)

	log.WithFields(*fields).Fatal(errorMessage)
}

func (l *standardLogger) Panic(event, topic, description, errorMessage string) {
	fields := errFields(event, topic, description)

	log.WithFields(*fields).Panic(errorMessage)
}

func PrintMemUsage() (alloc, totalAlloc, sys, numGC uint64) {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	return bToMb(m.Alloc), bToMb(m.TotalAlloc), bToMb(m.Sys), uint64(m.NumGC)
}

func bToMb(b uint64) uint64 {
	return b / 1024 / 1024
}
