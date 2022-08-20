package goserver

import (
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
	"github.com/lib/pq"
	"log"
	"net/http"
	"surelink-go/util"
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

	url := getEntryFromRedis(ctx, server, request.UUID)
	if len(url) != 0 {
		response = GetMapResponse{Url: url}
		ctx.JSON(http.StatusOK, response)
		return
	}

	// search in DB
	redirectionMap, err := server.store.Queries.GetRedirectionMap(ctx, request.UUID)
	if err != nil {
		log.Println(err)
		if pqErr, ok := err.(*pq.Error); ok {
			log.Println(pqErr.Code.Name())
		}
		ctx.JSON(http.StatusNotFound, errorResponse(&util.RecordNotFound{}))
		return
	}

	//TODO: go routine
	setEntryInRedis(ctx, server, request.UUID, redirectionMap.Url)

	response = GetMapResponse{Url: redirectionMap.Url}
	ctx.JSON(http.StatusOK, response)
	return

}

func getEntryFromRedis(ctx *gin.Context, server *Server, uuid string) string {
	redisKey := util.REDIS_REDIRECTION_KEY_PREFIX + uuid
	redisValue, err := server.redisStore.Client.Get(ctx, redisKey).Result()

	if err == redis.Nil {
		log.Printf("%s key does not exist\n", redisKey)
		return ""
	} else if err != nil {
		log.Println(err.Error())
		return ""
	}

	// set default ttl again if used
	server.redisStore.Client.Expire(ctx, redisKey, util.REDIS_REDIRECTION_TTL)

	return redisValue
}

func setEntryInRedis(ctx *gin.Context, server *Server, uuid string, url string) {
	redisKey := util.REDIS_REDIRECTION_KEY_PREFIX + uuid
	redisValue := url
	err := server.redisStore.Client.Set(ctx, redisKey, redisValue, util.REDIS_REDIRECTION_TTL).Err()
	if err != nil {
		log.Println(err.Error())
	}
}
