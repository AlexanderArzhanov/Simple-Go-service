package dataAccess

import (
	"context"
	"fmt"
	internallogic "simpleGoService/InternalLogic"
	"time"
)

func (PostgresConnection *PostgresConnection) PostgreSQLReadAll() []internallogic.Event {

	var SQLRequest = "select * from events"

	rows, err := PostgresConnection.Connection.Query(context.Background(), SQLRequest)
	if err != nil {

	}

	events := []internallogic.Event{}
	for rows.Next() {

		values, err := rows.Values()
		if err != nil {

		}

		id := values[0].(int32)
		msg := values[1].(string)

		var created_at time.Time
		value2 := values[2]
		if value2 != nil {
			created_at = value2.(time.Time)
		}

		var updated_at time.Time
		value3 := values[3]
		if value3 != nil {
			updated_at = value3.(time.Time)
		}

		event := internallogic.NewEvent(id, msg, created_at, updated_at)

		events = append(events, *event)
	}

	return events
}

func (PostgresConnection *PostgresConnection) PostgreSQLGetSelectionByID(id string) []internallogic.Event {

	SQLRequest := fmt.Sprintf("%s %s", "select * from events where id = ", id)

	rows, err := PostgresConnection.Connection.Query(context.Background(), SQLRequest)
	if err != nil {

	}

	events := []internallogic.Event{}
	for rows.Next() {

		values, err := rows.Values()
		if err != nil {

		}

		id := values[0].(int32)
		msg := values[1].(string)

		var created_at time.Time
		value2 := values[2]
		if value2 != nil {
			created_at = value2.(time.Time)
		}

		var updated_at time.Time
		value3 := values[3]
		if value3 != nil {
			updated_at = value3.(time.Time)
		}

		event := internallogic.NewEvent(id, msg, created_at, updated_at)

		events = append(events, *event)
	}

	return events
}

func (PostgresConnection *PostgresConnection) PostgreSQLDeleteSelectionByID(id string) string {

	SQLRequest := fmt.Sprintf("%s %s", "delete from events where id = ", id)

	_, err := PostgresConnection.Connection.Query(context.Background(), SQLRequest)
	if err != nil {
		return err.Error()
	} else {
		return "data was deleted"
	}
}
