package restql

import (
	"fmt"
	"strings"

	"github.com/sirupsen/logrus"
	"gopkg.in/natefinch/lumberjack.v2"
)

var Log *logrus.Logger

func init() {
	// Initialize the logger
	Log = logrus.New()

	// Set up log rotation
	Log.SetOutput(&lumberjack.Logger{
		Filename:   "restql.log", // Log file name
		MaxSize:    10,           // Maximum size in megabytes before rotating
		MaxBackups: 3,            // Maximum number of backup files
		MaxAge:     28,           // Maximum number of days to retain old log files
		Compress:   true,         // Compress rotated log files
	})

	// Optionally set log formatter
	Log.SetFormatter(&logrus.TextFormatter{
		FullTimestamp: true,
	})

	// Set log level if needed
	Log.SetLevel(logrus.InfoLevel)
}

func LogError(args ...interface{}) {
	var stringArgs []string
	for _, arg := range args {
		stringArgs = append(stringArgs, fmt.Sprint(arg))
	}
	message := strings.Join(stringArgs, " | ")
	Log.Error("RESTQL | " + message)
}
func LogInfo(args ...interface{}) {
	var stringArgs []string
	for _, arg := range args {
		stringArgs = append(stringArgs, fmt.Sprint(arg))
	}
	message := strings.Join(stringArgs, " | ")
	Log.Info("RESTQL | " + message)
}
