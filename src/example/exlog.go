package main

import (
	"io"
	"log"
	"os"
)

var cLogger *log.Logger
var fLogger *log.Logger

// Log Example
func main() {
	//1. General Logging
	log.Println("General Logging")

	//2. Customer Logging
	cLogger = log.New(os.Stdout, "CUSTOMER : ", log.LstdFlags)

	//3. File Logging
	fLog, err := os.OpenFile("flog.txt", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		panic(err)
	}
	defer fLog.Close()

	fLogger = log.New(fLog, "FILE : ", log.Ldate|log.Ltime|log.Lshortfile)

	if true {
		//4. General -> File Logging
		gfLog, err := os.OpenFile("gflog.txt", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
		if err != nil {
			panic(err)
		}
		defer gfLog.Close()

		log.SetOutput(gfLog)
	} else {
		//5. General + File Logging
		fpLog, err := os.OpenFile("logfile.txt", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
		if err != nil {
			panic(err)
		}
		defer fpLog.Close()

		multiWriter := io.MultiWriter(fpLog, os.Stdout)
		log.SetOutput(multiWriter)
	}

	///////////////////////////////////////////

	cLogger.Print("Customer Strat")
	fLogger.Print("File Strat")
	log.Print("Start")

	cLogger.Println("Customer End")
	fLogger.Println("File End")
	log.Print("End")
}
