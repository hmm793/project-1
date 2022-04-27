package delivery

import (
	"content-service-v3/services/content-service/domain/dto"
	"content-service-v3/services/content-service/internal/delivery/helper"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *bannerHandler) UpdateStatusBanner(c *gin.Context) {
	var input dto.UpdateStatusBannerInput
	err := c.ShouldBindJSON(&input)

	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{
			"message": errors,
		}
		c.JSON(http.StatusBadRequest, errorMessage)
		return
	}

	updatedBanner, err := h.bannerService.UpdateStatusBanner(input.ID, input)
	
	if err != nil {
		errorMessage := gin.H{
			"message": err.Error(),
		}
		c.JSON(http.StatusBadRequest, errorMessage)
		return
	}
	
	c.JSON(http.StatusOK, updatedBanner)
}