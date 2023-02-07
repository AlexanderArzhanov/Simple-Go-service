package dataAccess

import (
	internallogic "simpleGoService/InternalLogic"
)

func ReadAll() []internallogic.Event {
	return PGDBConn.PostgreSQLReadAll()
}

func GetEventByID(id string) []internallogic.Event {
	return PGDBConn.PostgreSQLGetSelectionByID(id)
}

func DeleteEventByID(id string) string {
	return PGDBConn.PostgreSQLDeleteEventByID(id)
}

func CreateEvent(event *internallogic.Event) string {
	return PGDBConn.PostgreSQLInsertEvent(event)
}
