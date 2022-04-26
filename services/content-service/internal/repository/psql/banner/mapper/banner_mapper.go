package mapper

import (
	"content-service-v3/services/content-service/domain/entity"
	"content-service-v3/services/content-service/internal/repository/psql/banner/model_banner"
)

func To_Model(bannerEntity entity.BannerEntity) model_banner.BannerModel {
	mappedBanner := model_banner.BannerModel{
		ID:            		bannerEntity.ID,
		Link:          		bannerEntity.Link,
		FileName:      		bannerEntity.FileName,
		CreatedById:   		bannerEntity.CreatedById,
		CreatedByName: 		bannerEntity.CreatedByName,
		CreatedAt:     		bannerEntity.CreatedAt,
		UpdatedAt:     		bannerEntity.UpdatedAt,
		UpdatedById:   		bannerEntity.UpdatedById,
		UpdatedByName: 		bannerEntity.UpdatedByName,
		DeletedByName: 		bannerEntity.DeletedByName,
		DeletedById:   		bannerEntity.DeletedById,
		Order:         		bannerEntity.Order,
		DeletedAt:     		bannerEntity.DeletedAt,
		Status: 	   		bannerEntity.Status,
		BannerCategoryID: 	bannerEntity.BannerCategoryID,
	}
	return mappedBanner
}

func To_Entity(bannerModel model_banner.BannerModel) entity.BannerEntity {
	mappedBanner := entity.BannerEntity{
		ID:            		bannerModel.ID,
		Link:          		bannerModel.Link,
		CreatedById:   		bannerModel.CreatedById,
		CreatedByName: 		bannerModel.CreatedByName,
		CreatedAt:     		bannerModel.CreatedAt,
		DeletedByName: 		bannerModel.DeletedByName,
		DeletedById:   		bannerModel.DeletedById,
		Order:         		bannerModel.Order,
		DeletedAt:     		bannerModel.DeletedAt,
		FileName:      		bannerModel.FileName,
		UpdatedAt:     		bannerModel.UpdatedAt,
		UpdatedById:   		bannerModel.UpdatedById,
		UpdatedByName: 		bannerModel.UpdatedByName,
		Status: 		  	bannerModel.Status,
		BannerCategoryID: 	bannerModel.BannerCategoryID,
	}
	return mappedBanner
}
