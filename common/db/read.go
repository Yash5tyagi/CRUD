package db

import (
	"context"
	"log"

	"github.com/jackc/pgx/v4"
)

func (db *DB) SelectParent() pgx.Rows {
	rows, err := db.Conn.Query(context.Background(), "SELECT father_name,mother_name FROM parents")

	if err != nil {
		log.Fatal(err)
	}

	return rows

}

func (db *DB) SelectStudent() pgx.Rows {
	rows, err := db.Conn.Query(context.Background(), "SELECT first_name,last_name,roll_no FROM students")

	if err != nil {
		log.Fatal(err)
	}

	return rows

}
