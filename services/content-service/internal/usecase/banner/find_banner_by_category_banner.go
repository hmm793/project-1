package usecase

import (
	"content-service-v3/services/content-service/internal/usecase/banner/formatter"
	"errors"
)

func (s *serviceBanner) FindBannerByCategoryBanner(name string) ([]formatter.FindBannerByCategoryBannerResponseFormatter, error) {
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

	// Formatter 
	formatterBanners := formatter.FormatFindBannerByCategoryBannerResponse(banners)

	if err != nil {
		return formatterBanners, err
	}

	return formatterBanners,nil
}