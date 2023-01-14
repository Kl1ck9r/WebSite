package main

import (
	"github.com/cmd/internal/repository"
	//"github.com/cmd/internal/database/storage"
	//"github.com/cmd/internal/entities"
	"log"
	"net/http"
	"time"
)

func main() {

	//p:=entities.DataUser{
	//	ID : "19",
	//}
	
	//storage.DeleteDB(&p)

	mux := http.NewServeMux()

	mux.HandleFunc("/registration/view", server.PageRegistration)

	mux.HandleFunc("/login/view", server.PageLogin)

	mux.HandleFunc("/recovery/password", server.PageResetPassword)

	mux.HandleFunc("/welcome/view", server.PageMain)

	mux.HandleFunc("/page/error", server.ErrorHandler)

	srv := &http.Server{
		Handler:      mux,
		Addr:         "127.0.0.1:8010",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Fatal(srv.ListenAndServe())
}
