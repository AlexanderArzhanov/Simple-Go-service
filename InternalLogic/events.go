package internallogic

import "time"

type Event struct {
	Id         int32     `json:"id"`
	Msg        string    `json:"msg"`
	Created_at time.Time `json:"created_at"`
	Updated_at time.Time `json:"updated_at"`
}

func NewEvent(id int32, msg string, created_at time.Time, updated_at time.Time) *Event {
	event := new(Event)

	event.Id = id
	event.Msg = msg
	event.Created_at = created_at
	event.Updated_at = updated_at

	return event
}
