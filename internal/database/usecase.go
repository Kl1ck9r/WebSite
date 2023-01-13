package usecase

import (
	"database/sql"
	"fmt"
	"log"

	dbstorage "github.com/cmd/internal/database/storage"
	"github.com/cmd/internal/entities"
	_ "github.com/lib/pq"
	"github.com/pkg/errors"
)

var Db *sql.DB

const (
	Psqlconnect = "user=postgres dbname=postgres password=Ruslan5655 host=localhost sslmode=disable"
)

func WebsiteAccess(ent *entities.DataUser) bool {
	Db, err := sql.Open("postgres", Psqlconnect)
	CheckError(err, "Failed to open  db")

	err = Db.Ping()
	CheckError(err, "Failed to connect db")

	var IsTrue bool

	storage := GetDataDB()

	for _, rng := range storage {
		if rng.Email == ent.Email && rng.Password == ent.Password && rng.UserName == ent.UserName {
			IsTrue = true
		}
	}

	return IsTrue
}

func ExistsUser(ent *entities.DataUser) bool {
	Db, err := sql.Open("postgres", Psqlconnect)
	CheckError(err, "Failed to open db")

	err = Db.Ping()
	CheckError(err, "Failed to connect db")

	rows, err := Db.Query("SELECT email FROM storage")

	if err != nil {
		errors.Wrap(err, "Failed to handle request db")
	}

	defer rows.Close()

	storage := []entities.DataUser{}

	for rows.Next() {
		p := entities.DataUser{}
		err := rows.Scan(&p.Email)
		CheckError(err, "Failed to copy of the db")

		storage = append(storage, p)
	}

	var IsTrue bool

	for _, rng := range storage {
		if rng.Email == ent.Email {
			//ChangePassword(ent)
			IsTrue = true
		} else {
			//log.Fatal("The email address not found database " +
			//"for password recovery an email address for db must be exists")
			//	IsTrue = false
		}
	}

	fmt.Println(ent.Email)

	return IsTrue

}

func ChangePassword(ent *entities.DataUser) {
	Db, err := sql.Open("postgres", Psqlconnect)
	CheckError(err, "Failed to open  db")

	err = Db.Ping()
	CheckError(err, "Failed to connect db")

	rows, err := Db.Query("SELECT password,email FROM notesdb")
	CheckError(err, "Failed to handle request")

	defer rows.Close()
	storage := []entities.DataUser{}
	p := entities.DataUser{}

	for rows.Next() {
		err := rows.Scan(&p.Password, &p.Email)
		CheckError(err, "Failed to copy of the db")

		storage = append(storage, p)
	}

	for _, rng := range storage {
		if rng.Password == ent.Password && rng.Email == ent.Email {
			log.Fatal("The new password matches the old password")
		} else {
			if dbstorage.UpdateDB(&p) {
				fmt.Println("Password success changed !")
			}
		}
	}
}

func GetDataDB() []entities.DataUser {
	Db, err := sql.Open("postgres", Psqlconnect)
	CheckError(err, "Failed to open request")

	err = Db.Ping()
	CheckError(err, "Failed to connect request")

	rows, err := Db.Query("SELECT * FROM storage")
	CheckError(err, "Failed to handle request")

	defer rows.Close()

	storage := []entities.DataUser{}

	for rows.Next() {
		p := entities.DataUser{}
		err = rows.Scan(&p.UserName, &p.Password, &p.Email)
		CheckError(err, "Failed to copy to variables")

		storage = append(storage, p)
	}

	for _, rng := range storage {
		fmt.Println(rng.UserName, rng.Password, rng.Email)
	}

	return storage
}

func CheckError(err error, msg string) {
	if err != nil {
		errors.Wrap(err, msg)
	}
}
