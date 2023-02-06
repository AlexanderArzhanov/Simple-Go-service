package dataAccess

import (
	"context"
	"fmt"

	//"os"

	//pgxpool "github.com/jackc/pgx/v4/pgxpool"
)

func (PostgresConnection *PostgresConnection) PostgreSQLReadAll() string {	

    //defer postgresConnection.Connection.Close()

	var SQLRequest = "select * from myschema.events"

    rows, err := PostgresConnection.Connection.Query(context.Background(), SQLRequest)
	if err != nil {

	}

	var dataResult string
	for rows.Next() {

		values, err := rows.Values()
		if err != nil {

		}

		id := values[0].(int32)


		//rowData := fmt.Sprintf("%d %s %s %s %d\n", id, firstName, lastName, email, age)
        rowData := fmt.Sprintf("%d\n", id)

		dataResult = fmt.Sprintf("%s %s\n", dataResult, rowData)
	}

	return dataResult
}
