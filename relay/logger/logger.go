package logger

import (
	"log"
	"os"
)

var (
	infoLogger  *log.Logger
	errorLogger *log.Logger
)

func Init(logFilePath string) {
	file, err := os.OpenFile(logFilePath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatal("Failed to open log file:", err)
	}

	infoLogger = log.New(file, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)
	errorLogger = log.New(file, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile)
}

func Info(message string, packageName string, actionName string) {
	infoLogger.Printf("PACKAGE NAME : %s  -  ACTION : %s - MESSAGE : %s\n", packageName, actionName, message)
}

func Error(message string, packageName string, actionName string) {
	errorLogger.Printf("PACKAGE NAME : %s  -  ACTION : %s - MESSAGE : %s\n", packageName, actionName, message)
}
