package service

import "surelink-go/infrastructure"

type RedirectionService struct {
	store        *infrastructure.Store
	cacheService CacheService
}

func NewRedirectionService(cacheService CacheService, store *infrastructure.Store) RedirectionService {
	return RedirectionService{
		store:        store,
		cacheService: cacheService,
	}
}
