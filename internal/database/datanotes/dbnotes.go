package notesdb

import (
	"database/sql"

	"github.com/cmd/internal/entities"
	"github.com/cmd/internal/utils"
	_ "github.com/lib/pq"
	"github.com/pkg/errors"
	"log"
)

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

	sqlInsert := `INSERT INTO notesdb(note)
	VALUES($1,$2)`

	_, err = tx.Exec(sqlInsert, nt.Note, nt.ID)
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

	sqlDelete := `DELETE FROM notesdb WHERE id_note=$1`

	_, err = tx.Exec(sqlDelete, nt.ID)
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

	sqlUpdate := `UPDATE notesdb SET note=$1 where id_note=$2 `

	_, err = tx.Exec(sqlUpdate, nt.Note, nt.ID)
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

	sqlFind := `SELECT note FROM notesdb WHERE id_note=$1`

	_, err = tx.Exec(sqlFind, ent.ID)
	CheckError(err, "Failed to handle request to db")

	return
}

func GetNotes() (sl []entities.Notes, err error) {
	Db, err := utils.ConnectDB()
	CheckError(err, "Failed to open request")

	defer Db.Close()

	rows, err := Db.Query("SELECT note,id_note FROM notesdb")
	CheckError(err, "Failed to handle request")

	defer rows.Close()

	notes := []entities.Notes{}

	for rows.Next() {
		p := entities.Notes{}
		err = rows.Scan(&p.Note, &p.ID)
		CheckError(err, "Failed to copy of database notesdb")

		notes = append(notes, p)
	}

	if err = rows.Err(); err != nil {
		log.Print(err)
	}

	return notes, nil
}

func CheckError(err error, msg string) {
	if err != nil {
		errors.Wrap(err, msg)
	}
}
