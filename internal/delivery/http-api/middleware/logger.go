package middleware

import (
	"log/slog"
	"net/http"
	"time"
)

type loggerResponseWriter struct {
	http.ResponseWriter
	statusCode    int
	contentLength int
}

func (lrw *loggerResponseWriter) Write(data []byte) (int, error) {
	n, err := lrw.ResponseWriter.Write(data)
	lrw.contentLength = n

	return n, err
}

func (lrw *loggerResponseWriter) WriteHeader(statusCode int) {
	lrw.ResponseWriter.WriteHeader(statusCode)
	lrw.statusCode = statusCode
}

func Logger(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var (
			now = time.Now()
			lrw = loggerResponseWriter{
				ResponseWriter: w,
			}
		)

		defer func(start time.Time) {
			slog.InfoContext(r.Context(), "http request",
				"method", r.Method,
				"path", r.URL.Path,
				"status_code", lrw.statusCode,
				"content_length", lrw.contentLength,
				"duration", time.Since(start).String(),
			)
		}(now)

		next.ServeHTTP(&lrw, r)
	})
}
