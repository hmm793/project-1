package formatter

import "content-service-v3/services/content-service/domain/entity"

func FormatUpdateStatusBannerResponse(banner entity.BannerEntity) UpdateStatusBannerResponseFormatter {
	updateStatusBannerFormatter := UpdateStatusBannerResponseFormatter{}
	updateStatusBannerFormatter.ID = banner.ID
	updateStatusBannerFormatter.Status = banner.Status
	updateStatusBannerFormatter.UpdatedById = banner.UpdatedById
	updateStatusBannerFormatter.UpdatedByName = banner.UpdatedByName
	return updateStatusBannerFormatter
}