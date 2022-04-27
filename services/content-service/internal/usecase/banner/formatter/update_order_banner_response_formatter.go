package formatter

import "content-service-v3/services/content-service/domain/entity"

func FormatUpdateOrderBannerResponse(banner entity.BannerEntity) UpdateOrderBannerResponseFormatter {
	updateOrderBannerFormatter := UpdateOrderBannerResponseFormatter{}
	updateOrderBannerFormatter.ID = banner.ID
	updateOrderBannerFormatter.Order = banner.Order
	updateOrderBannerFormatter.Status = banner.Status
	return updateOrderBannerFormatter
}