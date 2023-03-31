package main

import (
	_ "github.com/lib/pq"
	"log"
	"os"
	"surelink-go/api/controller"
	"surelink-go/api/routes"
	"surelink-go/api/service"
	_ "surelink-go/api/service"
	"surelink-go/infrastructure"
	"surelink-go/util"
)

func main() {
	initialTests()

	globalConfig, err := util.LoadGlobalConfig(".")
	if err != nil {
		log.Fatal("can not load global config", err)
	}

	//
	//conn, err := sql.Open(globalConfig.DBDriver, globalConfig.DBSource)
	//if err != nil {
	//	log.Fatal("can't connect to the database", err)
	//}
	//store := db.NewStore(conn)
	//
	//redisStore := gedis.NewRedisStore(globalConfig.RedisUrl)
	//
	//serverObj := goserver.NewServer(store, redisStore)
	//err = serverObj.Start(globalConfig.ServerAddress)
	//if err != nil {
	//	log.Fatal("can't start the server", err)
	//}

	//miscellaneous
	//random := rand.New(rand.NewSource(time.Now().UnixNano()))

	//database and cache
	cache := infrastructure.NewCache(globalConfig.RedisUrl)

	// initialize gin router
	log.Println("Initializing Routes")
	ginRouter := infrastructure.NewGinRouter()

	//initialize service
	//utilityService := service.NewUtilityService(cache, random)

	// captcha
	captchaService := service.NewCaptchaService(cache)
	captchaController := controller.NewCaptchaController(captchaService)
	captchaRoute := routes.NewCaptchaRoute(captchaController, ginRouter)
	captchaRoute.Setup()

	serverAddress := "0.0.0.0:9000"
	err = ginRouter.Gin.Run(serverAddress)
	if err != nil {
		log.Println(err)
		log.Fatal("could not start APIs")
	}

}

func initialTests() {
	if _, err := os.Stat(util.FONT_COMIC_PATH); err != nil {
		panic(err.Error())
	}
}
