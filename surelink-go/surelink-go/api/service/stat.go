package service

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
	"github.com/lib/pq"
	"log"
	"surelink-go/api/structs"
	"surelink-go/infrastructure"
	"surelink-go/util"
)

type StatService struct {
	cache *infrastructure.Cache
	store *infrastructure.Store
}

func NewStatService(cache *infrastructure.Cache, store *infrastructure.Store) StatService {
	return StatService{
		cache: cache,
		store: store,
	}
}

func (s StatService) GetHomePageStats(ctx *gin.Context) (response structs.GetHomePageStatsResponse, err error) {
	var redisModel infrastructure.HomePageStatModel

	redisModelStr, err := s.cache.Client.Get(ctx, util.RedisServiceStatKey).Result()
	if err == redis.Nil {
		return s.getHomePageStatFromStore(ctx)
	} else if err != nil {
		log.Print("redis lookup issue")
		log.Println(err)

		return response, &util.UnprecedentedDbError{}
	}

	err = json.Unmarshal([]byte(redisModelStr), &redisModel)
	if err != nil {
		log.Println("invalid value in redis for home stat data")
		log.Println(err)
		return response, &util.UnprecedentedDbError{}
	}

	response = structs.GetHomePageStatsResponse{NumUrlMapCreatedLifetime: redisModel.NumUrlMapCreatedLifetime,
		NumUrlMapRedirectedLifetime: redisModel.NumUrlMapRedirectedLifetime}

	return response, nil
}

func (s StatService) getHomePageStatFromStore(ctx *gin.Context) (response structs.GetHomePageStatsResponse, err error) {
	numUrlMap, err := s.store.Queries.GetUrlMapCount(ctx)
	if err != nil {
		log.Println(err)
		if pqErr, ok := err.(*pq.Error); ok {
			log.Println(pqErr.Code.Name())
		}
		return response, &util.UnprecedentedDbError{}
	}

	numUrlMapRedirection, err := s.store.Queries.GetUrlMapRedirectionCount(ctx)
	if err != nil {
		log.Println(err)
		if pqErr, ok := err.(*pq.Error); ok {
			log.Println(pqErr.Code.Name())
		}
		return response, &util.UnprecedentedDbError{}
	}

	response = structs.GetHomePageStatsResponse{NumUrlMapCreatedLifetime: numUrlMap,
		NumUrlMapRedirectedLifetime: numUrlMapRedirection}

	go s.saveHomePageStatInCache(ctx, response)

	return response, nil
}
func (s StatService) saveHomePageStatInCache(ctx *gin.Context, response structs.GetHomePageStatsResponse) {
	redisModel := infrastructure.HomePageStatModel{
		NumUrlMapRedirectedLifetime: response.NumUrlMapRedirectedLifetime,
		NumUrlMapCreatedLifetime:    response.NumUrlMapCreatedLifetime}

	err := s.cache.Client.Set(ctx, util.RedisServiceStatKey, redisModel, util.RedisServiceStatTtl).Err()
	if err != nil {
		log.Println("failed to save homePageStat in redis")
		log.Println(err)
		return
	}
}
