// handler.go
package handler

import (
	"fmt"
	"net/http"
)

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	// Handle index route
}

func UsersHandler(w http.ResponseWriter, r *http.Request) {
	// Handle users route
}

func UserHandler(w http.ResponseWriter, r *http.Request) {
	// Handle user route
}

func PingHandler(w http.ResponseWriter, r *http.Request) {
	//specify status code
	w.WriteHeader(http.StatusOK)

	//update response writer
	fmt.Fprintf(w, "API is up and running")
}
