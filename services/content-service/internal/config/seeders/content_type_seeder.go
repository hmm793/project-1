package seeders

import (
	"content-service-v3/services/content-service/internal/config/fakers"
	"content-service-v3/services/content-service/internal/repository/psql/banner_category/model_banner_category"

	"gorm.io/gorm"
)

type BannerSeeder struct {
	data []model_banner_category.BannerCategoryModel
}

func CreateBannerCategorySeeder(db *gorm.DB) BannerSeeder {
	return BannerSeeder{
		data : fakers.BannerCategoryFaker(db),
	}
}

func BannerCategoryDBSeed(db *gorm.DB) error {
	for _, datum := range CreateBannerCategorySeeder(db).data {
		err := db.Debug().Create(&datum).Error
		if err != nil {
			return err
		}
	}
	return nil
}