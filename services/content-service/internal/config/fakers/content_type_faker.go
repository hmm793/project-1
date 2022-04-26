package fakers

import (
	"content-service-v3/services/content-service/internal/repository/psql/banner_category/model_banner_category"

	"gorm.io/gorm"
)

func BannerCategoryFaker(db *gorm.DB) []model_banner_category.BannerCategoryModel {
	data := []model_banner_category.BannerCategoryModel{
		{ID: 1, Name: "Category 1"},
		{ID: 2, Name: "Category 2"},
		{ID: 3, Name: "Category 3"},
	}
	return data
}