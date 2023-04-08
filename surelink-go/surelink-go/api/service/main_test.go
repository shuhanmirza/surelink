package service

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	"log"
	"math/rand"
	"net/http/httptest"
	"os"
	"surelink-go/infrastructure"
	"surelink-go/util"
	"testing"
	"time"
)

var ctx *gin.Context
var utilityService UtilityService
var statService StatService

func TestMain(t *testing.M) {
	globalConfig, err := util.LoadGlobalConfig("../../")
	if err != nil {
		log.Fatal("can not load global config", err)
	}

	random := rand.New(rand.NewSource(time.Now().UnixNano()))
	cache := infrastructure.NewCache(globalConfig.RedisUrl)
	//database and cache
	conn, err := sql.Open(globalConfig.DBDriver, globalConfig.DBSource)
	if err != nil {
		log.Fatal("can't connect to the database", err)
	}
	store := infrastructure.NewStore(conn)

	utilityService = NewUtilityService(cache, random)
	statService = NewStatService(cache, store)

	gin.SetMode(gin.TestMode)
	ctx, _ = gin.CreateTestContext(httptest.NewRecorder())

	os.Exit(t.Run())
}
