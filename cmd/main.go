package main

import (
	root "github.com/roger-king/go-chi-wire/pkg"
	log "github.com/sirupsen/logrus"
)

func main() {
	server, _ := root.InitializeGoChiWire()

	log.Info("Application is now running")
	log.Fatal(server.ListenAndServe())
}
