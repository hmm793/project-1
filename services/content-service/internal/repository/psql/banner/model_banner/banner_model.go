package model_banner

import (
	"content-service-v3/services/content-service/internal/repository/psql/banner_category/model_banner_category"
	"time"
)

type BannerModel struct {
	ID            		int
	Link          		string
	FileName 	 		string
	CreatedById   		int
	CreatedByName	 	string
	CreatedAt     		time.Time
	UpdatedAt     		time.Time
	UpdatedById   		int
	UpdatedByName 		string
	DeletedByName 		string
	DeletedById   		int
	Order 		  		int
	DeletedAt     		time.Time
	Status 		  		int
	BannerCategoryID 	int
	BannerCategory model_banner_category.BannerCategoryModel
}