package db

import (
	"CRUD/common/views"
	"context"
	"fmt"
	"log"
)

func (db *DB) InsertStudent(stud *views.Student) {
	querystring := "INSERT INTO students (first_name,last_name,roll_no) VALUES ($1,$2,$3) ON CONFLICT(roll_no) DO NOTHING RETURNING id"
	err := db.Conn.QueryRow(context.Background(), querystring, stud.FirstName, stud.LastName, stud.RollNo).Scan(&stud.Id)
	if err != nil {
		fmt.Printf("unable to insert student \n")
		log.Fatal(err)
	}
}

func (db *DB) InsertParent(parent *views.Parents) {
	querystring := "INSERT INTO parents (father_name,mother_name,sid) VALUES ($1,$2,$3) ON CONFLICT(sid) DO NOTHING"
	_, err := db.Conn.Query(context.Background(), querystring, parent.FatherName, parent.MotherName, parent.SId)
	if err != nil {
		fmt.Printf("unable to insert parents\n")
		log.Fatal(err)

	}
}
