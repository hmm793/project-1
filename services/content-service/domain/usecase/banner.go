package usecase

import (
	"content-service-v3/services/content-service/domain/dto"
	"content-service-v3/services/content-service/internal/usecase/banner/formatter"
	"mime/multipart"
)

type ServiceBanner interface {
	CreateBanner(input dto.CreateBannerInput, file *multipart.FileHeader) (formatter.CreateBannerResponseFormatter, string, error)
	FindBannerById(id int) (formatter.FindBannerByIDResponseFormatter, error)
	FindBannerByCategoryBanner(name string) ([]formatter.FindBannerByCategoryBannerResponseFormatter, error)
}