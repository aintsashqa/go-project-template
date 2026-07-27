package middleware

import (
	slogext "github.com/aintsashqa/go-project-template/pkg/slog-ext"
	"github.com/go-chi/chi/v5/middleware"
)

var (
	Recoverer Middleware = middleware.Recoverer
	RequestID Middleware = middleware.RequestID
)

func initExternal() {
	slogext.AddContext(middleware.RequestIDKey, "request_id")
}
