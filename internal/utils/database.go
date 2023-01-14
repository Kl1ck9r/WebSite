package utils

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

const (
	Psqlconnect = "user=postgres dbname=postgres password=Ruslan5655 host=localhost sslmode=disable"
)

func ConnectDB() (*sql.DB, error) {
	db, err := sql.Open("postgres", Psqlconnect)

	if err != nil {
		log.Fatal(err)
	}

	err = db.Ping()

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(db)

	if err != nil {
		log.Printf("ERROR CONNECTING database: \n%v", err)
		return nil, err
	}
	return db, nil
}
