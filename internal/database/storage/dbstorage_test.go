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
		log.Fatal(err)
	}

	defer db.Close()

	var st entities.DataUser
	st.Email = "example@mail.ru"
	st.Password = "Password12345"
	st.UserName = "User-Agent"

	mock.ExpectExec("INSERT INTO storage password,").WithArgs(st.Password, st.UserName, st.Email).WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()

	check, err := utils.ConnectDB()

	if err != nil {
		log.Fatal(err)
	}

	check.Close()

	if err = InsertDB(&st); err != nil {
		t.Errorf("error was not expected while updating stats: %s", err)
	}
}

func TestDeleteDB(t *testing.T) {
	db, mock, err := sqlmock.New()

	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	var st entities.DataUser
	st.ID = "18"

	mock.ExpectExec("DELETE FROM storage WHERE id=$1").WithArgs(st.ID).WillReturnResult(sqlmock.NewResult(1, 1))

	check, err := utils.ConnectDB()

	if err != nil {
		log.Fatal(err)
	}

	check.Close()

	if err = DeleteDB(&st); err != nil {
		t.Errorf("error was not expected while updating stats: %s", err)
	}
}

func TestUpdateDB(t *testing.T) {
	db,mock,err:=sqlmock.New()

	if err!=nil{
		log.Fatal(err)
	}
	
	db.Close()

	var st entities.DataUser
	st.Password="NewPassword"
	st.Email="Anton@134"

	mock.ExpectExec("UPDATE storage SET password=$1 WHERE email =$2").WithArgs(st.Password,st.Email).WillReturnResult(sqlmock.NewResult(1,1))

	if err = UpdateDB(&st);err!=nil{
		t.Errorf("error was not expected while updating stats: %s", err)
	}
}

func TestFindUserBy(t *testing.T) {
	db, mock, err := sqlmock.New()

	if err != nil {
		log.Fatal(err)
	}

	db.Close()

	var st entities.DataUser
	st.ID = "15"

	mock.ExpectExec("SELECT username,password,email FROM storage WHERE id=$1").WithArgs(st.ID).WillReturnResult(sqlmock.NewResult(1, 1))

	check, err := utils.ConnectDB()

	if err != nil {
		log.Fatal(err)
	}
	check.Close()

	if err = FindUserByID(&st); err != nil {
		t.Errorf("error was not expected while updating stats: %s", err)
	}
}
