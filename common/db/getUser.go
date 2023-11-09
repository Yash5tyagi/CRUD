package db

import (
	"context"
	"log"

	"github.com/jackc/pgx/v4"
)

func (db *DB) SelectUser(username string) pgx.Rows {
	rows, err := db.Conn.Query(context.Background(), "SELECT uid,username,password FROM student_user WHERE username=$1", username)

	if err != nil {
		log.Fatal(err)
	}

	return rows
}
