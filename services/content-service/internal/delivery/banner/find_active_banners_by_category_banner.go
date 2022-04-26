package delivery

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *bannerHandler) GetActiveBannerByCategoryBanner(c *gin.Context) {
	categoryName := c.Param("categoryBanner")
	
	banners, err := h.bannerService.FindActiveBannerByCategoryBanner(categoryName)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, banners)
}