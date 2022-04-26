package repository

import (
	"content-service-v3/services/content-service/domain/entity"
)

type RepositoryBanner interface {
	SaveBanner(banner entity.BannerEntity) (entity.BannerEntity, error)
	FindBannerById(id int) (entity.BannerEntity, error)
	FindBannerByCategoryBannerId(idBannerCategory int) ([]entity.BannerEntity, error)
	FindActiveBannerByCategoryBannerId(idBannerCategory int) ([]entity.BannerEntity, error)
}
