package dataAccess

import (
	"context"
	"fmt"

	//"os"

	pgxpool "github.com/jackc/pgx/v4/pgxpool"
)

var PGDBConn PostgresConnection

type PostgresConnection struct {
	dbUrl string

	Connection *pgxpool.Pool
}

func NewPostgresConnection(user string, password string, addr string, port int, dbname string) {

	dbUrl := fmt.Sprintf("postgres://%s:%s@%s:%d/%s", user, password, addr, port, dbname)

	config, err := pgxpool.ParseConfig(dbUrl)
	if err != nil {
		panic(err)
	}

	conn, err := pgxpool.ConnectConfig(context.Background(), config)
	if err != nil {
		panic(err)
	}

	var postgresConnection PostgresConnection
	postgresConnection.dbUrl = dbUrl
	postgresConnection.Connection = conn

	PGDBConn = postgresConnection
}
