// api_gin.go

package rest

import (
	"errors"

	"github.com/gin-gonic/gin"
)

type GinAPIService struct {
	Engine *gin.Engine
}

func (g *GinAPIService) Handle(method, path string, handler gin.HandlerFunc) error {
	methodMap := map[string]func(string, ...gin.HandlerFunc) gin.IRoutes{
		"GET":     g.Engine.GET,
		"POST":    g.Engine.POST,
		"PUT":     g.Engine.PUT,
		"PATCH":   g.Engine.PATCH,
		"DELETE":  g.Engine.DELETE,
		"OPTIONS": g.Engine.OPTIONS,
		"HEAD":    g.Engine.HEAD,
	}

	if methodFunc, exists := methodMap[method]; exists {
		methodFunc(path, handler)
		return nil
	}

	return errors.New("unsupported HTTP method")
}
