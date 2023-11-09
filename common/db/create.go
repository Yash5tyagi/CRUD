package db

import (
	"CRUD/common/views"
	"context"
	"fmt"
	"log"
)

func (db *DB) InsertStudent(stud *views.Student) {
	querystring := "INSERT INTO students (first_name,last_name,roll_no) VALUES ($1,$2,$3) ON CONFLICT(roll_no) DO UPDATE SET first_name=$1,last_name=$2 RETURNING id"
	err := db.Conn.QueryRow(context.Background(), querystring, stud.FirstName, stud.LastName, stud.RollNo).Scan(&stud.Id)
	if err != nil {
		fmt.Printf("unable to insert student \n")
		log.Fatal(err)
	}
}

func (db *DB) InsertParent(parent *views.Parents) {
	querystring := "INSERT INTO parents (father_name,mother_name,sid) VALUES ($1,$2,$3) ON CONFLICT(sid) DO UPDATE SET father_name=$1,mother_name=$2"
	_, err := db.Conn.Query(context.Background(), querystring, parent.FatherName, parent.MotherName, parent.SId)
	if err != nil {
		fmt.Printf("unable to insert parents\n")
		log.Fatal(err)

	}
}
