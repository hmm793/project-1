package formatter

type UpdateOrderBannerResponseFormatter struct {
	ID   		  int       `json:"id"`
	Order         int       `json:"order"`
	Status        int       `json:"status"`
}