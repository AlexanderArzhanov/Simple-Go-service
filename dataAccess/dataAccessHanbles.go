package dataAccess

import (
	internallogic "simpleGoService/InternalLogic"
)

func ReadAll() []internallogic.Event {
	return PGDBConn.PostgreSQLReadAll()
}
