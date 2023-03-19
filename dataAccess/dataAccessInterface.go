package dataAccess

import internallogic "simpleGoService/InternalLogic"

var ServiceDataConn ServiceData

type ServiceData interface {
	Connect(map[string]string)
	ReadAll() []internallogic.Event
	GetEventByID(string) []internallogic.Event
	DeleteEventByID(string) string
	CreateEvent(*internallogic.Event) string
	UpdateEvent(*internallogic.Event) string
}

func NewConnection() *PostgresConnection {
	return &PostgresConnection{}
}
