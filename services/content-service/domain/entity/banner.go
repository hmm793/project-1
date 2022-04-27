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
func Update_BannerDTO_To_BannerEntity(input dto.UpdateBannerInput) BannerEntity {
	mappedBanner := BannerEntity{
		Link: input.Link,
		Order: input.Order,
		Status: input.Status,
		UpdatedById: input.UpdatedById,
		UpdatedByName: input.UpdatedByName,
		BannerCategoryID: input.BannerCategoryID,
	}
	return mappedBanner
}

func Delete_BannerDTO_To_BannerEntity(input dto.DeleteBannerInput) BannerEntity {
	mappedBanner := BannerEntity{
		ID: input.ID,
		DeletedById: input.DeletedById,
		DeletedByName: input.DeletedByName,
	}
	return mappedBanner
}

func Update_Status_BannerDTO_TO_BannerEntity(input dto.UpdateStatusBannerInput) BannerEntity {
	mappedBanner := BannerEntity{
		ID: input.ID,
		Status: input.Status,
		UpdatedById: input.UpdatedById,
		UpdatedByName: input.UpdatedByName,
	}
	return mappedBanner
}


type OrderEntity struct {
	ID     int `json:"id"`
	Order  int `json:"order"`
	Status int `json:"status"`
}

type UpdateOrderBannerEntity struct {
	UpdatedById   int           `json:"updatedById" binding:"number,required"`
	UpdatedByName string        `json:"updatedByName" binding:"required"`
	OrderData     []OrderEntity `json:"orderData"`
}

func Update_Order_BannerDTO_TO_BannerEntity(input dto.UpdateOrderBannerInput) UpdateOrderBannerEntity {
	var orders []OrderEntity
	for _, order := range input.OrderData {
		var newOrder OrderEntity
		newOrder.ID = order.ID
		newOrder.Order = order.Order
		newOrder.Status = order.Status 
		orders = append(orders, newOrder)
	}

	mappedOrder := UpdateOrderBannerEntity{
		UpdatedById: input.UpdatedById,
		UpdatedByName: input.UpdatedByName,
		OrderData: orders,
	}
	return mappedOrder
}