package middleware

import "net/http"

type Middleware func(http.Handler) http.Handler

func Pipe(handler http.Handler, middlewares ...Middleware) http.Handler {
	for _, middleware := range middlewares {
		handler = middleware(handler)
	}

	return handler
}

func PipeFunc(handlerFunc http.HandlerFunc, middlewares ...Middleware) http.Handler {
	return Pipe(http.Handler(handlerFunc), middlewares...)
}
