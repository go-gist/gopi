package rest

import (
	"fmt"
	"strings"

	"github.com/sirupsen/logrus"
	"gopkg.in/natefinch/lumberjack.v2"
)

var log *logrus.Logger

func init() {
	// Initialize the logger
	log = logrus.New()

	// Set up log rotation
	log.SetOutput(&lumberjack.Logger{
		Filename:   "rest.log", // log file name
		MaxSize:    10,         // Maximum size in megabytes before rotating
		MaxBackups: 3,          // Maximum number of backup files
		MaxAge:     28,         // Maximum number of days to retain old log files
		Compress:   true,       // Compress rotated log files
	})

	// Optionally set log formatter
	log.SetFormatter(&logrus.TextFormatter{
		FullTimestamp: true,
	})

	// Set log level if needed
	log.SetLevel(logrus.InfoLevel)
}

func logError(args ...interface{}) {
	var stringArgs []string
	for _, arg := range args {
		stringArgs = append(stringArgs, fmt.Sprint(arg))
	}
	message := strings.Join(stringArgs, " | ")
	log.Error("RESTQL | " + message)
}
func logInfo(args ...interface{}) {
	var stringArgs []string
	for _, arg := range args {
		stringArgs = append(stringArgs, fmt.Sprint(arg))
	}
	message := strings.Join(stringArgs, " | ")
	log.Info("RESTQL | " + message)
}
