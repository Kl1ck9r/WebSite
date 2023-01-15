package storage

import (
	"database/sql"
	"github.com/cmd/internal/entities"
	"github.com/cmd/internal/utils"
	_ "github.com/lib/pq"
)

var Db *sql.DB

func InsertDB(db *sql.DB, ent *entities.DataUser) (err error) {
	
	tx, err := db.Begin()
	if err != nil {
		return
	}

	defer func() {
		switch err {
		case nil:
			err = tx.Commit()
		default:
			tx.Rollback()
		}
	}()

	Db, err := utils.ConnectDB()
	CheckDB(err)

	defer Db.Close()

	sqlInsert := `INSERT INTO storage (username, password, email) 
    VALUES ($1, $2, $3)`

	_, err = Db.Exec(sqlInsert, ent.UserName, ent.Password, ent.Email)
	CheckDB(err)

	return nil
}

func DeleteDB(db *sql.DB, ent *entities.DataUser) (err error) {

	tx, err := db.Begin()

	if err != nil {
		return
	}

	defer func() {
		switch err {
		case nil:
			err = tx.Commit()
		default:
			tx.Rollback()
		}
	}()
	Db, err := utils.ConnectDB()
	CheckDB(err)

	defer Db.Close()

	sqlDelete := `DELETE FROM storage
		WHERE id=$1 RETURNING *`

	_, err = Db.Exec(sqlDelete, ent.ID)
	CheckDB(err)

	return
}

func UpdateDB(db *sql.DB, ent *entities.DataUser) (err error) {

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
	CheckDB(err)

	sqlUpdate := `UPDATE storage SET password=$1 WHERE email =$2`

	_, err = Db.Exec(sqlUpdate, ent.Password, ent.Email)

	CheckDB(err)
	return
}

func FindUserByID(db *sql.DB, ent *entities.DataUser) (err error) {
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
	CheckDB(err)

	sqlFind := `SELECT username,password,email FROM storage WHERE id=$1`

	_, err = Db.Exec(sqlFind, ent.ID)
	CheckDB(err)

	return
}

func CheckDB(err error) {
	if err != nil {
		panic(err)
	}
}
