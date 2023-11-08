package db

import (
	"CRUD/common/views"
	"context"
	"log"
)

func (db *DB) DelStudent(stud *views.Student) {
	QueryString := "DELETE FROM students WHERE id=$1"
	_, err := db.Conn.Query(context.Background(), QueryString, stud.Id)
	if err != nil {
		log.Fatal(err)
	}
}
