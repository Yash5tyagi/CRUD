package db

import (
	"context"
	"fmt"
	"log"
	"os"

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

	Dbdriver := os.Getenv("DB_DRIVER")
	DbHost := os.Getenv("DB_HOST")
	DbUser := os.Getenv("DB_USER")
	DbPassword := os.Getenv("DB_PASSWORD")
	DbName := os.Getenv("DB_NAME")
	DbPort := os.Getenv("DB_PORT")

	connString := fmt.Sprintf("%s://%s:%s@%s:%s/%s?sslmode=disable", Dbdriver, DbUser, DbPassword, DbHost, DbPort, DbName)

	conn, err := pgx.Connect(context.Background(), connString)
	if err != nil {
		log.Fatal(err)
	}
	return conn
}
