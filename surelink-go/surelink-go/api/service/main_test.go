package service

import (
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

func TestMain(t *testing.M) {
	globalConfig, err := util.LoadGlobalConfig("../../")
	if err != nil {
		log.Fatal("can not load global config", err)
	}

	random := rand.New(rand.NewSource(time.Now().UnixNano()))
	cache := infrastructure.NewCache(globalConfig.RedisUrl)
	utilityService = NewUtilityService(cache, random)

	gin.SetMode(gin.TestMode)
	ctx, _ = gin.CreateTestContext(httptest.NewRecorder())

	os.Exit(t.Run())
}
