package startup

import (
	"log"
	"os"
	"time"
)

var (
	InfoLog  *log.Logger
	ErrorLog *log.Logger
)

type Log struct{}

func (l *Log) InitialLog() {
	// Prepare log file paths with current date prefix
	currentDate := time.Now().Format("2006-01-02-")
	logPath := "./" + os.Getenv("LOGGING_DIR") + "/"

	// Define info and error log files
	infoLogFile := logPath + currentDate + os.Getenv("INFO_LOG_PATH")
	errorLogFile := logPath + currentDate + os.Getenv("ERROR_LOG_PATH")

	// Create loggers
	InfoLog = createLogger(infoLogFile, "INFO: ")
	ErrorLog = createLogger(errorLogFile, "ERROR: ")
}

func createLogger(filePath, prefix string) *log.Logger {
	logFile, err := os.OpenFile(filePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatalf("Failed to open log file %s: %v", filePath, err)
	}

	return log.New(logFile, prefix, log.Ldate|log.Ltime|log.Lshortfile)
}
