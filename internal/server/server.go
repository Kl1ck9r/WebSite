package server

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/cmd/internal/database"
	"github.com/cmd/internal/database/datanotes"
	"github.com/cmd/internal/database/storage"
	"github.com/cmd/internal/entities"
	"github.com/pkg/errors"
)

func PageLogin(w http.ResponseWriter, r *http.Request) {

	if r.URL.Path != "/login/view" {
		http.Error(w, "Invalid URL Address", http.StatusRequestURITooLong)
	}

	file, err := os.Open("./internal/repository/auth/login.html")
	CheckError(err, "Failed to open file")

	read, err := ioutil.ReadAll(file)
	CheckError(err, "Failed to read file")

	defer file.Close()

	switch r.Method {
	case "GET":
		w.Write(read)

	case "POST":
		if err := r.ParseForm(); err != nil {
			fmt.Fprintf(w, "ParseForm()err:%v", err)
			return
		}

		t_storage := entities.DataUser{
			Password: r.FormValue("password"),
			UserName: r.FormValue("username"),
			Email:    r.FormValue("address"),
		}

		if usecase.WebsiteAccess(&t_storage) {
			http.Redirect(w, r, "/welcome/view", http.StatusFound)
		} else {
			http.Error(w, "Error to login", http.StatusUnauthorized)
		}

	default:
		http.Redirect(w, r, "/page/error", http.StatusNotFound)
	}

}

func PageRegistration(w http.ResponseWriter, r *http.Request) {

	if r.URL.Path != "/registration/view" {
		http.Error(w, "Invalid URL Address", http.StatusRequestURITooLong)
	}

	fl, err := os.Open("./internal/repository/auth/signup.html")
	CheckError(err, "Failed to open file")

	read, err := ioutil.ReadAll(fl)
	CheckError(err, "Failed to read file")

	defer fl.Close()

	switch r.Method {
	case "GET":
		w.Write(read)

	case "POST":
		if err := r.ParseForm(); err != nil {
			fmt.Fprintf(w, "ParseForm()err:%v", err)
			return
		}

		t_storage := entities.DataUser{
			Password: r.FormValue("password"),
			UserName: r.FormValue("username"),
			Email:    r.FormValue("address"),
		}

		if usecase.ExistsUser(&t_storage) {
			fmt.Printf("User with email exists already")
		} else {
			storage.InsertDB(&t_storage)
		}

	default:
		http.Redirect(w, r, "/page/error", http.StatusNotFound)
	}
}

func PageResetPassword(w http.ResponseWriter, r *http.Request) {

	if r.URL.Path != "/recovery/password" {
		http.Error(w, "Invalid URL Address", http.StatusRequestURITooLong)
	}

	fl, err := os.Open("./internal/repository/recovery.html")
	CheckError(err, "Failed to open file")

	defer fl.Close()

	read, err := ioutil.ReadAll(fl)
	CheckError(err, "Failed to read file")

	switch r.Method {
	case "GET":
		w.Write(read)

	case "POST":
		if err := r.ParseForm(); err != nil {
			fmt.Fprintf(w, "ParseForm()err:%v", err)
			return
		}

		storage := entities.DataUser{
			Password: r.FormValue("password"), // get new password from website
			Email:    r.FormValue("addrress"),
		}

		if usecase.ExistsUser(&storage) {
			usecase.ChangePassword(&storage)
		} else {
			fmt.Println("User doens't exist")
		}

	default:
		http.Redirect(w, r, "/page/error", http.StatusNotFound)
	}
}

func PageMain(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/welcome/view" {
		http.Error(w, "Invalid URL Address", http.StatusRequestURITooLong)
	}

	fl, err := os.Open("./internal/repository/home.html")
	CheckError(err, "Failed to open db")

	defer fl.Close()

	read, err := ioutil.ReadAll(fl)
	CheckError(err, "Failed to read db")

	switch r.Method {

	case "GET":
		w.Write(read)

	case "POST":
		if err := r.ParseForm(); err != nil {
			fmt.Fprintf(w, "ParseForm()err:%v", err)
			return
		}

		note := entities.Notes{
			Note: r.FormValue("note"),
			ID:   r.FormValue("id"),
		}

		notesdb.InsertNoteDB(&note)

	default:
		http.Redirect(w, r, "/page/error", http.StatusNotFound)
	}
}

func ErrorHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/page/error" {
		http.Error(w, "Invalid URL Address", http.StatusRequestURITooLong)
	}

	w.Write([]byte("Error"))
}

func CheckError(err error, msg string) {
	if err != nil {
		errors.Wrap(err, msg)
	}
}
