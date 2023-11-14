package db

import (
	"CRUD/common/views"
	"context"
	"log"
)

func (db *DB) InsertParent(parent *views.Parents) {
	querystring := `WITH ins AS (
		INSERT INTO parents (father_name,mother_name) VALUES ($1,$2) ON CONFLICT(father_name,mother_name) DO NOTHING RETURNING pid
	)
	SELECT pid FROM ins
	UNION ALL
	SELECT pid FROM parents WHERE father_name = $1 AND mother_name = $2
	LIMIT 1`

	err := db.Conn.QueryRow(context.Background(), querystring, parent.FatherName, parent.MotherName).Scan(&parent.PId)
	if err != nil {
		log.Fatal("unable to insert parents ", err)

	}
}

func (db *DB) InsertStudent(stud *views.Student) {
	querystring := "INSERT INTO students (first_name,last_name,roll_no,pid) VALUES ($1,$2,$3,$4) ON CONFLICT(roll_no) DO UPDATE SET first_name=$1,last_name=$2 RETURNING id"
	err := db.Conn.QueryRow(context.Background(), querystring, stud.FirstName, stud.LastName, stud.RollNo, stud.PId).Scan(&stud.Id)
	if err != nil {
		log.Fatal("unable to insert student ", err)
	}
}
