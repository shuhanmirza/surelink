package service

import (
	"log"
	"testing"
)

func TestStatService_GetHomePageStats(t *testing.T) {
	response, err := statService.GetHomePageStats(ctx)

	log.Print(response)
	log.Println(err)
}
