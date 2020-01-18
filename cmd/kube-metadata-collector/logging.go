package main

import (
	"os"

	log "github.com/sirupsen/logrus"
)

func init() {
	log.SetFormatter(&log.JSONFormatter{})
	log.SetReportCaller(true)
	log.SetOutput(os.Stdout)
	log.SetLevel(log.InfoLevel)
}
