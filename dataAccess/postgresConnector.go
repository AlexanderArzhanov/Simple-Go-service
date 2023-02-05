package dataAccess

import (
    "context"
    "fmt"
    "os"

    pgx "github.com/jackc/pgx"
)

func PostgreSQL_ReadAll()  {

    pgdbUrl := "postgres://admin:password@localhost:5432/pgdb"
    conn, err := pgx.Connect()    .   Connect   (context.Background(), pgdbUrl)
    if err != nil{
        fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
        os.Exit(1)
    }

  //  defer conn.Close(context.Background())

   // var SQLRequest = "select * from pg_catalog.pg_aggregate"

   // Result :=




    //pgx.
}