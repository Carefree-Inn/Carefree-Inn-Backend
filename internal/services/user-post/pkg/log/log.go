package log

import (
	"fmt"
	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"github.com/pkg/errors"
	"github.com/rifflock/lfshook"
	"github.com/sirupsen/logrus"
	"os"
	"time"
)

var logger *logrus.Logger

func NewLogger(dir string) {
	logger = logrus.New()
	
	logger.SetOutput(os.Stdout)
	logger.SetFormatter(&logrus.JSONFormatter{
		TimestampFormat: "2006-01-02 15:04:05",
		PrettyPrint:     true,
	})
	
	logger.AddHook(loggerHook(dir))
}

func WithField(key string, value interface{}) *logrus.Entry {
	return logrus.WithField(key, value)
}

func WithFields(fields logrus.Fields) *logrus.Entry {
	return logrus.WithFields(fields)
}

func loggerHook(dir string) *lfshook.LfsHook {
	return lfshook.NewHook(
		lfshook.WriterMap{
			logrus.InfoLevel:  fileDivisionByTime(dir + "Info"),
			logrus.DebugLevel: fileDivisionByTime(dir + "Debug"),
			logrus.WarnLevel:  fileDivisionByTime(dir + "Warn"),
			logrus.PanicLevel: fileDivisionByTime(dir + "Panic"),
			logrus.FatalLevel: fileDivisionByTime(dir + "Fatal"),
		}, &logrus.JSONFormatter{
			TimestampFormat: "2006-01-02 15:04:05",
			PrettyPrint:     true,
		})
}

func fileDivisionByTime(level string) *rotatelogs.RotateLogs {
	division, err := rotatelogs.New(
		level+"/%Y-%m-%d.log",
		rotatelogs.WithMaxAge(time.Hour*24*7),
		rotatelogs.WithRotationTime(time.Hour*24),
	)
	if err != nil {
		logrus.WithError(err).WithField("stack", fmt.Sprintf("%+v", errors.WithStack(err))).Fatal()
	}
	return division
}

func Trace(entry *logrus.Entry) {
	if entry == nil {
		entry = logrus.NewEntry(logger)
	}
	entry.Trace()
}

func Info(entry *logrus.Entry, message ...string) {
	if entry == nil {
		entry = logrus.NewEntry(logger)
	}
	entry.Infoln(message)
}

func Debug(entry *logrus.Entry, err error, message ...string) {
	if entry == nil {
		entry = logrus.NewEntry(logger)
	}
	if err != nil {
		entry.WithError(err).WithField("stack", fmt.Sprintf("%+v", err)).Debugln(message)
	} else {
		entry.Debugln(message)
	}
}

func Warn(entry *logrus.Entry, err error, message ...string) {
	if entry == nil {
		entry = logrus.NewEntry(logger)
	}
	entry.WithError(err).WithField("stack", fmt.Sprintf("%+v", err)).Warnln(message)
}

func Panic(entry *logrus.Entry, err error, message ...string) {
	if entry == nil {
		entry = logrus.NewEntry(logger)
	}
	entry.WithError(err).WithField("stack", fmt.Sprintf("%+v", err)).Panicln(message)
}

func Fatal(entry *logrus.Entry, err error, message ...string) {
	if entry == nil {
		entry = logrus.NewEntry(logger)
	}
	
	entry.WithError(err).WithField("stack", fmt.Sprintf("%+v", err)).Fatalln(message)
}
