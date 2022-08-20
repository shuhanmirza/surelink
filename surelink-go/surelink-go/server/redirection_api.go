package server

import (
	"github.com/gin-gonic/gin"
	"github.com/lib/pq"
	"log"
	"net/http"
)

type GetMapRequest struct {
	UUID string `json:"uuid" binding:"required"`
}

type GetMapResponse struct {
	Url string `json:"url"`
}

func (server *Server) getMap(ctx *gin.Context) {
	var request GetMapRequest
	var response GetMapResponse

	err := ctx.ShouldBindJSON(&request)
	if err != nil {
		log.Println("validation error")
		log.Println(err)
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	redirectionMap, err := server.store.Queries.GetRedirectionMap(ctx, request.UUID)
	if err != nil {
		pq.ErrorCode.Name

		if pqErr, ok := err.(*pq.Error); ok {
			log.Println(pqErr.Code.Name())
		}
		ctx.JSON(http.StatusNotFound, errorResponse(err))
		return
	}

	response = GetMapResponse{Url: redirectionMap.Url}
	ctx.JSON(http.StatusOK, response)
	return

}
