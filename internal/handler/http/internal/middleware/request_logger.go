package middleware

import (
	"net/http"

	"github.com/Pochirify/pochirify-backend/internal/handler/logger"
)

func NewRequestLogger(lf logger.Factory) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			queryInfo, err := getQueryInfo(r)
			if err != nil {
				lf(r.Context()).Error(err, "failed to get query name")
			}

			lf(r.Context()).WithValues(
				"method", r.Method,
				"host", r.URL.Hostname(),
				"path", r.URL.Path,
				"operation", queryInfo.getOperation(),
				"queryName", queryInfo.getOperationName(),
			).Info("http request")

			lrw := newLoggingResponseWriter(r, w)
			next.ServeHTTP(lrw, r)

			switch {
			case lrw.statusCode >= 500:
				lf(r.Context()).WithValues(
					"method", r.Method,
					"code", lrw.statusCode,
				).Error(nil, "internal error")

			case lrw.statusCode >= 400:
				lf(r.Context()).WithValues(
					"method", r.Method,
					"code", lrw.statusCode,
				).Info("client error")

			case lrw.statusCode == http.StatusOK:
				lf(r.Context()).WithValues(
					"method", r.Method,
					"code", lrw.statusCode,
					// "body", lrw.body.String(),
				).Info("finished")

			default:
				lf(r.Context()).WithValues(
					"method", r.Method,
					"code", lrw.statusCode,
				).Error(nil, "unexpected status code")
			}
		})
	}
}

type loggingResponseWriter struct {
	http.ResponseWriter
	statusCode int
	// body       *bytes.Buffer
}

func newLoggingResponseWriter(r *http.Request, w http.ResponseWriter) *loggingResponseWriter {
	return &loggingResponseWriter{
		ResponseWriter: w,
		statusCode:     http.StatusOK,
		// body:           bytes.NewBufferString(""),
	}
}

func (w *loggingResponseWriter) Write(b []byte) (int, error) {
	// w.body.Write(b)
	return w.ResponseWriter.Write(b)
}

func (w *loggingResponseWriter) WriteHeader(statusCode int) {
	w.statusCode = statusCode
	w.ResponseWriter.WriteHeader(statusCode)
}
