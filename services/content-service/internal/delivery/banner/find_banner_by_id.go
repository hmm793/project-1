package delivery

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (h *bannerHandler) GetBanner(c *gin.Context) {

	idBanner, err := strconv.Atoi(c.Query("id"))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid Input"})
		return
	}

	banner, err := h.bannerService.FindBannerById(idBanner)

	c.JSON(http.StatusOK, banner)
}