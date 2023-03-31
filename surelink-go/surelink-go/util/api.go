package util

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func ErrorResponse(err error) gin.H {
	return gin.H{
		"error": err.Error(),
	}
}

func HandleCommonValidationError(ctx *gin.Context, err error) {
	fmt.Println("validation error")
	fmt.Println(err)
	ctx.JSON(http.StatusBadRequest, ErrorResponse(err))
}
