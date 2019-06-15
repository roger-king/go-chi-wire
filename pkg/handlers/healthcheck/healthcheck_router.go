package healthcheck

import (
	"net/http"

	"github.com/go-chi/chi"
)

// HeathCheckHandler -
type HeathCheckHandler http.Handler

// Router - Healthcheck Router
func Router() HeathCheckHandler {
	r := chi.NewRouter()

	r.Get("/", Get)

	return r
}
