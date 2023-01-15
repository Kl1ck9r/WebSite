package main

import (
	"github.com/cmd/internal/repository"
	"log"
	"net/http"
	"time"
)

func main() {

	mux := http.NewServeMux()

	mux.HandleFunc("page/registration", server.PageRegistration)

	mux.HandleFunc("page/login", server.PageLogin)

	mux.HandleFunc("page/reset/password", server.PageResetPassword)

	mux.HandleFunc("page/main", server.PageMain)

	mux.HandleFunc("/page/error", server.ErrorHandler)

	srv := &http.Server{
		Handler:      mux,
		Addr:         "127.0.0.1:8010",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Fatal(srv.ListenAndServe())
}
