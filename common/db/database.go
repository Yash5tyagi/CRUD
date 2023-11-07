package db

import "github.com/jackc/pgx/v4"

type DB struct {
	conn *pgx.Conn
}

func New(conn *pgx.Conn) (db *DB) {
	return &DB{
		conn: conn,
	}

}

func (db *DB) GetParent() {

}

func (db *DB) GetStudent() {

}
