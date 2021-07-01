package main

import (
	"os"

	log "github.com/sirupsen/logrus"
)

func init() {
	// Log as JSON instead of the default ASCII formatter.
	// default: log.SetFormatter(&log.TextFormatter{})
	log.SetFormatter(&log.JSONFormatter{})

	// Output to stdout instead of the default stderr
	// log.SetOutput(os.Stdout)
	// Can be any io.Writer
	f, err := os.OpenFile("demo.log", os.O_WRONLY|os.O_CREATE, 0755)
	if err != nil {
		log.Fatal("error: create log file")
	}
	log.SetOutput(f)

	// Only log the info severity or above.
	log.SetLevel(log.InfoLevel)
}

func main() {
	// A common pattern is to re-use fields between logging statements by re-using
	// the logrus.Entry returned from WithFields()
	contextLogger := log.WithFields(log.Fields{
		"common": "this is a common field",
		"other":  "I also should be logged always",
	})
	contextLogger.Info("I'll be logged with common and other field")
	contextLogger.Info("Me too")

	log.WithFields(log.Fields{
		"animal": "walrus",
		"size":   10,
	}).Info("A group of walrus emerges from the ocean")

	log.WithFields(log.Fields{
		"omg":    true,
		"number": 122,
	}).Warn("The group's number increased tremendously!")

	log.WithFields(log.Fields{
		"omg":    true,
		"number": 100,
	}).Fatal("The ice breaks!")
}
