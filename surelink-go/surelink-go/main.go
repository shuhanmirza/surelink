package main

import (
	"context"
	"database/sql"
	_ "github.com/lib/pq"
	"github.com/robfig/cron/v3"
	"log"
	"math/rand"
	"os"
	"surelink-go/api/controller"
	"surelink-go/api/routes"
	"surelink-go/api/service"
	_ "surelink-go/api/service"
	"surelink-go/cronjob"
	"surelink-go/infrastructure"
	"surelink-go/util"
	"time"
)

func main() {
	initialTests()

	globalConfig, err := util.LoadGlobalConfig(".")
	if err != nil {
		log.Fatal("can not load global config", err)
	}

	secretConfig, err := util.LoadSecretConfig(".")
	if err != nil {
		log.Fatal("can not load secret config", err)
	}

	//miscellaneous
	random := rand.New(rand.NewSource(time.Now().UnixNano()))
	cronScheduler := cron.New()
	go cronScheduler.Run()
	defer cronScheduler.Stop()

	//database and cache
	conn, err := sql.Open(globalConfig.DBDriver, globalConfig.DBSource)
	if err != nil {
		log.Fatal("can't connect to the database", err)
	}
	store := infrastructure.NewStore(conn)

	cache := infrastructure.NewCache(globalConfig.RedisUrl)

	// initialize gin router
	log.Println("Initializing Routes")
	ginRouter := infrastructure.NewGinRouter()

	//initialize service
	utilityService := service.NewUtilityService(cache, random)

	// captcha
	captchaService := service.NewCaptchaService(cache)
	captchaController := controller.NewCaptchaController(captchaService)
	captchaRoute := routes.NewCaptchaRoute(captchaController, ginRouter)
	captchaRoute.Setup()

	//redirection
	redirectionService := service.NewRedirectionService(store, cache, &utilityService)
	redirectionController := controller.NewRedirectionController(redirectionService)
	redirectionRoute := routes.NewRedirectionRoute(redirectionController, ginRouter)
	redirectionRoute.Setup()

	// stat
	statService := service.NewStatService(cache, store)
	statController := controller.NewStatController(statService)
	statRoute := routes.NewStatRoute(statController, ginRouter)
	statRoute.Setup()

	// link-preview
	linkPreviewService := service.NewLinkPreviewService(cache, redirectionService, secretConfig)
	linkPreviewController := controller.NewLinkPreviewController(linkPreviewService)
	linkPreviewRoute := routes.NewLinkPreviewRoute(linkPreviewController, ginRouter)
	linkPreviewRoute.Setup()

	go startCronJobs(cronScheduler, cache, utilityService)
	go onStartupTasks()

	//server
	serverAddress := globalConfig.ServerAddress
	err = ginRouter.Gin.Run(serverAddress)
	if err != nil {
		log.Println(err)
		log.Fatal("could not start APIs")
	}
}

func initialTests() {
	for _, fontPath := range util.CaptchaFontPathList {
		if _, err := os.Stat(fontPath); err != nil {
			panic(err.Error())
		}
	}

}

func onStartupTasks() {

}

func startCronJobs(cronScheduler *cron.Cron, cache *infrastructure.Cache, utilityService service.UtilityService) {
	cronJobCtx := context.Background()

	captchaCronJob := cronjob.NewCaptchaCronJob(cache, utilityService)
	_, errCron := cronScheduler.AddFunc(util.CronSpecEveryOneMin, func() {
		captchaCronJob.Run(cronJobCtx)
	})

	if errCron != nil {
		log.Println(errCron)
	}
}
