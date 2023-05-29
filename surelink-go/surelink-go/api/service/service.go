package service

type ServiceSilo struct {
	utilityService     *UtilityService
	redirectionService *RedirectionService
	linkPreviewService *LinkPreviewService
}

func NewServiceDiscovery() ServiceSilo {
	return ServiceSilo{}
}

func (s *ServiceSilo) UtilityService() *UtilityService {
	return s.utilityService
}

func (s *ServiceSilo) SetUtilityService(utilityService *UtilityService) {
	s.utilityService = utilityService
}

func (s *ServiceSilo) RedirectionService() *RedirectionService {
	return s.redirectionService
}

func (s *ServiceSilo) SetRedirectionService(redirectionService *RedirectionService) {
	s.redirectionService = redirectionService
}

func (s *ServiceSilo) LinkPreviewService() *LinkPreviewService {
	return s.linkPreviewService
}

func (s *ServiceSilo) SetLinkPreviewService(linkPreviewService *LinkPreviewService) {
	s.linkPreviewService = linkPreviewService
}
