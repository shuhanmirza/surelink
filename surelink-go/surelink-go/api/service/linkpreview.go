package service

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"log"
	"surelink-go/api/structs"
	"surelink-go/infrastructure"
	"surelink-go/util"
)

type LinkPreviewService struct {
	cache *infrastructure.Cache
}

func NewLinkPreviewService(cache *infrastructure.Cache) LinkPreviewService {
	return LinkPreviewService{
		cache: cache,
	}
}

func (s *LinkPreviewService) GetLinkPreview(ctx *gin.Context, request structs.GetLinkPreviewRequest) (response structs.GetLinkPreviewResponse, err error) {
	response, err = s.getLinkPreviewFromCache(ctx, request)
	if err == nil {
		return response, nil
	}

	return response, err
}

func (s *LinkPreviewService) getLinkPreviewFromCache(ctx *gin.Context, request structs.GetLinkPreviewRequest) (response structs.GetLinkPreviewResponse, err error) {
	key := util.RedisLinkPreviewPrefix + request.Uid
	valueJson, err := s.cache.Client.Get(ctx, key).Result()
	if err != nil {
		log.Println("could not get link preview from cache")
		log.Println(err)
		return response, err
	}

	err = json.Unmarshal([]byte(valueJson), &request)
	if err != nil {
		log.Printf("invalid link preview response in cache entry %s \n", key)
		log.Println(err)
		return response, err
	}

	return response, nil
}
