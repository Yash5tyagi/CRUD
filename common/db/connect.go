package db

import (
	"context"
	"log"

	"github.com/jackc/pgx/v4"
)

type DB struct {
	Conn *pgx.Conn
}

func New(conn *pgx.Conn) (db *DB) {
	return &DB{
		Conn: conn,
	}

}

func Init() *pgx.Conn {
	connString := "postgres://yash:yash1234@localhost:5432/student_info?sslmode=disable"

	conn, err := pgx.Connect(context.Background(), connString)
	if err != nil {
		log.Fatal(err)
	}
	return conn
}
