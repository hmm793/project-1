package formatter

import "content-service-v3/services/content-service/domain/entity"

func FormatCreateBannerResponse(banner entity.BannerEntity) CreateBannerResponseFormatter {
	createBannerFormatter := CreateBannerResponseFormatter{}
	createBannerFormatter.CreatedAt = banner.CreatedAt
	createBannerFormatter.CreatedById = banner.CreatedById
	createBannerFormatter.CreatedByName = banner.CreatedByName
	createBannerFormatter.FileName = banner.FileName
	createBannerFormatter.Link = banner.Link
	createBannerFormatter.Order = banner.Order
	createBannerFormatter.Status = banner.Status
	createBannerFormatter.BannerCategoryID = banner.BannerCategoryID
	return createBannerFormatter
}