package db

import (
	"CRUD/common/views"
	"context"
	"log"
)

func (db *DB) CreateUser(usr *views.User) {
	querystring := "INSERT INTO student_user (username,password) VALUES ($1,$2) RETURNING uid"
	err := db.Conn.QueryRow(context.Background(), querystring, usr.UserName, usr.Password).Scan(&usr.UId)
	if err != nil {

		log.Fatal("unable to create user", err)
	}
}
