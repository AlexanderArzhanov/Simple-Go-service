package dataAccess

import (
	"context"
	"fmt"

	//"os"

	pgxpool "github.com/jackc/pgx/v4/pgxpool"
)

func PostgreSQLReadAll() string {

	pgdbUrl := "postgres://admin:password@localhost:5432/pgdb"

	config, err := pgxpool.ParseConfig(pgdbUrl)
	if err != nil {
		return fmt.Sprintf("%s:%s", "Unable to connect to database: ", err)
	}

	conn, err := pgxpool.ConnectConfig(context.Background(), config)

	defer conn.Close()

	var SQLRequest = "select * from myschema.events"

	rows, err := conn.Query(context.Background(), SQLRequest)
	if err != nil {

	}

	var dataResult string
	for rows.Next() {

		values, err := rows.Values()
		if err != nil {

		}

		id := values[0].(int32)
		firstName := values[1].(string)
		lastName := values[2].(string)
		email := values[3].(string)
		age := values[4].(int32)

		rowData := fmt.Sprintf("%d %s %s %s %d\n", id, firstName, lastName, email, age)

		dataResult = fmt.Sprintf("%s %s\n", dataResult, rowData)
	}

	return dataResult
}
