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
	"github.com/cmd/internal/forms"
	"github.com/pkg/errors"
)

func PageLogin(w http.ResponseWriter, r *http.Request) {

	if r.URL.Path != "/login/view" {
		http.Error(w, "Invalid URL Address", http.StatusRequestURITooLong)
	}

	file, err := os.Open("./internal/templates/auth/login.html")
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

		datastorage := entities.DataUser{
			Password: r.FormValue("password"),
			UserName: r.FormValue("username"),
			Email:    r.FormValue("address"),
		}

		if forms.IsEmail(datastorage.Email) && forms.IsPassword(datastorage.Password) &&
			forms.IsUsername(datastorage.UserName) {

			if usecase.WebsiteAccess(&datastorage) {
				http.Redirect(w, r, "/welcome/view", http.StatusFound)
			} else {
				http.Error(w, "Login details are incorrect", http.StatusUnauthorized)
			}

		} else {
			fmt.Println("Invalid input data! ")
		}

	default:
		http.Redirect(w, r, "/page/error", http.StatusNotFound)
	}

}

func PageRegistration(w http.ResponseWriter, r *http.Request) {

	if r.URL.Path != "/registration/view" {
		http.Error(w, "Invalid URL Address", http.StatusRequestURITooLong)
	}

	fl, err := os.Open("./internal/templates/auth/signup.html")
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

		datastorage := entities.DataUser{
			Password: r.FormValue("password"),
			UserName: r.FormValue("username"),
			Email:    r.FormValue("address"),
		}

		if forms.IsEmail(datastorage.Email) && forms.IsPassword(datastorage.Password) &&
			forms.IsUsername(datastorage.UserName) {

			if usecase.ExistsUser(&datastorage) {
				fmt.Printf("User with email exists already")
			} else {
				storage.InsertDB(&datastorage)
			}
		} else {
			http.Error(w, "Login details are incorrect", http.StatusUnauthorized)
		}
	default:
		http.Redirect(w, r, "/page/error", http.StatusNotFound)
	}
}

func PageResetPassword(w http.ResponseWriter, r *http.Request) {

	if r.URL.Path != "/recovery/password" {
		http.Error(w, "Invalid URL Address", http.StatusRequestURITooLong)
	}

	fl, err := os.Open("./internal/templates/recovery.html")
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

		datastorage := entities.DataUser{
			Password: r.FormValue("password"), // get new password from website
			Email:    r.FormValue("addrress"),
		}

		if forms.IsEmail(datastorage.Email) && forms.IsPassword(datastorage.Password) {

			if usecase.ExistsUser(&datastorage) {
				usecase.ChangePassword(&datastorage)
			} else {
				fmt.Println("User doens't exist")
			}
		} else {
			http.Error(w, "Login details are incorrect", http.StatusUnauthorized)
		}

	default:
		http.Redirect(w, r, "/page/error", http.StatusNotFound)
	}
}

func PageMain(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/welcome/view" {
		http.Error(w, "Invalid URL Address", http.StatusRequestURITooLong)
	}

	fl, err := os.Open("./internal/templates/home.html")
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
