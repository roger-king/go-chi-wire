// +build wireinject
// The build tag makes sure the stub is not built in the final build.

package root

import (
	"net/http"

	"github.com/roger-king/go-chi-wire/pkg/server"

	"github.com/google/wire"
)

// GoChiWireSet - creates a wire set for the whole application.
var AppSet = wire.NewSet(server.ServerSet)

// InitializeGoChiWire - instructs wire to build dependency injection.
func InitializeApp() (*http.Server, error) {
	wire.Build(AppSet)
	return nil, nil
}
