package utils

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/cmd/internal/entities"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func Init() {
	if err := godotenv.Load(".env"); err != nil {
		log.Print("No .env file found")
	}
}

func ConnectDB() (*sql.DB, error) {

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
		log.Fatal("Failed to open database", err)
	}

	err = db.Ping()

	if err != nil {
		log.Printf("ERROR CONNECTING database: \n%v", err)
		return nil, err
	}

	return db, nil
}
