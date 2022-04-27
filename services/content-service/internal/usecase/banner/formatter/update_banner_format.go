package formatter

import "time"

type UpdateBannerResponseFormatter struct {
	CreatedByName string    `json:"createdByName"`
	CreatedById   int       `json:"createdById"`
	CreatedAt     time.Time `json:"createdAt"`
	Status        int       `json:"status"`
	Link          string    `json:"link"`
	FileName      string    `json:"filename"`
	Order         int       `json:"order"`
	UpdatedByName string    `json:"updatedByName"`
	UpdatedById   int       `json:"updatedById"`
	UpdatedAt     time.Time `json:"updatedAt"`
	BannerCategoryID int `json:"bannerCategoryID"`
}