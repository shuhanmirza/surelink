package service

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"log"
	"surelink-go/api/structs"
	"surelink-go/infrastructure"
	"surelink-go/util"
)

type CaptchaService struct {
	cache *infrastructure.Cache
}

func NewCaptchaService(cache *infrastructure.Cache) CaptchaService {
	return CaptchaService{
		cache: cache,
	}
}

func (s CaptchaService) GetNewCaptcha(ctx *gin.Context) (response structs.GetCaptchaResponse, err error) {
	captchaObj, err := s.getCaptchaFromQueue(ctx)
	if err != nil {
		log.Println("error while generating captcha")
		log.Println(err)
		return response, err
	}

	captchaUuid := uuid.New()

	go s.SaveNewCaptcha(ctx, captchaUuid.String(), captchaObj.Val)

	response = structs.GetCaptchaResponse{Uuid: captchaUuid.String(), Img: captchaObj.ImgB64}

	return response, nil
}

func (s CaptchaService) getCaptchaFromQueue(ctx *gin.Context) (captchaObj infrastructure.CaptchaModel, err error) {
	captchaObjJson, err := s.cache.Client.RPop(ctx, util.RedisCaptchaQueueKey).Result()
	if err != nil {
		log.Println("could not get captcha from the queue")
		log.Println(err)
		return captchaObj, &util.CaptchaValidationFailed{}
	}

	err = json.Unmarshal([]byte(captchaObjJson), &captchaObj)
	if err != nil {
		log.Println("invalid captcha object in the queue")
		log.Println(err)
		return captchaObj, &util.CaptchaValidationFailed{}
	}

	return captchaObj, nil
}

func (s CaptchaService) SaveNewCaptcha(ctx *gin.Context, captchaUuid string, captchaStr string) {
	redisKey := util.RedisCaptchaKeyPrefix + captchaUuid
	redisValue := captchaStr
	err := s.cache.Client.Set(ctx, redisKey, redisValue, util.RedisCaptchaTtl).Err()
	if err != nil {
		log.Println(err.Error())
	}
}
