package formatter

import "content-service-v3/services/content-service/domain/entity"

func FormatFindBannerByIDResponse(banner entity.BannerEntity) FindBannerByIDResponseFormatter {
	findBannerByIDFormatter := FindBannerByIDResponseFormatter{}
	findBannerByIDFormatter.ID = banner.ID
	findBannerByIDFormatter.CreatedByName = banner.CreatedByName
	findBannerByIDFormatter.CreatedById = banner.CreatedById
	findBannerByIDFormatter.CreatedAt = banner.CreatedAt
	findBannerByIDFormatter.Status = banner.Status
	findBannerByIDFormatter.Link = banner.Link
	findBannerByIDFormatter.FileName = banner.FileName
	findBannerByIDFormatter.Order = banner.Order
	findBannerByIDFormatter.UpdatedByName = banner.UpdatedByName
	findBannerByIDFormatter.UpdatedById = banner.UpdatedById
	findBannerByIDFormatter.UpdatedAt = banner.UpdatedAt
	findBannerByIDFormatter.BannerCategoryID = banner.BannerCategoryID
	return findBannerByIDFormatter
}