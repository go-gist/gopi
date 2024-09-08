package restql

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

type SQL struct {
	Connection *sql.DB
}

func SQLConnect(connectionString string) (*sql.DB, error) {
	db, err := sql.Open("mysql", connectionString)
	if err != nil {
		logError("Unable to connect to database", connectionString, err.Error())
		return nil, err
	}

	if err := db.Ping(); err != nil {
		logError("Unable to connect to database", connectionString, err.Error())
		db.Close()
		return nil, err
	}

	db.SetMaxOpenConns(99)

	return db, nil
}

func SQLDisconnect(db *sql.DB) {
	db.Close()
}

// Query executes a query and returns rows and an error
func (db *SQL) Query(query string, args ...interface{}) (*sql.Rows, error) {
	return db.Connection.Query(query, args)
}
