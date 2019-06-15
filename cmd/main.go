package main

import (
	root "github.com/roger-king/go-chi-wire/pkg"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"

	"os"
)

func init() {
	// Log as JSON instead of the default ASCII formatter.
	log.SetFormatter(&log.JSONFormatter{})

	// Output to stdout instead of the default stderr
	// Can be any io.Writer, see below for File example
	log.SetOutput(os.Stdout)

	// Only log the warning severity or above.
	log.SetLevel(log.InfoLevel)
}

func main() {
	server, _ := root.InitializeGoChiWire()

	log.Info("Application is starting on http://localhost:", viper.GetString("server.port"))
	log.Fatal(server.ListenAndServe())
}
