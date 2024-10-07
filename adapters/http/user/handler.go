package user

import (
	"errors"
	"taskgo/internal/core/user"
	"taskgo/pkg"
	"taskgo/pkg/exception"
	"taskgo/pkg/http"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	service *user.Service
}

func NewHandler(service *user.Service) *UserHandler {
	return &UserHandler{service: service}
}

func (uh *UserHandler) CreateUser(c *gin.Context) {
	var req user.CreateUserRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		ex := exception.New(404, err, err.Error(), err)
		http.Err(c, ex)
		return
	}

	if errs := pkg.Validate(req); len(errs) > 0 {
		ex := exception.New(400, errors.New("validation error"), "validation error", errs)
		http.Err(c, ex)
	}

	if err := uh.service.CreateUser(req); err != nil {
		http.Err(c, err)
		return
	}
	c.JSON(201, gin.H{"message": "user created successfully!"})

}
