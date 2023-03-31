package service

import (
	"github.com/gin-gonic/gin"
	"log"
	"surelink-go/infrastructure"
	"surelink-go/util"
)

type CacheService struct {
	cache *infrastructure.Cache
}

func NewCacheService(cache *infrastructure.Cache) CacheService {
	return CacheService{
		cache: cache,
	}
}

func (s CacheService) SaveNewCaptcha(ctx *gin.Context, captchaUuid string, captchaStr string) {
	redisKey := util.REDIS_CAPTCHA_KEY_PREFIX + captchaUuid
	redisValue := captchaStr
	err := s.cache.Client.Set(ctx, redisKey, redisValue, util.REDIS_CAPTCHA_TTL).Err()
	if err != nil {
		log.Println(err.Error())
	}
}
