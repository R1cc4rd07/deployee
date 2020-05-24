package webserver

import (
	"deployee/logger"
	"encoding/json"
	"net/http"
	"strings"
	"time"

	"github.com/gorilla/mux"
)

// Server struct holds the logger and Router
type Server struct {
	Log    *logger.Logging
	Router *mux.Router
	Srv    *http.Server
}

// InitServerStruct initalize the Server struct and the mux Router. Will create all existing routs. Returns the Server struct as a pointer
func InitServerStruct(l *logger.Logging, port string) *Server {
	r := mux.NewRouter()

	srv := &http.Server{
		Handler:      r,
		Addr:         "0.0.0.0:" + port,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	router := &Server{Log: l, Router: r, Srv: srv}

	// Create new subroute
	sv1 := r.PathPrefix("/rest/v1").Subrouter()
	// Add new route to endpoint
	// i.e: sv1.HandleFunc("/health", health).Methods(http.MethodGet)
	sv1.HandleFunc("/health", health).Methods(http.MethodGet)

	return router
}

// StartServer starts the HTTP server
func (server *Server) StartServer() {
	server.Log.Infof("Starting the webserver on port: %s", strings.TrimLeft(server.Srv.Addr, ":"))

	server.Log.Fatal(server.Srv.ListenAndServe())
}

// TODO --> REMOVE after creating real endpoints
func health(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(map[string]bool{"ok": true})
}
