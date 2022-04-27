package formatter

import "content-service-v3/services/content-service/domain/entity"

func FormatUpdateBannerResponse(banner entity.BannerEntity) UpdateBannerResponseFormatter {
	updateBannerFormatted := UpdateBannerResponseFormatter{}
	updateBannerFormatted.CreatedByName = banner.CreatedByName
	updateBannerFormatted.CreatedById = banner.CreatedById
	updateBannerFormatted.CreatedAt = banner.CreatedAt
	updateBannerFormatted.Status = banner.Status
	updateBannerFormatted.Link = banner.Link
	updateBannerFormatted.FileName = banner.FileName
	updateBannerFormatted.Order = banner.Order
	updateBannerFormatted.UpdatedByName = banner.UpdatedByName
	updateBannerFormatted.UpdatedById = banner.UpdatedById
	updateBannerFormatted.UpdatedAt = banner.UpdatedAt
	updateBannerFormatted.BannerCategoryID = banner.BannerCategoryID
	return updateBannerFormatted
}