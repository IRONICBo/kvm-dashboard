package utils

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"

	"github.com/sirupsen/logrus"
)

var Log *logrus.Logger

func init() {
	Log = logrus.New()
	Log.SetFormatter(&logrus.TextFormatter{})
	Log.SetReportCaller(true)

	Log.SetFormatter(&logrus.TextFormatter{
		TimestampFormat: "2006-01-02 15:04:05.000",
		FullTimestamp:   true,
		CallerPrettyfier: func(f *runtime.Frame) (string, string) {
			return "", fmt.Sprintf(" %s:%d", filepath.Base(f.File), f.Line)
		},
	})

	Log.SetOutput(os.Stdout)
	Log.SetLevel(logrus.DebugLevel) // Default log level
}

func SetLogLevel(level logrus.Level) {
	Log.SetLevel(level)
}
