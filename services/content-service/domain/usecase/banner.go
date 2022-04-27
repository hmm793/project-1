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
	FindActiveBannerByCategoryBanner(name string) ([]formatter.FindBannerByCategoryBannerResponseFormatter, error)
	DeleteBanner(id int, input dto.DeleteBannerInput) (int64, error)
	
	UpdateBanner(id int, input dto.UpdateBannerInput, file *multipart.FileHeader) (formatter.UpdateBannerResponseFormatter, string, error)
	UpdateStatusBanner(id int, input dto.UpdateStatusBannerInput) (formatter.UpdateStatusBannerResponseFormatter, error)
	UpdateOrderBanner(input dto.UpdateOrderBannerInput) ([]formatter.UpdateOrderBannerResponseFormatter, error)
}