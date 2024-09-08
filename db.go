package restql

import "database/sql"

type dbConnection interface {
	Query(query string, args ...interface{}) (*sql.Rows, error)
}

func handleDBOperation(action action, data map[string]interface{}, db dbConnection) error {
	dbQuery, _ := executeTemplate(action.queryTemplate, data)
	_, err := db.Query(dbQuery)
	if err != nil {
		return err
	}
	logInfo(dbQuery)
	return nil
}
