package service

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"surelink-go/api/structs"
	"surelink-go/infrastructure"
	"surelink-go/util"
)

type LinkPreviewService struct {
	cache              *infrastructure.Cache
	redirectionService RedirectionService
	secretConfig       util.SecretConfig
}

func NewLinkPreviewService(cache *infrastructure.Cache, redirectionService RedirectionService, secretConfig util.SecretConfig) LinkPreviewService {
	return LinkPreviewService{
		cache:              cache,
		redirectionService: redirectionService,
		secretConfig:       secretConfig,
	}
}

func (s *LinkPreviewService) GetLinkPreview(ctx *gin.Context, request structs.GetLinkPreviewRequest) (response structs.GetLinkPreviewResponse, err error) {

	urlMapResponse, err := s.redirectionService.GetMap(ctx, structs.GetMapRequest{Uid: request.Uid})
	if err != nil {
		return response, err
	}

	response, err = s.getLinkPreviewFromCache(ctx, urlMapResponse.Url)
	if err == nil {
		return response, nil
	}

	response, err = s.getLinkPreviewFromLinkPreviewDotNet(urlMapResponse.Url)
	if err == nil {
		go s.setLinkPreviewInCache(ctx, urlMapResponse.Url, response)
	}
	return response, err

}

func (s *LinkPreviewService) setLinkPreviewInCache(ctx *gin.Context, url string, response structs.GetLinkPreviewResponse) {
	key := util.RedisLinkPreviewPrefix + url
	valueBytes, _ := json.Marshal(response)
	value := string(valueBytes)

	err := s.cache.Client.Set(ctx, key, value, util.RedisLinkPreviewTtl).Err()
	if err != nil {
		log.Println(err.Error())
	}
}

func (s *LinkPreviewService) getLinkPreviewFromCache(ctx *gin.Context, url string) (response structs.GetLinkPreviewResponse, err error) {
	key := util.RedisLinkPreviewPrefix + url
	valueJson, err := s.cache.Client.Get(ctx, key).Result()
	if err != nil {
		log.Println("could not get link preview from cache")
		log.Println(err)
		return response, err
	}

	err = json.Unmarshal([]byte(valueJson), &response)
	if err != nil {
		log.Printf("invalid link preview response in cache entry %s \n", key)
		log.Println(err)
		return response, err
	}

	return response, nil
}

func (s *LinkPreviewService) getLinkPreviewFromLinkPreviewDotNet(url string) (response structs.GetLinkPreviewResponse, err error) {
	externalServiceUrl := "https://api.linkpreview.net/?key=" + s.secretConfig.LinkPreviewApiKey + "&q=" + url

	resp, err := http.Get(externalServiceUrl)
	if err != nil {
		log.Println("failed to retrieve preview from linkpreview.net")
		log.Println(err)
		return response, &util.LinkPreviewNotFoundError{}
	}
	defer resp.Body.Close()

	err = json.NewDecoder(resp.Body).Decode(&response)
	if err != nil {
		log.Println("failed to decode resp.Body from linkpreview.net")
		log.Println(err)
		return response, &util.LinkPreviewNotFoundError{}
	}

	return response, nil
}
