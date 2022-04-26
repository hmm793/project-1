package usecase

import (
	"content-service-v3/services/content-service/domain/entity"
	"errors"
)

func (s *serviceBanner) FindBannerByCategoryBanner(name string) ([]entity.BannerEntity, error) {
	banner, err := s.repositoryBannerCategory.FindBannerCategoryByName(name)

	if err != nil {
		return nil, err
	}

	idBannerCategory := banner.ID
	if idBannerCategory == 0 {
		return nil, errors.New("Banner Category Not Found")
	}

	// Teruskan ke repository
	banners,err := s.repository.FindBannerByCategoryBannerId(idBannerCategory)

	if err != nil {
		return banners, err
	}

	return banners,nil
}