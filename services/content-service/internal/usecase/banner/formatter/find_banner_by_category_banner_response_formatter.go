package formatter

import "content-service-v3/services/content-service/domain/entity"

func FormatFindBannerByCategoryBannerResponse(banners []entity.BannerEntity) []FindBannerByCategoryBannerResponseFormatter {
	var formattedBanners []FindBannerByCategoryBannerResponseFormatter

	for _, banner := range banners {
		formattedBanners = append(formattedBanners, convertToFindBannerByCategoryBannerResponseFormatter(banner))
	}
	return formattedBanners
}


func convertToFindBannerByCategoryBannerResponseFormatter(banner entity.BannerEntity) FindBannerByCategoryBannerResponseFormatter {
	formattedBanner := FindBannerByCategoryBannerResponseFormatter{}
	formattedBanner.CreatedByName = banner.CreatedByName
	formattedBanner.CreatedById = banner.CreatedById
	formattedBanner.CreatedAt = banner.CreatedAt
	formattedBanner.Status = banner.Status
	formattedBanner.Link = banner.Link
	formattedBanner.FileName = banner.FileName
	formattedBanner.Order = banner.Order
	formattedBanner.UpdatedByName = banner.UpdatedByName
	formattedBanner.UpdatedById = banner.UpdatedById
	formattedBanner.UpdatedAt = banner.UpdatedAt
	formattedBanner.BannerCategoryID = banner.BannerCategoryID
	return formattedBanner
}