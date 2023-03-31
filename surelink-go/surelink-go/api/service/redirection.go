package service

import "surelink-go/infrastructure"

type RedirectionService struct {
	store *infrastructure.Store
	cache *infrastructure.Cache
}

func NewRedirectionService(store *infrastructure.Store, cache *infrastructure.Cache) RedirectionService {
	return RedirectionService{
		store: store,
		cache: cache,
	}
}
