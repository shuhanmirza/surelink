package service

import (
	"github.com/gin-gonic/gin"
	"surelink-go/api/structs"
	"surelink-go/infrastructure"
)

type LinkPreviewService struct {
	cache *infrastructure.Cache
}

func NewLinkPreviewService(cache *infrastructure.Cache) LinkPreviewService {
	return LinkPreviewService{
		cache: cache,
	}
}

func (s *LinkPreviewService) GetLinkPreview(ctx *gin.Context, request structs.GetLinkPreviewRequest) (response structs.GetLinkPreviewResponse, err error) {

	return response, nil
}
