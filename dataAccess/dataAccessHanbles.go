package dataAccess

func ReadAll() string {
    return PGDBConn.PostgreSQLReadAll()
}