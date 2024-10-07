package http

import (
	// u "taskgo/adapters/http/user"

	"taskgo/adapters/http/user"

	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	server := gin.Default()
	rg := server.Group("/api")
	user.RegisterRoutes(rg)

	return server
}
