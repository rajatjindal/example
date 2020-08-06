package api

import (
	"net/http"

	"github.com/gorilla/mux"
)

//Server is api server
type Server struct {
	Router *mux.Router
}

//New returns new server
func New() (*Server, error) {
	router := mux.NewRouter().StrictSlash(true)
	server := &Server{
		Router: router,
	}

	server.addRoutes()
	return server, nil
}

//order matters
func (s *Server) addRoutes() {
	s.Router.Methods(http.MethodOptions).HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})

	//endpoints
	s.Router.Methods(http.MethodGet).Path("/api/endpoints").HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("hello from /endpoints"))
	})

	//add middleware after all endpoints are added
	s.Router.Use(CorsHandler())
}
