package main

import (
	"deployee/logger"
	"deployee/webserver"
)

func main() {
	log := logger.Logging{}
	log.Info("Starting Deployee with the following configuration:\n")
	webserver := webserver.InitServerStruct(&log, "9999")

	// Will be the last line in the main function
	webserver.StartServer()
}
