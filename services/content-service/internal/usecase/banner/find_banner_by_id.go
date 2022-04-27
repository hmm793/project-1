package usecase

import (
	"content-service-v3/services/content-service/internal/usecase/banner/formatter"
	"errors"
)

func (s *serviceBanner) FindBannerById(id int) (formatter.FindBannerByIDResponseFormatter, error) {
	banner, err := s.repository.FindBannerById(id)
	
	formattedBanner := formatter.FormatFindBannerByIDResponse(banner)
	if err != nil {
		return formattedBanner, err
	}
	if formattedBanner.ID == 0 {
		return formattedBanner, errors.New("Banner Not Found")
	}
	return formattedBanner, nil
}