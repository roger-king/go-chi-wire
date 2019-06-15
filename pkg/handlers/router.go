package handlers

import (
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/roger-king/go-chi-wire/pkg/handlers/healthcheck"
	"github.com/rs/cors"
)

// NewRouter - creates go-chi application router
func NewRouter() http.Handler {
	r := chi.NewRouter()

	debug, _ := strconv.ParseBool(os.Getenv("DEBUG"))

	c := cors.New(cors.Options{
		AllowedOrigins: []string{"*"},                                       // All origins
		AllowedMethods: []string{"POST", "GET", "PUT", "DELETE", "OPTIONS"}, // Allowing only get, just an example
		Debug:          debug,
	})

	r.Use(c.Handler)
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	// Set a timeout value on the request context (ctx), that will signal
	// through ctx.Done() that the request has timed out and further
	// processing should be stopped.
	r.Use(middleware.Timeout(60 * time.Second))

	r.Mount("/api/status", healthcheck.Router())

	return r
}
