package log

import (
	"fmt"
	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"github.com/pkg/errors"
	"github.com/rifflock/lfshook"
	"github.com/sirupsen/logrus"
	"io"
	"os"
	"time"
)

var logger *logrus.Logger
var syncLog *SyncLog

func NewLogger(buffer uint32) {
	logger = logrus.New()
	syncLog = NewSyncLog(buffer)

	logger.SetOutput(os.Stdout)
	logger.SetFormatter(&logrus.JSONFormatter{
		TimestampFormat: "2006-01-02 15:04:05",
		PrettyPrint:     true,
	})

	logger.AddHook(loggerHook())
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

func loggerHook() *lfshook.LfsHook {
	return lfshook.NewHook(
		lfshook.WriterMap{
			logrus.InfoLevel:  fileDivisionByTime("Info"),
			logrus.DebugLevel: fileDivisionByTime("Debug"),
			logrus.WarnLevel:  io.MultiWriter(fileDivisionByTime("Warn"), syncLog),
			logrus.PanicLevel: io.MultiWriter(fileDivisionByTime("Panic"), syncLog),
			logrus.FatalLevel: io.MultiWriter(fileDivisionByTime("Fatal"), syncLog),
		}, &logrus.JSONFormatter{
			TimestampFormat: "2006-01-02 15:04:05",
			PrettyPrint:     true,
		})
}

func fileDivisionByTime(level string) *rotatelogs.RotateLogs {
	division, err := rotatelogs.New(
		"./log/"+level+"/%Y-%m-%d.log",
		rotatelogs.WithMaxAge(time.Hour*24*7),
		rotatelogs.WithRotationTime(time.Hour*24),
	)
	if err != nil {
		logrus.WithError(err).WithField("stack", fmt.Sprintf("%+v", errors.WithStack(err))).Fatal()
	}
	return division
}
