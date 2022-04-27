package usecase

import (
	"content-service-v3/services/content-service/domain/dto"
	"content-service-v3/services/content-service/domain/entity"
	"errors"
	"strconv"
)

func (s *serviceBanner) UpdateOrderBanner(input dto.UpdateOrderBannerInput) ([]entity.OrderEntity, error) {
	var updatedBannersByOrder []entity.OrderEntity
	
	// Mapper From DTO to Entity
	mappedOder := entity.Update_Order_BannerDTO_TO_BannerEntity(input)

	// Looping Through mappedOrder.OrderData
	for _, order := range mappedOder.OrderData {
		banner, err := s.repository.FindBannerById(order.ID)
		if err != nil {
			return updatedBannersByOrder, err
		}

		if banner.ID == 0 {
			return updatedBannersByOrder, errors.New("Banner With ID:"+strconv.Itoa(order.ID)+" Not Found")
		}

		banner.ID = order.ID
		banner.Status = order.Status
		banner.Order = order.Order
		banner.UpdatedById = mappedOder.UpdatedById
		banner.UpdatedByName = mappedOder.UpdatedByName

		// Update Status Banner
		updatedBanner, err := s.repository.UpdateBanner(banner)

		if err != nil {
			return updatedBannersByOrder, err
		}
		
		// Setelah Di Update Maka Di Map ke bantuk []entity.OrderEntity
		updatedBannerByOrder := entity.OrderEntity{
			ID: updatedBanner.ID,
			Order: updatedBanner.Order,
			Status: updatedBanner.Status,
		}
		updatedBannersByOrder = append(updatedBannersByOrder, updatedBannerByOrder)
	}

	return updatedBannersByOrder, nil
}