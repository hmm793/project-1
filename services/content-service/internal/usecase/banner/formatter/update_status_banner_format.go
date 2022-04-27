package formatter

type UpdateStatusBannerResponseFormatter struct {
	ID            int    `json:"id"`
	Status        int    `json:"status"`
	UpdatedById   int    `json:"updatedById"`
	UpdatedByName string `json:"updatedByName"`
}