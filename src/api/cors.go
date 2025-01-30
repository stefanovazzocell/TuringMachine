package api

import (
	"log/slog"
	"net/http"
)

const (
	CorsMethods = ", OPTIONS"
	CorsMaxAge  = "86400" // 24h
)

// Wrapper for requests to add cors handler
func (a *api) corsWrapper(method string, handler func(http.ResponseWriter, *http.Request)) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		slog.Debug("request",
			"remote", r.RemoteAddr,
			"uri", r.RequestURI,
			"method", r.Method)

		w.Header().Set("Access-Control-Allow-Origin", a.config.CorsOrigins)
		w.Header().Set("Access-Control-Request-Method", method+CorsMethods)
		w.Header().Set("Access-Control-Max-Age", CorsMaxAge)
		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusNoContent)
			return
		}

		handler(w, r)
	}
}
