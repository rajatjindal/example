package api

import "net/http"

type corsMiddleware struct {
	handler http.Handler
}

func (m *corsMiddleware) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("access-control-allow-origin", "*")
	w.Header().Set("access-control-allow-methods", "OPTIONS, GET, POST, PATCH, PUT, HEAD")
	w.Header().Set("access-control-allow-headers", "authorization, content-type")
	m.handler.ServeHTTP(w, r)
}

//CorsHandler is for gorrilla middleware
func CorsHandler() func(h http.Handler) http.Handler {
	return func(h http.Handler) http.Handler {
		return &corsMiddleware{
			handler: h,
		}
	}
}
