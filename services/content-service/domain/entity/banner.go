package entity

import (
	"content-service-v3/services/content-service/domain/dto"
	"time"
)

type BannerEntity struct {
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
}

// Convert DTO ke Entity
func Create_BannerDTO_To_BannerEntity(input dto.CreateBannerInput) BannerEntity {
	mappedBanner := BannerEntity{
		Link: input.Link,
		Order: input.Order,
		Status: input.Status,
		CreatedById: input.CreatedById,
		CreatedByName: input.CreatedByName,
		BannerCategoryID: input.BannerCategoryID,
	}
	return mappedBanner
}