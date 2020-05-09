package decorator

import (
	"github.com/lanvard/contract/inter"
	"lanvard/config"
)

type BootProviders struct{}

// Providers are located in config/providers/providers.go
func (r BootProviders) Bootstrap(app inter.App) inter.App {
	for _, bootstrapper := range config.Providers.BootProviders {
		app = bootstrapper.Boot(app)
	}

	return app
}