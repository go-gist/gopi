package rest

import (
	"database/sql"
	"fmt"
)

type dbConnection interface {
	Query(query string, args ...interface{}) (*sql.Rows, error)
}

func handleDBOperation(action action, data map[string]interface{}, db dbConnection) ([]map[string]interface{}, error) {
	dbQuery, err := executeTemplate(action.queryTemplate, data)
	if err != nil {
		return nil, fmt.Errorf("failed to execute template: %w", err)
	}

	// Log the generated query
	fmt.Println("Generated SQL Query:", dbQuery)

	rows, err := db.Query(dbQuery)
	if err != nil {
		return nil, fmt.Errorf("query execution error: %w", err)
	}
	defer rows.Close()

	// Get column names
	columns, err := rows.Columns()
	if err != nil {
		return nil, fmt.Errorf("failed to get columns: %w", err)
	}

	// Prepare a slice to hold the column values
	values := make([]interface{}, len(columns))
	valuePtrs := make([]interface{}, len(columns))
	for i := range columns {
		valuePtrs[i] = &values[i]
	}

	// Create a slice to hold the results
	results := []map[string]interface{}{}
	for rows.Next() {
		// Scan the row into the valuePtrs
		if err := rows.Scan(valuePtrs...); err != nil {
			return nil, fmt.Errorf("failed to scan row: %w", err)
		}

		// Convert row values to a map
		rowMap := make(map[string]interface{})
		for i, colName := range columns {
			val := values[i]
			// Handle NULL values
			if val == nil {
				rowMap[colName] = nil
			} else {
				switch v := val.(type) {
				case []byte:
					rowMap[colName] = string(v)
				default:
					rowMap[colName] = v
				}
			}
		}
		results = append(results, rowMap)
	}

	// Check for errors from iterating over rows
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error occurred while iterating rows: %w", err)
	}

	// Check if any records were retrieved
	if len(results) == 0 {
		fmt.Println("No records found")
	} else {
		fmt.Printf("Number of records retrieved: %d\n", len(results))
	}

	return results, nil
}
