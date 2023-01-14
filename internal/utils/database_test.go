package utils

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"testing"

	"github.com/cmd/internal/entities"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/stretchr/testify/assert"
)

func TestInit(t *testing.T) {
	if err := godotenv.Load(".env"); err != nil {
		assert.Nil(t, err)
	}
}

func TestConnectDB(t *testing.T) {

	psqlConnect := entities.DataBase{
		Host:     os.Getenv("DB_HOST"),
		Password: os.Getenv("DB_PASSWORD"),
		UserName: os.Getenv("DB_NAME"),
		DBName:   os.Getenv("DB_USERNAME"),
		Port:     os.Getenv("DB_PORT"),
		SSLMode:  "disable",
	}

	db, err := sql.Open("postgres", fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		psqlConnect.Host, psqlConnect.Port, psqlConnect.UserName, psqlConnect.Password, psqlConnect.DBName, psqlConnect.SSLMode))

	if err != nil {
		log.Fatal("Failed top open database", err)
	}

	defer db.Close()

	err = db.Ping()

	if err != nil {
		log.Fatal("Failed to connect to database", err)
	}
}
