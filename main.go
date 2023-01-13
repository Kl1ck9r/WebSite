package main

import (
	"github.com/cmd/internal/server"
	"net/http"
	"github.com/pkg/errors"
)

const(
	//tslCertFile = "./mirgations/ssl/"
	//tslKeyFile = "./migrations/ssl/"
)
// err:=http.ListenAndServeTLS()

func main() {

	mux := http.NewServeMux()
	
	mux.HandleFunc("/registration/view", server.PageRegistration)

	mux.HandleFunc("/login/view", server.PageLogin)

	mux.HandleFunc("/recovery/password", server.PageResetPassword)

	mux.HandleFunc("/welcome/view", server.PageMain)

	mux.HandleFunc("/page/error", server.ErrorHandler)

	fs := http.FileServer(http.Dir("./internal/repository/"))
	mux.Handle("/", fs)

	err := http.ListenAndServe(":8010", mux)

	if err != nil {
		errors.Wrap(err, "Failed to start server")
	}
}
