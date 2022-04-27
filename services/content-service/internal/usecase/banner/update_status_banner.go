package usecase

import (
	"content-service-v3/services/content-service/domain/dto"
	"content-service-v3/services/content-service/domain/entity"
	"content-service-v3/services/content-service/internal/usecase/banner/formatter"
	"errors"
)

func (s *serviceBanner) UpdateStatusBanner(id int, input dto.UpdateStatusBannerInput) (formatter.UpdateStatusBannerResponseFormatter, error) {
	var updatedBanner formatter.UpdateStatusBannerResponseFormatter

	// Mapper From DTO to Entity
	mappedBanner := entity.Update_Status_Banner(input)

	// Cek Apakah Ada Banner Dengan ID = id
	currBanner, err := s.repository.FindBannerById(id)

	if err != nil {
		return updatedBanner, err
	}

	if currBanner.ID == 0 {
		return updatedBanner, errors.New("Banner Not Found")
	}

	currBanner.ID = mappedBanner.ID
	currBanner.Status = mappedBanner.Status
	currBanner.UpdatedById = mappedBanner.UpdatedById
	currBanner.UpdatedByName = mappedBanner.UpdatedByName

	// Update Status Banner
	banner, err := s.repository.UpdateBanner(currBanner)

	updatedBanner = formatter.FormatUpdateStatusBannerResponse(banner)
	
	if err != nil {
		return updatedBanner, err
	}

	return updatedBanner, nil
}