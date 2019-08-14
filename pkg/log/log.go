package log

import (
	"flag"
	"log"
	"os"
	"path/filepath"

	"github.com/sirupsen/logrus"
)

var logPath string

func init() {

	flag.StringVar(&logPath, "logs", "./logs", "default config path")

	// Log as JSON instead of the default ASCII formatter.
	logrus.SetFormatter(&logrus.JSONFormatter{})

	if err := os.MkdirAll(logPath, 0777); err != nil {
		log.Fatalf("conf path make failed.err：%+v", err)
	}

	logFilePath := filepath.Join(logPath, "app.log")
	logFile, err := os.Create(logFilePath)

	if err != nil {
		log.Fatalf("log file create failed.err:%+v", err)
	}

	// Output to stdout instead of the default stderr
	// Can be any io.Writer, see below for File example
	// logrus.SetOutput(os.Stdout)
	logrus.SetOutput(logFile)

	// Only log the warning severity or above.
	logrus.SetLevel(logrus.TraceLevel)
}

// Dir log dir
func Dir() string {
	return logPath
}

// Infof log infof
func Infof(format string, args ...interface{}) {
	logrus.Infof(format, args...)
}
