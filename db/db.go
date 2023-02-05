package db

import pgxpool "github.com/jackc/pgx/v4/pgxpool"

type DB struct {
	dbType   string
	user     string
	password string
	addr     string
	port     int32
	dbName   string

	dbConn *pgxpool.Conn
}

func NewDB() {

}

func (db *DB) Connect() {

}
