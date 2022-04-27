package dto

type CreateBannerInput struct {
	Link             string `form:"link" binding:"required"`
	Order            int    `form:"order" binding:"required"`
	Status           int    `form:"status" binding:"required"`
	CreatedById      int    `form:"createdById" binding:"number,required"`
	CreatedByName    string `form:"createdByName" binding:"required"`
	BannerCategoryID int    `form:"bannerCategoryId" binding:"required"`
}

type UpdateBannerInput struct {
	Link             string `form:"link" binding:"required"`
	Order            int    `form:"order" binding:"required"`
	Status           int    `form:"status" binding:"required"`
	UpdatedById      int    `form:"updatedById" binding:"number,required"`
	UpdatedByName    string `form:"updatedByName" binding:"required"`
	BannerCategoryID int    `form:"bannerCategoryId" binding:"required"`
}
type IDInput struct {
	ID int `uri:"id" binding:"required"`
}

type DeleteBannerInput struct {
	ID            int    `form:"id" binding:"required"`
	DeletedById   int    `form:"deletedById" binding:"required"`
	DeletedByName string `form:"deletedByName" binding:"required"`
}

type UpdateStatusBannerInput struct {
	ID            int    `json:"id" binding:"required"`
	Status        int    `json:"status" binding:"required"`
	UpdatedById   int    `json:"updatedById" binding:"number,required"`
	UpdatedByName string `json:"updatedByName" binding:"required"`
}

type OrderStruct struct {
	ID     int `json:"id"`
	Order  int `json:"order"`
	Status int `json:"status"`
}

type UpdateOrderBannerInput struct {
	UpdatedById   int           `json:"updatedById" binding:"number,required"`
	UpdatedByName string        `json:"updatedByName" binding:"required"`
	OrderData     []OrderStruct `json:"orderData"`
}
