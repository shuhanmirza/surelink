package service

import (
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
	"github.com/lib/pq"
	"log"
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
		return response, nil
	}

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
	return response, nil
}

func (s RedirectionService) getUrlMapFromRedis(ctx *gin.Context, uid string) string {
	redisKey := util.REDIS_REDIRECTION_KEY_PREFIX + uid
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

func (s RedirectionService) setUrlMapInRedis(ctx *gin.Context, urlMap sqlc.UrlMap) {
	redisKey := util.REDIS_REDIRECTION_KEY_PREFIX + urlMap.Uid
	redisValue := urlMap.Url
	err := s.cache.Client.Set(ctx, redisKey, redisValue, 0).Err()
	if err != nil {
		log.Println(err.Error())
	}
}
