package server

import (
	"fmt"
	"io/ioutil"
	"log"

	"net/http"
	"os"

	"github.com/cmd/internal/database"
	"github.com/cmd/internal/database/datanotes"
	"github.com/cmd/internal/database/storage"
	"github.com/cmd/internal/entities"
	"github.com/cmd/internal/forms"
	"github.com/cmd/internal/repository/parser"
	"github.com/cmd/internal/utils"
	"github.com/pkg/errors"
)

func PageLogin(w http.ResponseWriter, r *http.Request) {

	file, err := os.Open("./templates/auth/login.html")
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

		if forms.IsEmail(datastorage.Email) && forms.IsPassword(datastorage.Password) {

			if usecase.WebsiteAccess(&datastorage) {
				http.Redirect(w, r, "/page/main", http.StatusSeeOther)
			} else {
				http.Error(w, "Login details are incorrect", http.StatusUnauthorized)
			}

		} else {
			fmt.Fprintf(w, "Invalid input data!")
		}

	default:
		http.Redirect(w, r, "/page/error", http.StatusNotFound)
	}
}

func PageRegistration(w http.ResponseWriter, r *http.Request) {

	fl, err := os.Open("./templates/auth/signup.html")
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
				fmt.Fprintf(w, "User with so email exists already")
			} else {
				db, err := utils.ConnectDB()
				if err != nil {
					log.Fatal("Failed to connect db ")
				}

				if err = storage.InsertDB(db, &datastorage); err != nil {
					log.Fatal("Failed to insert data user to database storage")
				}

				http.Redirect(w, r, "/page/login", http.StatusFound)
			}
		} else {
			http.Error(w, "Login details are incorrect", http.StatusUnauthorized)
		}
	default:
		http.Redirect(w, r, "/page/error", http.StatusNotFound)
	}
}

func PageResetPassword(w http.ResponseWriter, r *http.Request) {

	fl, err := os.Open("./templates/recovery.html")
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
				fmt.Fprintf(w, "User with so email doesn't exist")
			}
		} else {
			http.Error(w, "Login details are incorrect", http.StatusUnauthorized)
		}

	default:
		http.Redirect(w, r, "/page/error", http.StatusNotFound)
	}
}

func PageMain(w http.ResponseWriter, r *http.Request) {

	fl, err := os.Open("./templates/home.html")
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

		db, err := utils.ConnectDB()
		CheckError(err, "Failed to connect database")

		if err := notesdb.InsertNoteDB(db, &note); err != nil {
			log.Print(err)
		}

	default:
		http.Redirect(w, r, "/page/error", http.StatusNotFound)
	}
}

func ShowNotesHandler(w http.ResponseWriter, r *http.Request) {
	notes, err := notesdb.GetNotes()
	CheckError(err, "Failed to get notes from database notesdb")

	parser.RenderTemplate(w, "./templates/showNotes", notes)
}

func DeleteNotesHandler(w http.ResponseWriter, r *http.Request) {
	file, err := os.Open("./templates/deleteNote.html")
	CheckError(err, "Failed to open file")

	defer file.Close()

	read, err := ioutil.ReadAll(file)
	CheckError(err, "Failed to read of the html file")

	switch r.Method {
	case "GET":
		w.Write(read)

	case "POST":
		db, err := utils.ConnectDB()
		CheckError(err, "Failed to connect database")

	
		noteID := entities.Notes{
			ID: r.FormValue("deleteNote"),
		}
		notesdb.DeleteNoteDB(db, &noteID)

	default:
		http.Redirect(w, r, "/page/error", http.StatusNotFound)
	}
}

func ErrorHandler(w http.ResponseWriter, r *http.Request) {

	title := r.URL.Path[len("page/error"):]
	p, err := parser.LoadPage(title)

	if err != nil {
		p = &entities.Page{Title: title}
	}

	parser.RenderTemplate(w, "./templates/errorpage", p)
}

func CheckError(err error, msg string) {
	if err != nil {
		errors.Wrap(err, msg)
	}
}
