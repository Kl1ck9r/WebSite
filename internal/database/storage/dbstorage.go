package storage

import (
	"database/sql"
	"github.com/cmd/internal/entities"
	_ "github.com/lib/pq"
)

var Db *sql.DB

const (
	Psqlconnect = "user=postgres dbname=postgres password=Ruslan5655 host=localhost sslmode=disable"
)

func InsertDB(ent *entities.DataUser) {

	Db, err := sql.Open("postgres", Psqlconnect)
	CheckDB(err)

	err = Db.Ping()
	CheckDB(err)

	defer Db.Close()

	sqlInsert := `INSERT INTO storage (username, password, email) 
    VALUES ($1, $2, $3)`

	_, err = Db.Exec(sqlInsert, ent.UserName, ent.Password, ent.Email)
	CheckDB(err)
}

func DeleteDB(ent *entities.DataUser) {

	Db, err := sql.Open("postgres", Psqlconnect)
	CheckDB(err)

	err = Db.Ping()
	CheckDB(err)

	defer Db.Close()

	sqlDelete := `DELETE FROM storage(username,password,email)
		WHERE username=$1,password=$2,email=$3)`

	_, err = Db.Exec(sqlDelete,ent.UserName,ent.Password,ent.Email)
	CheckDB(err)
}

func UpdateDB(ent *entities.DataUser)bool{
	Db, err := sql.Open("postgres", Psqlconnect)
	CheckDB(err)

	err = Db.Ping()
	CheckDB(err)

	sqlUpdate:=`UPDATE storage SET password=$1 WHERE email =$2`

	_,err=Db.Exec(sqlUpdate,ent.Password,ent.Email)
	
	CheckDB(err)

	return true 
}

func CheckDB(err error) {
	if err != nil {
		panic(err)
	}
}
