package service

import (
	"bytes"
	"encoding/base64"
	"github.com/afocus/captcha"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"image/color"
	"image/png"
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
	captchaStr, captchaImgB64, err := generateCaptchaImage()
	if err != nil {
		log.Println("error while generating captcha")
		log.Println(err)
		return response, err
	}

	captchaUuid := uuid.New()

	go s.SaveNewCaptcha(ctx, captchaUuid.String(), captchaStr)

	response = structs.GetCaptchaResponse{Uuid: captchaUuid.String(), Img: captchaImgB64}

	return response, nil
}

func generateCaptchaImage() (captchaStr string, captchaImgB64 string, err error) {

	captchaGenerator := captcha.New()

	err = captchaGenerator.SetFont(util.FONT_COMIC_PATH)
	if err != nil {
		log.Println("error occurred while setting font" + err.Error())
		return "", "", &util.FontNotFound{}
	}
	captchaGenerator.SetSize(128, 64)
	captchaGenerator.SetDisturbance(captcha.MEDIUM)
	captchaGenerator.SetFrontColor(color.RGBA{R: 255, G: 255, B: 255, A: 255}) // white
	captchaGenerator.SetBkgColor(
		color.RGBA{R: 255, A: 255}, //red
		color.RGBA{B: 255, A: 255}, //blue
		color.RGBA{G: 153, A: 255}) //light-green

	captchaImg, captchaStr := captchaGenerator.Create(util.CAPTCHA_TEXT_LENGTH, captcha.ALL)

	var buff bytes.Buffer
	err = png.Encode(&buff, captchaImg)
	if err != nil {
		log.Println("error while png encoding" + err.Error())
		return "", "", &util.ImgEncodingFailed{}
	}
	captchaImgB64 = base64.StdEncoding.EncodeToString(buff.Bytes())

	return captchaStr, captchaImgB64, nil

}

func (s CaptchaService) SaveNewCaptcha(ctx *gin.Context, captchaUuid string, captchaStr string) {
	redisKey := util.REDIS_CAPTCHA_KEY_PREFIX + captchaUuid
	redisValue := captchaStr
	err := s.cache.Client.Set(ctx, redisKey, redisValue, util.REDIS_CAPTCHA_TTL).Err()
	if err != nil {
		log.Println(err.Error())
	}
}
