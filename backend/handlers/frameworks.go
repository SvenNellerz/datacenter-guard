package handlers

import (
	"net/http"

	"github.com/SvenNellerz/datacenter-guard/backend/models"
	"github.com/SvenNellerz/datacenter-guard/backend/services"
	"github.com/gin-gonic/gin"
)

func GetFrameworks(c *gin.Context) {
	frameworks, err := services.LoadFrameworks()
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.Error("Failed to load frameworks"))
		return
	}
	c.JSON(http.StatusOK, models.SuccessWithCount(frameworks, len(frameworks)))
}

func GetFrameworkChecklist(c *gin.Context) {
	id := c.Param("id")
	frameworks, err := services.LoadFrameworks()
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.Error("Failed to load frameworks"))
		return
	}

	for _, fw := range frameworks {
		if fw.ID == id {
			checklist := services.GenerateChecklist(fw)
			c.JSON(http.StatusOK, models.Success(checklist))
			return
		}
	}

	c.JSON(http.StatusNotFound, models.Error("Framework not found"))
}
