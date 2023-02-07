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

func (PostgresConnection *PostgresConnection) PostgreSQLDeleteEventByID(id string) string {

	SQLRequest := fmt.Sprintf("%s %s", "delete from events where id = ", id)

	_, err := PostgresConnection.Connection.Query(context.Background(), SQLRequest)
	if err != nil {
		return err.Error()
	} else {
		return "data was deleted"
	}
}

func (PostgresConnection *PostgresConnection) PostgreSQLInsertEvent(event *internallogic.Event) string {

	mgg := event.Msg
	created_at := event.Created_at

	SQLRequest := fmt.Sprintf("insert into events (msg, created_at) values (%s, %s)", mgg, created_at)

	_, err := PostgresConnection.Connection.Query(context.Background(), SQLRequest)
	if err != nil {
		return err.Error()
	} else {
		return "data was created"
	}
}
