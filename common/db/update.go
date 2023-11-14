package db

import (
	"CRUD/common/views"
	"context"
	"log"
)

func (db *DB) UpdateStudent(stud *views.Student) {
	QueryString := "UPDATE students SET first_name=$1 ,last_name=$2 ,pid=$3 where roll_no=$4"
	_, err := db.Conn.Query(context.Background(), QueryString, stud.FirstName, stud.LastName, stud.PId, stud.RollNo)
	if err != nil {
		log.Fatal(err)
	}
}

func (db *DB) UpdateParent(parent *views.Parents) {
	QueryString := "UPDATE parents SET father_name=$1 ,mother_name=$2 where pid=$3"
	_, err := db.Conn.Query(context.Background(), QueryString, parent.FatherName, parent.MotherName, parent.PId)
	if err != nil {
		log.Fatal(err)
	}
}
