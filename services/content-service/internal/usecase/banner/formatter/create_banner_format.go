package formatter

import "time"

type CreateBannerResponseFormatter struct {
	CreatedByName string    `json:"createdByName"`
	CreatedById   int       `json:"createdById"`
	CreatedAt     time.Time `json:"createdAt"`
	Status        int       `json:"status"`
	Link          string    `json:"link"`
	FileName      string    `json:"filename"`
	Order         int       `json:"order"`
	BannerCategoryID int `json:"bannerCategoryID"`
}