package main

import (
	"log"
	"os"
)

func main() {
	logfile, _ := os.Create("./22-log-to-writer/log.txt")
	defer logfile.Close()

	logger := log.New(logfile, "example ", log.LstdFlags|log.Lshortfile)

	logger.Println("This is a regular message")
	logger.Fatalln("This is a fatal error")
	
}