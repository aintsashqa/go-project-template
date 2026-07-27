package httpapi

import (
	"net/http"

	"github.com/aintsashqa/go-project-template/internal/delivery/http-api/handler"
	"github.com/aintsashqa/go-project-template/internal/delivery/http-api/middleware"
)

type Registrar interface {
	Register(*http.ServeMux)
}

func NewHandler(registrars ...Registrar) http.Handler {
	mux := http.NewServeMux()

	mux.HandleFunc("/health", handler.Health)

	for _, registrar := range registrars {
		registrar.Register(mux)
	}

	return middleware.Pipe(mux,
		middleware.Logger,
		middleware.RequestID,
		middleware.Recoverer)
}
