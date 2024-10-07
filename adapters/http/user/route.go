package user

import (
	"taskgo/adapters/repositories/memory"
	u "taskgo/internal/core/user"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(rg *gin.RouterGroup) *gin.RouterGroup {

	userService := u.NewService(memory.NewUserRepo())
	uh := NewHandler(userService)
	ur := rg.Group("/users")
	{
		ur.POST("", uh.CreateUser)
	}
	return ur
}
