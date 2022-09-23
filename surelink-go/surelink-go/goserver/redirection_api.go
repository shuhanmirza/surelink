package goserver

import (
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
	"github.com/lib/pq"
	"log"
	"net/http"
	db "surelink-go/sqlc"
	"surelink-go/util"
)

type GetMapRequest struct {
	Uid string `json:"uid" binding:"required"`
}

type GetMapResponse struct {
	Url string `json:"url"`
}

type SetMapRequest struct {
	CaptchaUuid  string `json:"captcha_uuid" binding:"required"`
	CaptchaValue string `json:"captcha_value" binding:"required"`
	Url          string `json:"url" binding:"required"`
}

type SetMapResponse struct {
	ShortUrl string `json:"short_url"`
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

	url := getRedirectionMapFromRedis(ctx, server, request.Uid)
	if len(url) != 0 {
		response = GetMapResponse{Url: url}
		ctx.JSON(http.StatusOK, response)
		return
	}

	// search in DB
	redirectionMap, err := server.store.Queries.GetRedirectionMap(ctx, request.Uid)
	if err != nil {
		log.Println(err)
		if pqErr, ok := err.(*pq.Error); ok {
			log.Println(pqErr.Code.Name())
		}
		ctx.JSON(http.StatusNotFound, errorResponse(&util.RecordNotFound{}))
		return
	}

	go setRedirectionMapInRedis(ctx, server, redirectionMap.Uid, redirectionMap.Url)

	response = GetMapResponse{Url: redirectionMap.Url}
	ctx.JSON(http.StatusOK, response)
	return

}

func (server *Server) setMap(ctx *gin.Context) {
	var request SetMapRequest
	var response SetMapResponse

	err := ctx.ShouldBindJSON(&request)
	if err != nil {
		log.Println("validation error")
		log.Println(err)
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	if !validateCaptchaFromRedis(ctx, server, request.CaptchaUuid, request.CaptchaValue) {
		log.Println("captcha validation failed")
		ctx.JSON(http.StatusBadRequest, errorResponse(&util.CaptchaValidationFailed{}))
		return
	}

	urlValidity, err := util.IsValidHttpsUrl(request.Url)
	if err != nil {
		log.Println("error while checking url validity")
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	} else if urlValidity == false {
		log.Println("invalid url")
		ctx.JSON(http.StatusBadRequest, errorResponse(&util.UrlParsingError{}))
		return
	}

	shortUrlUid, err := generateShortUrlUid(ctx, server)
	if err != nil {
		log.Println("error while generating unique short url")
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	redirectionMap, err := server.store.Queries.CreateRedirectionMap(ctx, db.CreateRedirectionMapParams{Url: request.Url, Uid: shortUrlUid})
	if err != nil {
		log.Println(err)
		if pqErr, ok := err.(*pq.Error); ok {
			log.Println(pqErr.Code.Name())
		}
		ctx.JSON(http.StatusInternalServerError, &util.UnprecedentedDbError{})
		return
	}

	go setRedirectionMapInRedis(ctx, server, redirectionMap.Uid, redirectionMap.Url)

	response = SetMapResponse{ShortUrl: shortUrlUid}
	ctx.JSON(http.StatusCreated, response)
	return

}

func generateShortUrlUid(ctx *gin.Context, server *Server) (string, error) {
	uidString := util.RandomStringAlphabet(util.SHORT_URL_UID_LENGTH)
	count, err := server.store.Queries.CheckIfUidExists(ctx, uidString)
	if err != nil {
		log.Println(err)
		if pqErr, ok := err.(*pq.Error); ok {
			log.Println(pqErr.Code.Name())
		}
		return "", &util.UnprecedentedDbError{}
	}

	if count > 0 {
		return generateShortUrlUid(ctx, server)
	}

	return uidString, nil
}

func validateCaptchaFromRedis(ctx *gin.Context, server *Server, captchaUuid string, captchaValue string) bool {
	redisKey := util.REDIS_CAPTCHA_KEY_PREFIX + captchaUuid
	redisValue, err := server.redisStore.Client.Get(ctx, redisKey).Result()

	if err == redis.Nil {
		log.Printf("%s key does not exist\n", redisKey)
		return false
	} else if err != nil {
		log.Println(err.Error())
		return false
	}

	if captchaValue == redisValue {
		go server.redisStore.Client.Del(ctx, redisKey) //delete key upon validation
		return true
	}

	return false
}

func getRedirectionMapFromRedis(ctx *gin.Context, server *Server, uid string) string {
	redisKey := util.REDIS_REDIRECTION_KEY_PREFIX + uid
	redisValue, err := server.redisStore.Client.Get(ctx, redisKey).Result()

	if err == redis.Nil {
		log.Printf("%s key does not exist\n", redisKey)
		return ""
	} else if err != nil {
		log.Println(err.Error())
		return ""
	}

	// set default ttl again if used
	go server.redisStore.Client.Expire(ctx, redisKey, util.REDIS_REDIRECTION_TTL)

	return redisValue
}

func setRedirectionMapInRedis(ctx *gin.Context, server *Server, uid string, url string) {
	redisKey := util.REDIS_REDIRECTION_KEY_PREFIX + uid
	redisValue := url
	err := server.redisStore.Client.Set(ctx, redisKey, redisValue, util.REDIS_REDIRECTION_TTL).Err()
	if err != nil {
		log.Println(err.Error())
	}
}
