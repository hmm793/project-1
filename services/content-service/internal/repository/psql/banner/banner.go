package banner

import (
	"content-service-v3/services/content-service/domain/entity"
	"content-service-v3/services/content-service/internal/repository/psql/banner/mapper"
	"content-service-v3/services/content-service/internal/repository/psql/banner/model_banner"
	"time"

	"gorm.io/gorm"
)

type repositoryBanner struct {
	db *gorm.DB
}

func NewRepositoryBanner(db *gorm.DB) *repositoryBanner {
	return &repositoryBanner{db}
}

func (r *repositoryBanner) SaveBanner(banner entity.BannerEntity) (entity.BannerEntity, error) {
	// Mapper To Model
	mappedBannerToModel := mapper.To_Model(banner)

	// Save To Database
	err := r.db.Create(&mappedBannerToModel).Error

	// Mapper To Entity
	mappedBannerToEntity := mapper.To_Entity(mappedBannerToModel)

	// Error
	if err != nil {
		return mappedBannerToEntity, err
	}
	// Not error
	return mappedBannerToEntity, nil
}


func (r *repositoryBanner) FindBannerById(id int) (entity.BannerEntity, error) {
	var bannerInModel model_banner.BannerModel
	var t time.Time
	err := r.db.Where("id = ?", id).Where("deleted_at = ?", t).Preload("BannerCategory").Find(&bannerInModel).Error
	
	// Mapper To Entity
	mappedBannerToEntity := mapper.To_Entity(bannerInModel)

	// Error
	if err != nil {
		return mappedBannerToEntity, err
	}

	// Not error
	return mappedBannerToEntity, nil
}

func (r *repositoryBanner) FindBannerByCategoryBannerId(idBannerCategory int) ([]entity.BannerEntity, error) {
	var bannersInModel []model_banner.BannerModel
	var t time.Time
	err := r.db.Where("banner_category_id = ?", idBannerCategory).Where("deleted_at = ?", t).Find(&bannersInModel).Error
	
	// Mapper To Entity
	var mappedBannersToEntity []entity.BannerEntity

	for _, banner := range bannersInModel {
		mappedBannersToEntity = append(mappedBannersToEntity, mapper.To_Entity(banner))
	}

	// Error
	if err != nil {
		return mappedBannersToEntity, err
	}
	// Not error
	return mappedBannersToEntity, nil
}
func (r *repositoryBanner) FindActiveBannerByCategoryBannerId(idBannerCategory int) ([]entity.BannerEntity, error) {
	var bannersInModel []model_banner.BannerModel
	var t time.Time
	err := r.db.Where("banner_category_id = ?", idBannerCategory).Where("deleted_at = ?", t).Where("status = ?", 1).Find(&bannersInModel).Error
	
	// Mapper To Entity
	var mappedBannersToEntity []entity.BannerEntity

	for _, banner := range bannersInModel {
		mappedBannersToEntity = append(mappedBannersToEntity, mapper.To_Entity(banner))
	}

	// Error
	if err != nil {
		return mappedBannersToEntity, err
	}
	// Not error
	return mappedBannersToEntity, nil
}

func (r *repositoryBanner) UpdateBanner(banner entity.BannerEntity) (entity.BannerEntity, error) {
	// Ubah Ke Model
	mappedBanner := mapper.To_Model(banner)
	
	// Save
	err := r.db.Save(&mappedBanner).Error
	
	// Ubah Ke Entity
	mappedBannerToEntity := mapper.To_Entity(mappedBanner)
	
	// Klo Error
	if err != nil {
		return mappedBannerToEntity, err
	}
	// Klo Tidak Error
	return mappedBannerToEntity, nil
}

func (r *repositoryBanner) DeleteBannerById(banner entity.BannerEntity) (int64, error) {
	// Ubah Ke Model
	mappedBanner := mapper.To_Model(banner)
	
	result := r.db.Save(&mappedBanner)
	err := result.Error
	rowsAffected := result.RowsAffected
	if err != nil {
		return  rowsAffected, err
	}
	return rowsAffected, nil
}