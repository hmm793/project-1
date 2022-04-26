package dto

type CreateBannerInput struct {
	Link             string `form:"link" binding:"required"`
	Order            int    `form:"order" binding:"required"`
	Status           int    `form:"status" binding:"required"`
	CreatedById      int    `form:"createdById" binding:"required"`
	CreatedByName    string `form:"createdByName" binding:"required"`
	BannerCategoryID int    `form:"bannerCategoryId" binding:"required"`
}

type IDBannerQuery struct {
	ID int `uri:"id"`
}