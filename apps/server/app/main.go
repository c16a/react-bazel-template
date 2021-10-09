package main

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {
	r := mux.NewRouter()

	srv := &http.Server{
		Handler: r,
		Addr:    ":8080",
	}

	log.Fatalln(srv.ListenAndServe())
}
