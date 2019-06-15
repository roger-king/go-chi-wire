package server

import (
	"net/http"
	"os"

	"github.com/google/wire"
	_ "github.com/joho/godotenv/autoload"
	"github.com/roger-king/go-chi-wire/pkg/config"
	"github.com/roger-king/go-chi-wire/pkg/handlers"
	log "github.com/sirupsen/logrus"
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

// ServerSet -
var ServerSet = wire.NewSet(config.NewConfig, handlers.NewRouter, CreateServer)

// CreateServer - creates the http server
func CreateServer(h http.Handler, c *config.Config) *http.Server {
	port := c.Server.Port
	server := &http.Server{
		Addr:    ":" + port,
		Handler: h,
	}

	return server
}
