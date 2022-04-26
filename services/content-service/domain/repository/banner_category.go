package repository

import (
	"content-service-v3/services/content-service/domain/entity"
)

type RepositoryBannerCategory interface {
	FindBannerCategoryByName(name string) (entity.BannerCategoryEntity, error)
}
