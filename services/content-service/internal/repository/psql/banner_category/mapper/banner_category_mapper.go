package mapper

import (
	"content-service-v3/services/content-service/domain/entity"
	"content-service-v3/services/content-service/internal/repository/psql/banner_category/model_banner_category"
)

func To_Entity(bannerCategoryModel model_banner_category.BannerCategoryModel) entity.BannerCategoryEntity {
	mappedBannerCategory := entity.BannerCategoryEntity{
		ID: bannerCategoryModel.ID,
		Name: bannerCategoryModel.Name,
	}
	return mappedBannerCategory
}
