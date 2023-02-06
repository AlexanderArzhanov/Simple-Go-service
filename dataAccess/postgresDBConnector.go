package dataAccess

import (
    "context"
    //"fmt"

    //"os"

    pgxpool "github.com/jackc/pgx/v4/pgxpool"
)

var PGDBConn PostgresConnection

type PostgresConnection struct {
    dbUrl string

    Connection *pgxpool.Pool
}

func NewPostgresConnection() {

    dbUrl := "postgres://admin:password@localhost:5432/pgdb"

    config, err := pgxpool.ParseConfig(dbUrl)
    if err != nil {
        //return fmt.Sprintf("%s:%s", "Unable to connect to database: ", err)
    }

    conn, err := pgxpool.ConnectConfig(context.Background(), config)
    if err != nil {
        //return fmt.Sprintf("%s:%s", "Unable to connect to database: ", err)
    }

    var postgresConnection PostgresConnection
    postgresConnection.dbUrl = dbUrl
    postgresConnection.Connection = conn

    PGDBConn = postgresConnection
}

