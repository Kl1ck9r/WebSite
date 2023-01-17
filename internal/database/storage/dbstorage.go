package storage

import (
	"database/sql"

	"github.com/cmd/internal/entities"
	_ "github.com/lib/pq"
)

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

	sqlInsert := `INSERT INTO storage (username, password, email) 
    VALUES ($1, $2, $3)`

	_, err = tx.Exec(sqlInsert, ent.UserName, ent.Password, ent.Email)
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

	sqlDelete := `DELETE FROM storage
		WHERE id=$1 RETURNING *`

	_, err = tx.Exec(sqlDelete, ent.ID)
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

	sqlUpdate := `UPDATE storage SET password=$1 WHERE email =$2`

	_, err = tx.Exec(sqlUpdate, ent.Password, ent.Email)

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

	sqlFind := `SELECT username,password,email FROM storage WHERE id=$1`

	_, err = tx.Exec(sqlFind, ent.ID)
	CheckDB(err)

	return
}

func CheckDB(err error) {
	if err != nil {
		panic(err)
	}
}
