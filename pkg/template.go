// +build wireinject
// The build tag makes sure the stub is not built in the final build.

package root

import (
	"net/http"

	"github.com/roger-king/go-chi-wire/pkg/server"

	"github.com/google/wire"
)

// GoChiWireSet - creates a wire set for the whole application.
var GoChiWireSet = wire.NewSet(server.ServerSet)

// InitializeGoChiWire - instructs wire to build dependency injection.
func InitializeGoChiWire() (*http.Server, error) {
	wire.Build(GoChiWireSet)
	return nil, nil
}
