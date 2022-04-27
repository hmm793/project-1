package usecase

import (
	"content-service-v3/services/content-service/domain/dto"
	"content-service-v3/services/content-service/domain/entity"
	"content-service-v3/services/content-service/internal/usecase/banner/formatter"
	"mime/multipart"
)

type ServiceBanner interface {
	CreateBanner(input dto.CreateBannerInput, file *multipart.FileHeader) (formatter.CreateBannerResponseFormatter, string, error)
	FindBannerById(id int) (formatter.FindBannerByIDResponseFormatter, error)
	FindBannerByCategoryBanner(name string) ([]formatter.FindBannerByCategoryBannerResponseFormatter, error)
	FindActiveBannerByCategoryBanner(name string) ([]formatter.FindBannerByCategoryBannerResponseFormatter, error)
	UpdateBanner(id int, input dto.UpdateBannerInput, file *multipart.FileHeader) (entity.BannerEntity, string, error)
	DeleteBanner(id int, input dto.DeleteBannerInput) (int64, error)
	UpdateStatusBanner(id int, input dto.UpdateStatusBannerInput) (entity.BannerEntity, error)
}