package flint

import (
	"fmt"
	"os"
	"time"
)

var logFile *os.File

func CreateLogFile(fileName string) error {
	var err error
	logFile, err = os.OpenFile(fileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return fmt.Errorf("failed to create or open log file: %w", err)
	}

	fmt.Printf("Log file '%s' created successfully.\n", fileName)
	return nil
}

func LogError(function string, message string) {

	timestamp := time.Now().Format(time.RFC3339)

	fmt.Printf("\033[31m[Flint ERROR]\033[0m [%s] [%s] %s\n", timestamp, function, message)

	if logFile != nil {
		_, err := fmt.Fprintf(logFile, "[%s] [ERROR] [%s] %s\n", timestamp, function, message)
		if err != nil {
			fmt.Printf("Failed to write to log file: %s\n", err.Error())
		}
	}
}
