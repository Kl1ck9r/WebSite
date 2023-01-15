package storage

import (
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/cmd/internal/entities"
	"github.com/cmd/internal/utils"
	"log"
	"testing"
)

func TestInsertDB(t *testing.T) {
	db, mock, err := sqlmock.New()

	if err != nil {
		log.Fatal("Failed to connect sqlmock database",err)
	}

	defer db.Close()

	var st entities.DataUser
	st.Email = "example@mail.ru"
	st.Password = "Password12345"
	st.UserName = "User-Agent"

	mock.ExpectBegin()
	mock.ExpectExec("INSERT INTO storage password,username,email").WithArgs(st.Password, st.UserName, st.Email).WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()

	check, err := utils.ConnectDB()

	if err != nil {
		log.Fatal(err)
	}

	defer check.Close()

	if err = InsertDB(db,&st); err != nil {
		t.Errorf("error was not expected while updating stats: %s", err)
	}
}

func TestDeleteDB(t *testing.T) {
	db, mock, err := sqlmock.New()

	if err != nil {
		log.Fatal("Failed to connect sqlmock database",err)
	}

	defer db.Close()

	var st entities.DataUser
	st.ID = "18"

	mock.ExpectExec("DELETE FROM storage WHERE id=$1").WithArgs(st.ID).WillReturnResult(sqlmock.NewResult(1, 1))

	conndb, err := utils.ConnectDB()

	if err != nil {
		log.Fatal(err)
	}

	defer conndb.Close()

	if err = DeleteDB(db,&st); err != nil {
		t.Errorf("error was not expected while updating stats: %s", err)
	}
}

func TestUpdateDB(t *testing.T) {
	db,mock,err:=sqlmock.New()

	if err!=nil{
		log.Fatal("Failed to connect sqlmock database",err)
	}
	
	db.Close()

	conndb, err := utils.ConnectDB()

	if err != nil {
		log.Fatal(err)
	}

	defer conndb.Close()

	var st entities.DataUser
	st.Password="NewPassword"
	st.Email="Anton@134"

	mock.ExpectExec("UPDATE storage SET password=$1 WHERE email =$2").WithArgs(st.Password,st.Email).WillReturnResult(sqlmock.NewResult(1,1))

	if err = UpdateDB(db,&st);err!=nil{
		t.Errorf("error was not expected while updating stats: %s", err)
	}
}

func TestFindUserBy(t *testing.T) {
	db, mock, err := sqlmock.New()

	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	var st entities.DataUser
	st.ID = "15"

	mock.ExpectExec("SELECT username,password,email FROM storage WHERE id=$1").WithArgs(st.ID).WillReturnResult(sqlmock.NewResult(1, 1))

	conndb, err := utils.ConnectDB()

	if err != nil {
		log.Fatal(err)
	}
	defer conndb.Close()

	if err = FindUserByID(db,&st); err != nil {
		t.Errorf("error was not expected while updating stats: %s", err)
	}
}
