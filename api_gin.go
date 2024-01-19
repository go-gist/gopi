// api_gin.go

package restql

import (
	"errors"

	"github.com/gin-gonic/gin"
)

// GinAPIService is a concrete implementation of the APIService interface using the Gin Engine.
type GinAPIService struct {
	Engine *gin.Engine
}

// Handle implements the Handle method of the APIService interface.
func (g *GinAPIService) Handle(method, path string, handler gin.HandlerFunc) error {
	switch method {
	case "GET":
		g.Engine.GET(path, handler)
	case "POST":
		g.Engine.POST(path, handler)
	default:
		return errors.New("unsupported HTTP method")
	}

	return nil
}
