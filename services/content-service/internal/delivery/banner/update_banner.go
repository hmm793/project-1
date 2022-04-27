package delivery

import (
	"content-service-v3/services/content-service/domain/dto"
	"content-service-v3/services/content-service/internal/delivery/helper"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *bannerHandler) UpdateBanner(c *gin.Context) {
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

	var input dto.UpdateBannerInput
	err = c.ShouldBind(&input)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{
			"message": errors,
		}
		c.JSON(http.StatusBadRequest, errorMessage)
		return
	}

	file, err := c.FormFile("image")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message" :err.Error()})
		return
	}
	
	banner, pathForSave, err := h.bannerService.UpdateBanner(inputID.ID, input,file)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message" :err.Error()})
		return
	}

	if pathForSave != "" {
		err = c.SaveUploadedFile(file, pathForSave)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"message" :err.Error()})
			return
		}
	}

	c.JSON(http.StatusOK, banner)
}