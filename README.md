# gogean/rest

[![CodeQL](https://github.com/gogean/rest/actions/workflows/codeql.yml/badge.svg)](https://github.com/gogean/rest/actions/workflows/codeql.yml)
[![Build](https://github.com/gogean/rest/actions/workflows/go.yml/badge.svg)](https://github.com/gogean/rest/actions/workflows/go.yml)
[![Go Report Card](https://goreportcard.com/badge/github.com/gogean/rest)](https://goreportcard.com/report/github.com/gogean/rest)
[![GoDoc](https://pkg.go.dev/badge/github.com/gogean/rest?status.svg)](https://pkg.go.dev/github.com/gogean/rest?tab=doc)

This library automatically generates REST APIs from a given configuration. 

> It is currently a work in progress; below is a sample configuration.


### rest.yml
```
apis:
  - name: "Foo"
    path: "/foo"
    method: "GET"
    description: "Get an item"
    query:
      schema: "foo_query_schema.json"
    actions:
      - type: "db"
        query: foo.sql.tpl

  - name: "Bar"
    path: "/bar"
    method: "POST"
    description: "Creates an item"
    payload:
      schema: "bar_payload_schema.json"

```

### foo.sql.tpl
```
SELECT       
	*
FROM 
	users
LIMIT
	{{.start}}, {{.size}}
```
### foo_query_schema.json
```
{
  "$schema": "http://json-schema.org/draft-07/schema#",
  "type": "object",
  "properties": {
    "size": {
      "type": "number",
      "minimum": 0,
      "maximum": 1000
    },
    "start": {
      "type": "number",
      "minimum": 0,
      "maximum": 5
    }
  },
  "required": ["size", "start"]
}

```

### Example use

```
// main.go

package main

import (
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/gogean/rest"
)

func main() {
	dbConnectionString := fmt.Sprintf(
		"%s:%s@(%s:%s)/%s",
		os.Getenv("APP_DB_USER"),
		os.Getenv("APP_DB_PWD"),
		os.Getenv("APP_DB_HOST"),
		os.Getenv("APP_DB_PORT"),
		os.Getenv("APP_DB_NAME"))

	db, _ := rest.SQLConnect(dbConnectionString)

	dbConnection := rest.SQL{Connection: db}

	// Create a Gin engine instance
	router := gin.Default()

	// Create an instance of GinAPIService
	ginAPIService := rest.GinAPIService{Engine: router}

	// Define an api object
	apiConfig, _ := rest.GetAPIConfig("../rest/test-config/rest.yml")
	apis := rest.GetAPIs(apiConfig)

	// Use generateAPI to add a handler to the Gin engine
	rest.GenerateAPIs(apis, &ginAPIService, &dbConnection)

	// Start the Gin server
	router.Run(":8080")
}

```



