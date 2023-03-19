package dataAccess

import (
	"context"
	"fmt"
	internallogic "simpleGoService/InternalLogic"
	"time"

	pgxpool "github.com/jackc/pgx/v4/pgxpool"
)

type PostgresConnection struct {
	dbUrl      string
	Connection *pgxpool.Pool
}

func (PostgresConn *PostgresConnection) Connect(connData map[string]string) {

	user := connData["user"]
	password := connData["password"]
	addr := connData["addr"]
	port := connData["port"]
	dbname := connData["dbname"]

	dbUrl := fmt.Sprintf("postgres://%s:%s@%s:%s/%s", user, password, addr, port, dbname)

	config, err := pgxpool.ParseConfig(dbUrl)
	if err != nil {
		panic(err)
	}

	conn, err := pgxpool.ConnectConfig(context.Background(), config)
	if err != nil {
		panic(err)
	}

	PostgresConn.dbUrl = dbUrl
	PostgresConn.Connection = conn
}

func (PostgresConn *PostgresConnection) ReadAll() []internallogic.Event {

	var SQLRequest = "select * from events"

	rows, err := PostgresConn.Connection.Query(context.Background(), SQLRequest)
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

func (PostgresConn *PostgresConnection) GetEventByID(id string) []internallogic.Event {

	SQLRequest := fmt.Sprintf("%s %s", "select * from events where id = ", id)

	rows, err := PostgresConn.Connection.Query(context.Background(), SQLRequest)
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

func (PostgresConn *PostgresConnection) DeleteEventByID(id string) string {

	SQLRequest := fmt.Sprintf("%s %s", "delete from events where id = ", id)

	_, err := PostgresConn.Connection.Query(context.Background(), SQLRequest)
	if err != nil {
		return err.Error()
	} else {
		return "data was deleted"
	}
}

func (PostgresConn *PostgresConnection) CreateEvent(event *internallogic.Event) string {

	created_at := time.Now().Format("2006-01-02 15:04:05")

	SQLRequest := fmt.Sprintf("insert into events (msg, created_at) values ('%s', '%s')", event.Msg, created_at)
	_, err := PostgresConn.Connection.Query(context.Background(), SQLRequest)
	if err != nil {
		return err.Error()
	} else {
		return "data was created"
	}
}

func (PostgresConn *PostgresConnection) UpdateEvent(event *internallogic.Event) string {

	updated_at := time.Now().Format("2006-01-02 15:04:05")

	SQLRequest := fmt.Sprintf("UPDATE events SET msg='%s', updated_at='%s' WHERE id=%d", event.Msg, updated_at, event.Id)
	_, err := PostgresConn.Connection.Query(context.Background(), SQLRequest)
	if err != nil {
		return err.Error()
	} else {
		return "data was updated"
	}
}
