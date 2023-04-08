package service

import (
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
	"github.com/lib/pq"
	"log"
	"net/http"
	"surelink-go/api/structs"
	"surelink-go/infrastructure"
	"surelink-go/infrastructure/sqlc"
	"surelink-go/util"
)

type RedirectionService struct {
	store          *infrastructure.Store
	cache          *infrastructure.Cache
	utilityService *UtilityService
}

func NewRedirectionService(store *infrastructure.Store, cache *infrastructure.Cache, utilityService *UtilityService) RedirectionService {
	return RedirectionService{
		store:          store,
		cache:          cache,
		utilityService: utilityService,
	}
}

func (s RedirectionService) GetMap(ctx *gin.Context, request structs.GetMapRequest) (response structs.GetMapResponse, err error) {
	url := s.getUrlMapFromRedis(ctx, request.Uid)
	if len(url) != 0 {
		response = structs.GetMapResponse{Url: url}
	} else {
		response, err = s.getMapFromStore(ctx, request)
	}

	//TODO: Cache non-existent url map lookup
	go s.incrementRedirectionCount(ctx, request.Uid)

	return response, err
}

func (s RedirectionService) SetMap(ctx *gin.Context, request structs.SetMapRequest) (response structs.SetMapResponse, err error) {
	if !s.validateCaptchaFromRedis(ctx, request.CaptchaUuid, request.CaptchaValue) {
		log.Println("captcha validation failed")
		return response, &util.CaptchaValidationFailed{}
	}

	urlValidity, err := s.utilityService.IsValidHttpsUrl(ctx, request.Url)
	if err != nil {
		log.Println("error while checking url validity")
		return response, err
	} else if urlValidity == false {
		log.Println("invalid url")
		return response, &util.UrlParsingError{}
	}

	shortUrlUid, err := s.generateShortUrlUid(ctx)
	if err != nil {
		log.Println("error while generating unique short url")
		return response, err
	}

	urlMap, err := s.store.Queries.CreateUrlMap(ctx, sqlc.CreateUrlMapParams{Url: request.Url, Uid: shortUrlUid})
	if err != nil {
		log.Println(err)
		if pqErr, ok := err.(*pq.Error); ok {
			log.Println(pqErr.Code.Name())
		}
		ctx.JSON(http.StatusInternalServerError, &util.UnprecedentedDbError{})
		return
	}

	go s.setUrlMapInRedis(ctx, urlMap)

	response = structs.SetMapResponse{ShortUrl: shortUrlUid}
	return response, nil
}

func (s RedirectionService) generateShortUrlUid(ctx *gin.Context) (uidString string, err error) {
	uidString = s.utilityService.RandomStringAlphabet(util.ShortUrlUidLength)

	count, err := s.store.Queries.CheckIfUidExistsInUrlMap(ctx, uidString)
	if err != nil {
		log.Println(err)
		if pqErr, ok := err.(*pq.Error); ok {
			log.Println(pqErr.Code.Name())
		}
		return "", &util.UnprecedentedDbError{}
	}

	if count > 0 {
		return s.generateShortUrlUid(ctx)
	}

	return uidString, nil
}

func (s RedirectionService) validateCaptchaFromRedis(ctx *gin.Context, captchaUuid string, captchaValue string) bool {
	redisKey := util.RedisCaptchaKeyPrefix + captchaUuid
	redisValue, err := s.cache.Client.Get(ctx, redisKey).Result()

	if err == redis.Nil {
		log.Printf("%s key does not exist\n", redisKey)
		return false
	} else if err != nil {
		log.Println(err.Error())
		return false
	} else if captchaValue != redisValue {
		return false
	}

	go s.cache.Client.Del(ctx, redisKey) //delete key upon validation
	return true
}

func (s RedirectionService) getUrlMapFromRedis(ctx *gin.Context, uid string) string {
	redisKey := util.RedisRedirectionKeyPrefix + uid
	redisValue, err := s.cache.Client.Get(ctx, redisKey).Result()

	if err == redis.Nil {
		log.Printf("%s key does not exist in redis\n", redisKey)
		return ""
	} else if err != nil {
		log.Println(err.Error())
		return ""
	}

	return redisValue
}

func (s RedirectionService) getMapFromStore(ctx *gin.Context, request structs.GetMapRequest) (response structs.GetMapResponse, err error) {

	urlMap, err := s.store.Queries.GetUrlMap(ctx, request.Uid)
	if err != nil {
		log.Println(err)
		if pqErr, isPqErr := err.(*pq.Error); isPqErr {
			log.Println(pqErr.Code.Name())
		}
		return response, &util.RecordNotFound{}
	}

	go s.setUrlMapInRedis(ctx, urlMap)

	response = structs.GetMapResponse{Url: urlMap.Url}
	return response, err
}

func (s RedirectionService) setUrlMapInRedis(ctx *gin.Context, urlMap sqlc.UrlMap) {
	redisKey := util.RedisRedirectionKeyPrefix + urlMap.Uid
	redisValue := urlMap.Url
	err := s.cache.Client.Set(ctx, redisKey, redisValue, util.RedisUrlMapTtl).Err()
	if err != nil {
		log.Println(err.Error())
	}
}

func (s RedirectionService) incrementRedirectionCount(ctx *gin.Context, uid string) {
	err := s.store.Queries.IncrementUrlMapTimeRedirected(ctx, uid)
	if err != nil {
		log.Printf("failed to incremeent redirection count for %s\n", uid)
		log.Println(err)
	}
}
