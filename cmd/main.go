package main

import (
	root "github.com/roger-king/go-chi-wire/pkg"
	log "github.com/sirupsen/logrus"
)

func main() {
	server, _ := root.InitializeApp()
	log.Fatal(server.ListenAndServe())
}
