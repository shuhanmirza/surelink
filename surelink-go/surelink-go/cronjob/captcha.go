package cronjob

import (
	"bytes"
	"context"
	"encoding/base64"
	"encoding/json"
	"github.com/afocus/captcha"
	"image/color"
	"image/png"
	"log"
	"surelink-go/api/service"
	"surelink-go/infrastructure"
	"surelink-go/util"
)

type CaptchaCronJob struct {
	cache          *infrastructure.Cache
	utilityService service.UtilityService
}

func NewCaptchaCronJob(cache *infrastructure.Cache, utilityService service.UtilityService) CaptchaCronJob {
	return CaptchaCronJob{
		cache:          cache,
		utilityService: utilityService,
	}
}

func (cj *CaptchaCronJob) Run(ctx context.Context) {
	captchaQueueSize, err := cj.cache.Client.LLen(ctx, util.RedisCaptchaQueueKey).Result()
	if err != nil {
		log.Println("cronjob: error while checking queue size in captcha queue")
		log.Println(err.Error())
		return
	}

	for captchaQueueSize < util.CaptchaQueueMaxSize {
		captchaObj, err := cj.generateCaptchaImage()
		if err != nil {
			log.Println("cronjob: error while generating captcha")
			log.Println(err)
			return
		}

		captchaObjJson, _ := json.Marshal(captchaObj)
		_, err = cj.cache.Client.LPush(ctx, util.RedisCaptchaQueueKey, captchaObjJson).Result()
		if err != nil {
			log.Println("cronjob: error while pushing new captcha in captcha queue")
			log.Println(err.Error())
			return
		} else {
			captchaQueueSize++
		}
	}

}

func (cj *CaptchaCronJob) generateCaptchaImage() (captchaObj infrastructure.CaptchaModel, err error) {

	captchaGenerator := captcha.New()

	pathIndex := cj.utilityService.RandomInt(0, int64(len(util.CaptchaFontPathList)-1))
	err = captchaGenerator.SetFont(util.CaptchaFontPathList[pathIndex])
	if err != nil {
		log.Println("error occurred while setting font" + err.Error())
		return captchaObj, &util.FontNotFound{}
	}
	captchaGenerator.SetSize(128, 64)
	captchaGenerator.SetDisturbance(captcha.NORMAL)
	captchaGenerator.SetFrontColor(color.RGBA{R: 255, G: 255, B: 255, A: 255}) // white
	captchaGenerator.SetBkgColor(
		color.RGBA{R: 255, A: 255}, //red
		color.RGBA{B: 255, A: 255}, //blue
		color.RGBA{G: 153, A: 255}) //light-green

	captchaImg, captchaStr := captchaGenerator.Create(util.CaptchaTextLength, captcha.ALL)

	var buff bytes.Buffer
	err = png.Encode(&buff, captchaImg)
	if err != nil {
		log.Println("error while png encoding" + err.Error())
		return captchaObj, &util.ImgEncodingFailed{}
	}
	captchaImgB64 := base64.StdEncoding.EncodeToString(buff.Bytes())

	return infrastructure.CaptchaModel{Val: captchaStr, ImgB64: captchaImgB64}, nil
}
