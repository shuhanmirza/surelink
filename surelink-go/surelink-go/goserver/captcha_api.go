package goserver

import (
	"bytes"
	"encoding/base64"
	"github.com/afocus/captcha"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"image/color"
	"image/png"
	"log"
	"net/http"
	"surelink-go/util"
)

type GetCaptchaResponse struct {
	Uuid string `json:"uuid"`
	Img  string `json:"img"`
}

func (server *Server) getCaptcha(ctx *gin.Context) {
	var response GetCaptchaResponse

	captchaStr, captchaImgB64, err := generateCaptchaImage()
	if err != nil {
		log.Println("error while generating captcha")
		log.Println(err)
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
	}

	captchaUuid := uuid.New()

	go enterCaptchaToRedis(ctx, server, captchaUuid.String(), captchaStr)

	response = GetCaptchaResponse{Uuid: captchaUuid.String(), Img: captchaImgB64}
	ctx.JSON(http.StatusOK, response)
	return

}

func enterCaptchaToRedis(ctx *gin.Context, server *Server, captchaUuid string, captchaStr string) {
	redisKey := util.REDIS_CAPTCHA_KEY_PREFIX + captchaUuid
	redisValue := captchaStr
	err := server.redisStore.Client.Set(ctx, redisKey, redisValue, util.REDIS_CAPTCHA_TTL).Err()
	if err != nil {
		log.Println(err.Error())
	}
}

func generateCaptchaImage() (string, string, error) {

	captchaGenerator := captcha.New()

	err := captchaGenerator.SetFont(util.FONT_COMIC_PATH)
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

	captchaImg, captchaStr := captchaGenerator.Create(util.CAPTHCA_TEXT_LENGTH, captcha.ALL)

	var buff bytes.Buffer
	err = png.Encode(&buff, captchaImg)
	if err != nil {
		log.Println("error while png encoding" + err.Error())
		return "", "", &util.ImgEncodingFailed{}
	}
	captchaImgB64 := base64.StdEncoding.EncodeToString(buff.Bytes())

	return captchaStr, captchaImgB64, nil

}
