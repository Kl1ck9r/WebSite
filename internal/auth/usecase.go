package usecase

import (
	"log"

	dbstorage "github.com/cmd/internal/database/storage"
	"github.com/cmd/internal/entities"
	"github.com/cmd/internal/utils"
	_ "github.com/lib/pq"
	"github.com/pkg/errors"
)

func WebsiteAccess(ent *entities.DataUser) bool {
	var IsTrue bool = false

	storage := GetDataDB()

	for _, rng := range storage {
		if rng.Email == ent.Email && rng.Password == ent.Password && rng.UserName == ent.UserName {
			IsTrue = true
		}
	}

	return IsTrue
}

func ExistsUser(ent *entities.DataUser) bool {
	Db, err := utils.ConnectDB()
	CheckError(err, "Failed to open db")

	defer Db.Close()

	rows, err := Db.Query("SELECT email FROM storage")
	CheckError(err, "Failed to hanld request")

	defer rows.Close()

	storage := []entities.DataUser{}

	for rows.Next() {
		p := entities.DataUser{}
		err := rows.Scan(&p.Email)
		CheckError(err, "Failed to copy of the db")

		storage = append(storage, p)
	}

	var IsTrue bool = false

	for _, rng := range storage {
		if rng.Email == ent.Email {
			IsTrue = true
		}
	}

	return IsTrue
}

func ChangePassword(ent *entities.DataUser) (err error) {
	Db, err := utils.ConnectDB()
	CheckError(err, "Failed to open  db")

	defer Db.Close()

	rows, err := Db.Query("SELECT password,email FROM storage")
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
			dbstorage.UpdateDB(Db, &p)
		}
	}

	return
}

func GetDataDB() []entities.DataUser {
	Db, err := utils.ConnectDB()
	CheckError(err, "Failed to open request")

	defer Db.Close()

	rows, err := Db.Query("SELECT username,password,email FROM storage")
	CheckError(err, "Failed to handle request")

	defer rows.Close()

	storage := []entities.DataUser{}

	for rows.Next() {
		p := entities.DataUser{}
		err = rows.Scan(&p.UserName, &p.Password, &p.Email)
		CheckError(err, "Failed to copy to variables")

		storage = append(storage, p)
	}

	if err = rows.Err(); err != nil {
		log.Print(err)
	}

	return storage
}

func GetByUserName() (sl []string, err error) {
	Db, err := utils.ConnectDB()
	CheckError(err, "Failed to open request")

	rows, err := Db.Query("SELECT username FROM storage")
	CheckError(err, "Faield to handle request db")

	defer rows.Close()

	username := []string{}
	var names string

	for rows.Next() {
		err = rows.Scan(&names)
		CheckError(err, "Failed to copy to variable")
		username = append(username, names)
	}

	return username, nil
}

func CheckError(err error, msg string) {
	if err != nil {
		errors.Wrap(err, msg)
	}
}
