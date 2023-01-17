package notesdb

import (
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/cmd/internal/entities"
	"github.com/cmd/internal/utils"
	"log"
	"testing"
)

func TestInsertNoteDB(t *testing.T) {
	db, mock, err := sqlmock.New()

	if err != nil {
		log.Fatal("Failed to connect sqlmock database", err)
	}

	defer db.Close()

	var nt entities.Notes
	nt.Note = "Some information"

	mock.ExpectBegin()
	mock.ExpectExec("INSERT INTO notesdb note").WithArgs(nt.Note).WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()

	conndb, err := utils.ConnectDB()

	if err != nil {
		log.Fatal(err)
	}
	defer conndb.Close()

	if err = InsertNoteDB(db, &nt); err != nil {
		t.Errorf("error was not expected while updating stats: %s", err)
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func TestDeleteNoteDB(t *testing.T) {
	db, mock, err := sqlmock.New()

	if err != nil {
		log.Fatal("Failed to connect sqlmock database", err)
	}

	defer db.Close()

	var nt entities.Notes
	nt.ID = "3"

	mock.ExpectBegin()
	mock.ExpectExec("DELETE FROM notesdb WHERE id=$1").WithArgs(nt.ID).WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()

	conndb, err := utils.ConnectDB()

	if err != nil {
		log.Fatal(err)
	}

	defer conndb.Close()

	if err = DeleteNoteDB(db, &nt); err != nil {
		t.Errorf("error was not expected while updating stats: %s", err)
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func TestUpdateNoteDB(t *testing.T) {
	db, mock, err := sqlmock.New()

	if err != nil {
		log.Fatal("Failed to connect sqlmock database", err)
	}

	db.Close()

	conndb, err := utils.ConnectDB()

	if err != nil {
		log.Fatal(err)
	}
	defer conndb.Close()

	var nt entities.Notes
	nt.Note = "New Note"
	nt.ID = "2" // where id

	mock.ExpectBegin()
	mock.ExpectExec("UPDATE notesdb SET password=$1 WHERE email =$2").WithArgs(nt.Note, nt.ID).WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()

	if err = UpdateNoteDB(db, &nt); err != nil {
		t.Errorf("error was not expected while updating stats: %s", err)
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func TestFindRecordByID(t *testing.T) {
	db, mock, err := sqlmock.New()

	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	var nt entities.Notes
	nt.ID = "5"

	mock.ExpectBegin()
	mock.ExpectExec("SELECT note FROM notesdb WHERE id=$1").WithArgs(nt.ID).WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()

	conndb, err := utils.ConnectDB()

	if err != nil {
		log.Fatal(err)
	}
	defer conndb.Close()

	if err = FindRecordByID(db, &nt); err != nil {
		t.Errorf("error was not expected while updating stats: %s", err)
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}
