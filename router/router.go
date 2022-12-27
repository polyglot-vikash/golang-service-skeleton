// router.go
package router

import (
	"net/http"
	"sample-app/handler"

	"github.com/gorilla/mux"
)

// Set up router
func SetupRouter() *mux.Router {
	r := mux.NewRouter()

	r.HandleFunc("/ping", handler.PingHandler).Methods("GET")
	r.HandleFunc("/users", handler.UsersHandler).Methods("GET")
	r.HandleFunc("/users/{id}", handler.UserHandler).Methods("GET")

	http.Handle("/", r)
	http.ListenAndServe(":8080", r)
	return r
}
