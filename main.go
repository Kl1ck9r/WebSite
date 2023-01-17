package main

import (
	"log"
	"net/http"
	"time"
	
	"github.com/cmd/internal/middleware"
	"github.com/cmd/internal/repository"
)

func main() {


	mux := http.NewServeMux()

	mux.HandleFunc("/page/registration", server.PageRegistration)

	mux.HandleFunc("/page/login", server.PageLogin)

	mux.HandleFunc("/page/reset/password", server.PageResetPassword)

	mux.HandleFunc("/page/main", server.PageMain)

	mux.HandleFunc("/page/show/notes", server.ShowNotesHander)

	mux.HandleFunc("/page/error", server.ErrorHandler)

	handler := middleware.Logging(mux)
    handler = middleware.PanicRecovery(handler)

	srv := &http.Server{
		Handler:        mux,
		Addr:         "127.0.0.1:8080",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Fatal(srv.ListenAndServe())
}
