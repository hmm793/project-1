package delivery

import "content-service-v3/services/content-service/domain/usecase"

type bannerHandler struct {
	bannerService usecase.ServiceBanner
}

func NewBannerHandler(bannerService usecase.ServiceBanner) *bannerHandler {
	return &bannerHandler{bannerService}
}
