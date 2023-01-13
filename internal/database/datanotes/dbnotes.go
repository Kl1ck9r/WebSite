package notesdb

import (
	"database/sql"

	"github.com/cmd/internal/entities"
	_ "github.com/lib/pq"
	"github.com/pkg/errors"
)

var Db *sql.DB

const (
	Psqlconnect = "user=postgres dbname=postgres password=Ruslan5655 host=localhost sslmode=disable"
)

func InsertNoteDB(nt *entities.Notes) {
	Db, err := sql.Open("postgres", Psqlconnect)
	CheckError(err, "Failed to open db")

	err = Db.Ping()
	CheckError(err, "Failed to connect db")

	sqlInsert := `INSERT INTO notesdb(note)
	VALUES($1,$2)`

	_, err = Db.Exec(sqlInsert, nt.Note, nt.ID)
	CheckError(err, "Failed to handle request db")
}

func DeleteNoteDB(nt *entities.Notes) {
	Db, err := sql.Open("postgres", Psqlconnect)
	CheckError(err, "Failed to open db")

	err = Db.Ping()
	CheckError(err, "Failed to connect db")

	sqlDelete := `DELETE FROM notesdb WHERE id_note=$1`

	_, err = Db.Exec(sqlDelete, nt.ID)
	CheckError(err, "Failed to handle request from db")
}

func UpdateNoteDB(nt *entities.Notes) {
	Db, err := sql.Open("postgres", Psqlconnect)
	CheckError(err, "Failed open db")

	err = Db.Ping()
	CheckError(err, "Failed to connect db")

	sqlUpdate := `UPDATE notesdb SET note=$1 where id_note=$2 `

	_, err = Db.Exec(sqlUpdate, nt.Note, nt.ID)
	CheckError(err, "Failed to handle request from db")
}

func CheckError(err error, msg string) {
	if err != nil {
		errors.Wrap(err, msg)
	}
}
