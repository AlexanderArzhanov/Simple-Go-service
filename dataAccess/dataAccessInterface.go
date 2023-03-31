package dataAccess

import internallogic "simpleGoService/InternalLogic"

var ServiceDataConn ServiceDataConnection

type ServiceDataConnection interface {
	Connect(map[string]string)
	ReadAll() []internallogic.Event
	GetEventByID(string) []internallogic.Event
	DeleteEventByID(string) string
	CreateEvent(*internallogic.Event) string
	UpdateEvent(*internallogic.Event) string
}

func NewConnection() ServiceDataConnection {
	/*var SDConn ServiceDataConnection //
	SDConn = &PostgresConnection{}
	return SDConn*/

	return &PostgresConnection{}
}
