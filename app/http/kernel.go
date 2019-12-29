package http

import (
	"github.com/lanvard/foundation"
	"github.com/lanvard/foundation/http"
	foundationMiddleware "github.com/lanvard/foundation/http/middleware"
	"lanvard/app/http/middleware"
)

func NewKernel(app *foundation.Application) http.Kernel {
	return http.Kernel{
		App:        app,
		Middleware: pipes(),
	}
}

func pipes() []middleware.PipeInterface {
	return []middleware.PipeInterface{
		// todo remove or use ValidatePostSize
		foundationMiddleware.ValidatePostSize{},
	}
}
