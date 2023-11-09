package db

import (
	"CRUD/common/views"
	"context"
	"log"
)

func (db *DB) InsertStudent(stud *views.Student) {
	querystring := "INSERT INTO students (first_name,last_name,roll_no) VALUES ($1,$2,$3) ON CONFLICT(roll_no) DO UPDATE SET first_name=$1,last_name=$2 RETURNING id"
	err := db.Conn.QueryRow(context.Background(), querystring, stud.FirstName, stud.LastName, stud.RollNo).Scan(&stud.Id)
	if err != nil {
		log.Fatal("unable to insert student ", err)
	}
}

func (db *DB) InsertParent(parent *views.Parents) {
	querystring := "INSERT INTO parents (father_name,mother_name,sid) VALUES ($1,$2,$3) ON CONFLICT(sid) DO UPDATE SET father_name=$1,mother_name=$2"
	_, err := db.Conn.Query(context.Background(), querystring, parent.FatherName, parent.MotherName, parent.SId)
	if err != nil {
		log.Fatal("unable to insert parents ", err)

	}
}
