package usecase

import (
	"content-service-v3/services/content-service/internal/usecase/banner/formatter"
	"errors"
)

func (s *serviceBanner) FindActiveBannerByCategoryBanner(name string) ([]formatter.FindBannerByCategoryBannerResponseFormatter, error) {
	banner, err := s.repositoryBannerCategory.FindBannerCategoryByName(name)

	if err != nil {
		return nil, err
	}

	idBannerCategory := banner.ID
	if idBannerCategory == 0 {
		return nil, errors.New("Banner Category Not Found")
	}

	// Teruskan ke repository
	banners,err := s.repository.FindActiveBannerByCategoryBannerId(idBannerCategory)
	
	// Formatter 
	formatterBanners := formatter.FormatFindBannerByCategoryBannerResponse(banners)
	
	if err != nil {
		return formatterBanners, err
	}

	if len(formatterBanners) == 0 {
		return formatterBanners, errors.New("Banners Not Found")
	}

	return formatterBanners,nil
}