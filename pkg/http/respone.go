package http

import (
	"taskgo/pkg/exception"

	"github.com/gin-gonic/gin"
)

func Err(ctx *gin.Context, ex *exception.Exception) {
	ctx.JSON(ex.Code, gin.H{"error": ex.Err.Error(), "message": ex.Message, "details": ex.Details})
}
