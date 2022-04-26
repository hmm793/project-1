package usecase

import (
	"content-service-v3/services/content-service/internal/usecase/banner/formatter"
	"errors"
	"fmt"
)

func (s *serviceBanner) FindBannerById(id int) (formatter.FindBannerByIDResponseFormatter, error) {

	fmt.Println("ID", id)
	banner, err := s.repository.FindBannerById(id)
	fmt.Println("Banner", banner)
	
	formattedBanner := formatter.FormatFindBannerByIDResponse(banner)
	fmt.Println("formattedBanner", formattedBanner.ID)
	if err != nil {
		return formattedBanner, err
	}
	if formattedBanner.ID == 0 {
		return formattedBanner, errors.New("Banner Not Found")
	}
	return formattedBanner, nil
}