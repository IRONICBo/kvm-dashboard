package utils

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"

	"github.com/sirupsen/logrus"
)

var log *logrus.Logger

func init() {
	log = logrus.New()
	log.SetFormatter(&logrus.TextFormatter{})
	log.SetReportCaller(true)

	log.SetFormatter(&logrus.TextFormatter{
		TimestampFormat: "2006-01-02 15:04:05.000",
		FullTimestamp:   true,
		CallerPrettyfier: func(f *runtime.Frame) (string, string) {
			return "", fmt.Sprintf(" %s:%d", filepath.Base(f.File), f.Line)
		},
	})

	log.SetOutput(os.Stdout)
	log.SetLevel(logrus.DebugLevel) // Default log level
}

func SetLogLevel(level logrus.Level) {
	log.SetLevel(level)
}
