package service

import (
	"bytes"
	"encoding/json"
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
	store            *infrastructure.Store
	cache            *infrastructure.Cache
	serviceDiscovery *ServiceSilo
	secretConfig     util.SecretConfig
	httpClient       http.Client
}

func NewRedirectionService(store *infrastructure.Store, cache *infrastructure.Cache, secretConfig util.SecretConfig, serviceDiscovery *ServiceSilo) RedirectionService {
	return RedirectionService{
		store:            store,
		cache:            cache,
		secretConfig:     secretConfig,
		serviceDiscovery: serviceDiscovery,
		httpClient:       http.Client{},
	}
}

func (s *RedirectionService) GetMap(ctx *gin.Context, request structs.GetMapRequest) (response structs.GetMapResponse, err error) {
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

func (s *RedirectionService) SetMapV1(ctx *gin.Context, request structs.SetMapRequest) (response structs.SetMapResponse, err error) {
	if !s.validateCaptchaFromRedis(ctx, request.CaptchaUuid, request.CaptchaValue) {
		log.Println("captcha validation failed")
		return response, &util.CaptchaValidationFailed{}
	}

	return s.setMap(ctx, request.Url)
}

func (s *RedirectionService) SetMapV2(ctx *gin.Context, request structs.SetMapRequestV2) (response structs.SetMapResponse, err error) {

	if !s.validateRecaptcha(s.secretConfig.RecaptchaSecretKey, request.RecaptchaToken) {
		log.Println("captcha validation failed")
		return response, &util.CaptchaValidationFailed{}
	}

	return s.setMap(ctx, request.Url)
}

func (s *RedirectionService) setMap(ctx *gin.Context, url string) (response structs.SetMapResponse, err error) {
	urlValidity, err := s.serviceDiscovery.UtilityService().IsValidHttpsUrl(ctx, url)
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

	urlMap, err := s.store.Queries.CreateUrlMap(ctx, sqlc.CreateUrlMapParams{Url: url, Uid: shortUrlUid})
	if err != nil {
		log.Println(err)
		if pqErr, ok := err.(*pq.Error); ok {
			log.Println(pqErr.Code.Name())
		}
		ctx.JSON(http.StatusInternalServerError, &util.UnprecedentedDbError{})
		return
	}

	go s.setUrlMapInRedis(ctx, urlMap)
	go s.serviceDiscovery.LinkPreviewService().SetLinkPreviewOnLinkCreation(ctx, url)

	response = structs.SetMapResponse{ShortUrl: shortUrlUid}
	return response, nil
}

func (s *RedirectionService) generateShortUrlUid(ctx *gin.Context) (uidString string, err error) {
	uidString = s.serviceDiscovery.UtilityService().RandomStringAlphabet(util.ShortUrlUidLength)

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

func (s *RedirectionService) validateCaptchaFromRedis(ctx *gin.Context, captchaUuid string, captchaValue string) bool {
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

func (s *RedirectionService) getUrlMapFromRedis(ctx *gin.Context, uid string) string {
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

func (s *RedirectionService) getMapFromStore(ctx *gin.Context, request structs.GetMapRequest) (response structs.GetMapResponse, err error) {

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

func (s *RedirectionService) setUrlMapInRedis(ctx *gin.Context, urlMap sqlc.UrlMap) {
	redisKey := util.RedisRedirectionKeyPrefix + urlMap.Uid
	redisValue := urlMap.Url
	err := s.cache.Client.Set(ctx, redisKey, redisValue, util.RedisUrlMapTtl).Err()
	if err != nil {
		log.Println(err.Error())
	}
}

func (s *RedirectionService) incrementRedirectionCount(ctx *gin.Context, uid string) {
	err := s.store.Queries.IncrementUrlMapTimeRedirected(ctx, uid)
	if err != nil {
		log.Printf("failed to incremeent redirection count for %s\n", uid)
		log.Println(err)
	}
}

// TODO: move all external api calls to utility service
func (s *RedirectionService) validateRecaptcha(secretKey string, token string) (isHuman bool) {

	externalServiceUrl := "https://www.google.com/recaptcha/api/siteverify?secret=" + secretKey + "&response=" + token

	httpRequest, err := http.NewRequest("POST", externalServiceUrl, bytes.NewBuffer([]byte{}))
	httpRequest.Header.Add("Content-Type", "application/json")

	httpResponse, err := s.httpClient.Do(httpRequest)
	if err != nil {
		log.Println("failed to verify recaptcha")
		log.Println(err)
		return false
	}
	defer httpResponse.Body.Close()

	var recaptchaResponse structs.RecaptchaResponse
	err = json.NewDecoder(httpResponse.Body).Decode(&recaptchaResponse)
	if err != nil {
		log.Println("failed to decode resp.Body from google.com recaptcha")
		log.Println(err)
		return false
	}

	return recaptchaResponse.Score >= 0.3
}
