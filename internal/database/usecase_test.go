package usecase

import (
	"fmt"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/cmd/internal/entities"
	"github.com/cmd/internal/utils"
	"log"
	"testing"
)

// just simple testing

func TestWebSiteAccess(t *testing.T) {

	personallyData := entities.DataUser{
		Password: "Ruslan12345",
		Email:    "ruslan@mail.ru",
		UserName: "Ruslan",
	}

	if WebsiteAccess(&personallyData) {
		fmt.Println("Login completed successfully")
	} else {
		t.Errorf("Site access denied")
	}
}

func TestExistsUser(t *testing.T) {

}

func TestChangePassword(t *testing.T) {

}

func TestGetDataDB(t *testing.T) {
	db, mock, err := sqlmock.New()

	if err != nil {
		log.Fatal("Failed to connect sqlmock database", err)
	}
	defer db.Close()

	conndb, err := utils.ConnectDB()

	if err != nil {
		log.Fatal("Faield to connect to database storage", err)
	}

	defer conndb.Close()

	mock.ExpectExec("SELECT username,password,email FROM storage")

	storage := GetDataDB()

	for _, rng := range storage {
		if rng.Email == "" && rng.Password == "" && rng.UserName == "" {
			t.Errorf("Failed to get some data from database storage")
		}
	}
}

func TestGetByUserName(t *testing.T) {

}
