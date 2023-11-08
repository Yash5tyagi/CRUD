package db

import (
	"context"
	"log"

	"github.com/google/uuid"
)

func (db *DB) DelStudent(id uuid.UUID) {
	QueryString := "DELETE FROM students WHERE id=$1"
	_, err := db.Conn.Query(context.Background(), QueryString, id)
	if err != nil {
		log.Fatal(err)
	}
}
