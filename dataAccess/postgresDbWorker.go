package dataAccess

import (
	"context"
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
