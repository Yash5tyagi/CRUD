package db

import (
	"CRUD/common/views"
	"context"
	"log"
)

func (db *DB) CreateUser(usr *views.User) {
	querystring := "INSERT INTO student_user (username,password) VALUES ($1,$2) ON CONFLICT(username) DO NOTHING "
	_, err := db.Conn.Query(context.Background(), querystring, usr.UserName, usr.Password)
	if err != nil {

		log.Fatal("unable to create user", err)
	}
}
