// restql.go

package restql

import (
	"github.com/gin-gonic/gin"
)

// APIPath is a constant representing the default API path.
const APIPath = "/api"

// APIService defines the interface for a service that can handle API requests.
type APIService interface {
	Handle(method, path string, handler gin.HandlerFunc)
}

// GinAPIService is a concrete implementation of the APIService interface using the Gin Engine.
type GinAPIService struct {
	Engine *gin.Engine
}

// Handle implements the Handle method of the APIService interface.
func (g GinAPIService) Handle(method, path string, handler gin.HandlerFunc) {
	switch method {
	case "GET":
		g.Engine.GET(path, handler)
	case "POST":
		g.Engine.POST(path, handler)
	// Add more cases for other HTTP methods if needed.
	default:
		// Handle unsupported methods if necessary.
	}
}

// API defines the structure for an API object.
type API struct {
	Path   string // Making the 'Path' field public
	Method string // Making the 'Method' field public
}

// GenerateAPI takes an API object containing properties such as path, method, etc., and an APIService object.
// It adds a handler specific to the method (defaulting to GET if not provided).
// For example, it adds APIService.Handle for the provided method.
func GenerateAPI(api interface{}, service APIService) {
	// Retrieve the method and path from the API object; default to "GET" if not provided.
	method := api.(API).Method
	if method == "" {
		method = "GET"
	}
	path := api.(API).Path
	if path == "" {
		path = APIPath
	}

	// Call the Handle method on APIService for the specified HTTP method.
	service.Handle(method, path, generateHandler(api))
}

// generateHandler is a helper function to generate the appropriate handler function based on the API object.
func generateHandler(api interface{}) gin.HandlerFunc {
	// Implement the handler logic based on the API object.
	return func(c *gin.Context) {
		// Customize the handler logic here based on the provided API object.
		// You can access the API object properties such as api.(API).Path, api.(API).Method, etc.
		c.JSON(200, gin.H{
			"message": "Handler for " + api.(API).Method + " " + api.(API).Path,
		})
	}
}
