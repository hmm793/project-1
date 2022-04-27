package usecase

import (
	"content-service-v3/services/content-service/domain/dto"
	"content-service-v3/services/content-service/domain/entity"
	"errors"
)

func (s *serviceBanner) UpdateStatusBanner(id int, input dto.UpdateStatusBannerInput) (entity.BannerEntity, error) {
	// Mapper From DTO to Entity
	mappedBanner := entity.Update_Status_BannerDTO_TO_BannerEntity(input)

	// Cek Apakah Ada Banner Dengan ID = id
	currBanner, err := s.repository.FindBannerById(id)

	if err != nil {
		return currBanner, err
	}

	if currBanner.ID == 0 {
		return currBanner, errors.New("Banner Not Found")
	}

	currBanner.ID = mappedBanner.ID
	currBanner.Status = mappedBanner.Status
	currBanner.UpdatedById = mappedBanner.UpdatedById
	currBanner.UpdatedByName = mappedBanner.UpdatedByName

	// Update Status Banner
	banner, err := s.repository.UpdateBanner(currBanner)

	if err != nil {
		return banner, err
	}

	return banner, nil
}