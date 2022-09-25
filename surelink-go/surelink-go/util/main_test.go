package util

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http/httptest"
	"os"
	gedis "surelink-go/redisStore"
	"testing"
)

var redisStore *gedis.RedisStore
var ctx *gin.Context

func TestMain(t *testing.M) {
	globalConfig, err := LoadGlobalConfig("../")
	if err != nil {
		log.Fatal("can not load global config", err)
	}

	redisStore = gedis.NewRedisStore(globalConfig.RedisUrl)

	gin.SetMode(gin.TestMode)
	ctx, _ = gin.CreateTestContext(httptest.NewRecorder())

	os.Exit(t.Run())
}
