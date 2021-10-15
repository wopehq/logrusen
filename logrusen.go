package logrusen

import (
	"fmt"
	"os"
	"runtime"

	"github.com/evalphobia/logrus_sentry"
	"github.com/rifflock/lfshook"
	log "github.com/sirupsen/logrus"
)

type StandardLogger interface {
	Debug(message string, fields Fields)
	Info(message string, fields Fields)
	Warn(message string, err error, fields Fields)
	Error(message string, err error, fields Fields)
	Fatal(message string, err error, fields Fields)
	Panic(message string, err error, fields Fields)

	Setup() error
	SetupWithSentry(dsn string) error
}

type standardLogger struct {
	*log.Logger
}

type Fields log.Fields

func New() StandardLogger {
	baseLogger := log.New()
	standardLogger := &standardLogger{baseLogger}
	return standardLogger
}

func (l *standardLogger) Setup() error {
	setDefault()
	return nil
}

func (l *standardLogger) SetupWithSentry(dsn string) error {
	if dsn != "" {
		err := setDefaultWithSentry(dsn)
		if err != nil {
			return err
		}
	} else {
		setDefault()
	}
	return nil
}

func setDefault() {
	log.SetOutput(os.Stdout)

	log.SetLevel(log.DebugLevel)

	log.SetFormatter(&log.TextFormatter{
		FullTimestamp: true,
		ForceColors:   true,
	})
}

func setDefaultWithSentry(dsn string) error {
	log.SetOutput(os.Stdout)

	log.SetLevel(log.DebugLevel)

	log.SetFormatter(&log.TextFormatter{
		FullTimestamp: true,
		ForceColors:   true,
	})

	log.AddHook(lfshook.NewHook(
		lfshook.PathMap{
			log.InfoLevel:  "logs.log",
			log.ErrorLevel: "logs.log",
			log.DebugLevel: "logs.log",
			log.FatalLevel: "logs.log",
			log.PanicLevel: "logs.log",
			log.TraceLevel: "logs.log",
			log.WarnLevel:  "logs.log",
		},
		&log.JSONFormatter{},
	))

	sentryHook, err := logrus_sentry.NewSentryHook(dsn, []log.Level{
		log.PanicLevel,
		log.FatalLevel,
		log.ErrorLevel,
	})
	if err != nil {
		return err
	}

	log.AddHook(sentryHook)
	return nil
}

// Const Log Variables
func constFields(fields Fields) Fields {
	if fields == nil {
		fields = Fields{}
	}
	pc, _, line, _ := runtime.Caller(2)
	fields["caller"] = fmt.Sprintf("%s:%d", runtime.FuncForPC(pc).Name(), line)
	return fields
}

// Error Log Variables
func errorFields(fields Fields) Fields {
	if fields == nil {
		fields = Fields{}
	}

	pc, _, line, _ := runtime.Caller(2)
	fields["caller"] = fmt.Sprintf("%s:%d", runtime.FuncForPC(pc).Name(), line)
	return fields
}

// DEBUG:
// message* user friendly error messega
// fields* can be nil or can be env and system status variables
func (l *standardLogger) Debug(message string, fields Fields) {
	fields = errorFields(fields)

	log.WithFields(log.Fields(fields)).Debug(message)
}

// INFO:
// message* user friendly error messega
// fields* can be nil or can be env and system status variables
func (l *standardLogger) Info(message string, fields Fields) {
	fields = constFields(fields)

	log.WithFields(log.Fields(fields)).Info(message)
}

// Warn:
// message* user friendly error message
// err (error): An error obtained from a failed call to a previous method or function
// fields* can be nil or can be env and system status variables
func (l *standardLogger) Warn(message string, err error, fields Fields) {
	fields = errorFields(fields)
	fields["error"] = err

	log.WithFields(log.Fields(fields)).Warn(message)
}

// Error writes a message to the log of Error level status.
// message* user friendly error message
// err (error): An error obtained from a failed call to a previous method or function
// fields* can be nil or can be env and system status variables
func (l *standardLogger) Error(message string, err error, fields Fields) {
	fields = errorFields(fields)
	fields["error"] = err

	log.WithFields(log.Fields(fields)).Error(message)
}

// Fatal writes a message to the log of Fatal level status.
// Note: Calling a Fatal() error will exit execution of the current program. Goroutines will not
// execute on deferral. Only call Fatal() if you are sure that the program should exit as well.
func (l *standardLogger) Fatal(message string, err error, fields Fields) {
	fields = errorFields(fields)
	fields["error"] = err

	log.WithFields(log.Fields(fields)).Fatal(message)
}

// DONT PANIC
// so, i use fatal function
func (l *standardLogger) Panic(message string, err error, fields Fields) {
	l.Fatal(message, err, fields)
}
