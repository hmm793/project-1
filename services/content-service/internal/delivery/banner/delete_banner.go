package delivery

import (
	"content-service-v3/services/content-service/domain/dto"
	"content-service-v3/services/content-service/internal/delivery/helper"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *bannerHandler) DeleteBanner(c *gin.Context) {
	var inputID dto.IDInput
	err := c.ShouldBindUri(&inputID)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{
			"message": errors,
		}
		c.JSON(http.StatusBadRequest, errorMessage)
		return
	}

	var input dto.DeleteBannerInput
	err = c.ShouldBind(&input)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{
			"message": errors,
		}
		c.JSON(http.StatusBadRequest, errorMessage)
		return
	}
	
	rowAffected,err := h.bannerService.DeleteBanner(input.ID, input)

	if err != nil {
		errorMessage := gin.H{
			"message": err.Error(),
		}
		c.JSON(http.StatusBadRequest, errorMessage)
		return
	}

	response := gin.H{
		"rowAffected" :rowAffected,
	}
	c.JSON(http.StatusOK, response)
}