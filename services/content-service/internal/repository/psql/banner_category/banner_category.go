package banner_category

import (
	"content-service-v3/services/content-service/domain/entity"
	"content-service-v3/services/content-service/internal/repository/psql/banner_category/mapper"
	"content-service-v3/services/content-service/internal/repository/psql/banner_category/model_banner_category"

	"gorm.io/gorm"
)

type repositoryBannerCategory struct {
	db *gorm.DB
}

func NewRepositoryBannerCategory(db *gorm.DB) *repositoryBannerCategory {
	return &repositoryBannerCategory{db}
}


func (r *repositoryBannerCategory) FindBannerCategoryByName(name string) (entity.BannerCategoryEntity, error) {
	var bannerCategoryInModel model_banner_category.BannerCategoryModel

	err := r.db.Where("name = ?", name).Find(&bannerCategoryInModel).Error
	
	// Mapper To Entity
	mappedBannerCategoryToEntity := mapper.To_Entity(bannerCategoryInModel)

	// Error
	if err != nil {
		return mappedBannerCategoryToEntity, err
	}
	// Not error
	return mappedBannerCategoryToEntity, nil
}