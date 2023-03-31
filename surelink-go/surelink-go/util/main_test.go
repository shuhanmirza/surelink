package util

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http/httptest"
	"os"
	"testing"
)

var ctx *gin.Context

func TestMain(t *testing.M) {
	_, err := LoadGlobalConfig("../")
	if err != nil {
		log.Fatal("can not load global config", err)
	}

	gin.SetMode(gin.TestMode)
	ctx, _ = gin.CreateTestContext(httptest.NewRecorder())

	os.Exit(t.Run())
}
