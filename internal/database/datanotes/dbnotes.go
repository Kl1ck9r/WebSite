package notesdb

import (
	"database/sql"

	"github.com/cmd/internal/entities"
	"github.com/cmd/internal/utils"
	_ "github.com/lib/pq"
	"github.com/pkg/errors"
)

var Db *sql.DB

func InsertNoteDB(db *sql.DB, nt *entities.Notes) (err error) {

	tx, err := db.Begin()

	if err != nil {
		return
	}

	defer func() {
		switch err {
		case nil:
			tx.Commit()
		default:
			tx.Rollback()
		}
	}()

	Db, err := utils.ConnectDB()
	CheckError(err, "Failed to open db")

	sqlInsert := `INSERT INTO notesdb(note)
	VALUES($1,$2)`

	_, err = Db.Exec(sqlInsert, nt.Note, nt.ID)
	CheckError(err, "Failed to handle request db")

	return
}

func DeleteNoteDB(db *sql.DB, nt *entities.Notes) (err error) {

	tx, err := db.Begin()

	if err != nil {
		return
	}

	defer func() {
		switch err {
		case nil:
			tx.Commit()
		default:
			tx.Rollback()
		}
	}()

	Db, err := utils.ConnectDB()
	CheckError(err, "Failed to open db")

	sqlDelete := `DELETE FROM notesdb WHERE id_note=$1`

	_, err = Db.Exec(sqlDelete, nt.ID)
	CheckError(err, "Failed to handle request from db")

	return
}

func UpdateNoteDB(db *sql.DB, nt *entities.Notes) (err error) {

	tx, err := db.Begin()

	if err != nil {
		return
	}

	defer func() {
		switch err {
		case nil:
			tx.Commit()
		default:
			tx.Rollback()
		}
	}()

	Db, err := utils.ConnectDB()
	CheckError(err, "Failed open db")

	sqlUpdate := `UPDATE notesdb SET note=$1 where id_note=$2 `

	_, err = Db.Exec(sqlUpdate, nt.Note, nt.ID)
	CheckError(err, "Failed to handle request from db")

	return
}

func FindRecordByID(db *sql.DB, ent *entities.Notes) (err error) {

	tx, err := db.Begin()

	if err != nil {
		return
	}

	defer func() {
		switch err {
		case nil:
			tx.Commit()
		default:
			tx.Rollback()
		}
	}()

	Db, err := utils.ConnectDB()
	CheckError(err, "Failed to open db")

	sqlFind := `SELECT note FROM notesdb WHERE id_note=$1`

	_, err = Db.Exec(sqlFind, ent.ID)
	CheckError(err, "Failed to handle request to db")

	return
}

func CheckError(err error, msg string) {
	if err != nil {
		errors.Wrap(err, msg)
	}
}
