package logging

import (
	"log"
	"os"
)

// TODO Finish logging with output in console and file
var (
	InfoLogger    *log.Logger
	ErrorLogger   *log.Logger
	WarningLogger *log.Logger
)

func Init() {
	err := os.MkdirAll("logs", 0644)
	if err != nil {
		panic(err)
	}

	file, err := os.OpenFile("logs.txt", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatal(err)
	}

	//	defer file.Close()

	InfoLogger = log.New(file, "INFO:", log.LstdFlags|log.Lshortfile)
	ErrorLogger = log.New(file, "ERROR:", log.LstdFlags|log.Lshortfile)
	WarningLogger = log.New(file, "WARNING:", log.LstdFlags|log.Lshortfile)
}
