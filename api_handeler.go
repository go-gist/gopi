package restql

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func generateHandler(api api) gin.HandlerFunc {
	return func(c *gin.Context) {
		queryParams := parseQueryParams(c.Request.URL.Query())

		if api.Query != nil && api.Query.Schema != "" {
			if errors := validateJSON(queryParams, api.Query.Schema); len(errors) > 0 {
				responseError(c, http.StatusBadRequest, errors, queryParams)
				return
			}
		}

		params, err := bindRequestBody(c)
		if len(params) > 0 && err != nil {
			responseError(c, http.StatusBadRequest, []ValidationError{{
				Key:     "body",
				Message: err.Error(),
			}}, params)
			return
		}

		if api.Payload != nil && api.Payload.Schema != "" {
			if errors := validateJSON(params, api.Payload.Schema); len(errors) > 0 {
				responseError(c, http.StatusBadRequest, errors, params)
				return
			}
		}

		if api.DB != nil {
			handleDatabaseOperation()
		}

		responseData := generateResponseData(params)
		c.JSON(http.StatusOK, responseData)
	}
}

func parseQueryParams(queryValues map[string][]string) map[string]interface{} {
	queryParams := make(map[string]interface{})
	for key, values := range queryValues {
		value := values[0]
		switch {
		case isInt(value):
			intValue, _ := strconv.Atoi(value)
			queryParams[key] = intValue
		case isFloat(value):
			floatValue, _ := strconv.ParseFloat(value, 64)
			queryParams[key] = floatValue
		case isBool(value):
			boolValue, _ := strconv.ParseBool(value)
			queryParams[key] = boolValue
		default:
			queryParams[key] = value
		}
	}
	return queryParams
}

func isInt(value string) bool {
	_, err := strconv.Atoi(value)
	return err == nil
}

func isFloat(value string) bool {
	_, err := strconv.ParseFloat(value, 64)
	return err == nil
}

func isBool(value string) bool {
	switch value {
	case "true", "false":
		return true
	default:
		return false
	}
}

func bindRequestBody(c *gin.Context) (map[string]interface{}, error) {
	var params map[string]interface{}
	err := c.ShouldBindJSON(&params)
	return params, err
}

func handleDatabaseOperation() {
	fmt.Println()
}
