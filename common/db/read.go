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
	rows, err := db.Conn.Query(context.Background(), "SELECT first_name,last_name,roll_no,pid FROM students")

	if err != nil {
		log.Fatal(err)
	}

	return rows

}

func (db *DB) SelectAllDet() pgx.Rows {
	QueryString := "select students.first_name,students.last_name,parents.father_name,parents.mother_name,students.roll_no from  students inner join parents on parents.pid=students.pid;"
	rows, err := db.Conn.Query(context.Background(), QueryString)

	if err != nil {
		log.Fatal(err)
	}

	return rows
}

func (db *DB) SelectPid(father_name string, mother_name string) pgx.Rows {
	rows, err := db.Conn.Query(context.Background(), "SELECT pid FROM parents where father_name=$1,mother_name=$2", father_name, mother_name)

	if err != nil {
		log.Fatal(err)
	}

	return rows

}
