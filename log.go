package main

import (
	"fmt"
	"log"
	"os"
)

// Alog is ...
func Alog(str string) {

	//set logfile Stdout
	logFile, logErr := os.OpenFile(*logFileName, os.O_CREATE|os.O_RDWR|os.O_APPEND, 0666)
	if logErr != nil {
		fmt.Println("Fail to find", *logFile, "cServer start Failed")
		os.Exit(1)
	}
	log.SetOutput(logFile)
	log.SetFlags(log.Ldate | log.Ltime)

	//write log
	log.Printf(":%v : %v \n", "dat", str)
}
