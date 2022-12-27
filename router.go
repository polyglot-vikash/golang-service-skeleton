// router.go
package main

import "github.com/gorilla/mux"

// Set up router
func setupRouter() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/", indexHandler).Methods("GET")
	r.HandleFunc("/users", usersHandler).Methods("GET")
	r.HandleFunc("/users/{id}", userHandler).Methods("GET")
	return r
}
