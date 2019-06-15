package server

import (
	"net/http"

	"github.com/google/wire"
	"github.com/roger-king/go-chi-wire/pkg/config"
	"github.com/roger-king/go-chi-wire/pkg/handlers"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

// ServerSet -
var ServerSet = wire.NewSet(handlers.NewRouter, CreateServer)

func init() {
	config.Read()
}

// CreateServer - creates the http server
func CreateServer(h http.Handler) *http.Server {
	log.Info("Server")
	port := viper.GetString("PORT")
	server := &http.Server{
		Addr:    ":" + port,
		Handler: h,
	}

	return server
}
