package logger

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/rs/zerolog"
)

var Log zerolog.Logger

func InitLogger(logFilePath string) {
	// Ensure the directory exists
	dir := filepath.Dir(logFilePath)
	fmt.Print(dir)
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		if err := os.MkdirAll(dir, 0755); err != nil {
			fmt.Println("Warning: Failed to create log directory, using default location")
			// Fallback to a safe location
			homeDir, _ := os.UserHomeDir()
			fmt.Print(homeDir)
			logFilePath = filepath.Join(homeDir, "logs", "app.txt")
			_ = os.MkdirAll(filepath.Dir(logFilePath), 0755)
		}
	}

	// Open or create the log file
	logFile, err := os.OpenFile(logFilePath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		panic(fmt.Sprintf("Failed to open log file: %v", err))
	}

	// Log to both file and stdout
	multi := zerolog.MultiLevelWriter(os.Stdout, logFile)
	Log = zerolog.New(multi).With().Timestamp().Logger()
}

func Info(tag string, msg string) {
	Log.Info().Msg(fmt.Sprintf("INFO: (%s) %s", tag, msg))
}

func Warning(msg string) {
	Log.Warn().Msg(msg)
}

func Error(tag string, msg string, err error) {
	Log.Error().Err(err).Msg(fmt.Sprintf("ERROR: (%s) %s", tag, msg))
}

func Critical(msg string, err error) {
	Log.Fatal().Err(err).Msg(msg)
}
