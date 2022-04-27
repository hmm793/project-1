package usecase

import (
	"content-service-v3/services/content-service/domain/dto"
	"content-service-v3/services/content-service/domain/entity"
	"errors"
	"time"
)

func (s *serviceBanner) DeleteBanner(id int, input dto.DeleteBannerInput) (int64, error) {
	// Mapper From DTO to Entity
	mappedBanner := entity.Delete_BannerDTO_To_BannerEntity(input)

	// Cek Apakah Ada Banner Dengan ID = id
	currBanner, err := s.repository.FindBannerById(id)

	if err != nil {
		return 0, err
	}

	if currBanner.ID == 0 {
		return 0, errors.New("Banner Not Found")
	}

	currBanner.ID = mappedBanner.ID
	currBanner.DeletedById = mappedBanner.DeletedById
	currBanner.DeletedByName = mappedBanner.DeletedByName
	currBanner.DeletedAt = time.Now()

	// Delete Banner
	rowAffected, err := s.repository.DeleteBannerById(currBanner)

	if err != nil {
		return rowAffected,err
	}
	
	return rowAffected, nil
}