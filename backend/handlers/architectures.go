package handlers

import (
	"net/http"

	"github.com/SvenNellerz/datacenter-guard/backend/models"
	"github.com/SvenNellerz/datacenter-guard/backend/services"
	"github.com/gin-gonic/gin"
)

func GetArchitectures(c *gin.Context) {
	architectures, err := services.LoadArchitectures()
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.Error("Failed to load architectures"))
		return
	}
	c.JSON(http.StatusOK, models.SuccessWithCount(architectures, len(architectures)))
}

func GetPlaybooks(c *gin.Context) {
	playbooks, err := services.LoadPlaybooks()
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.Error("Failed to load playbooks"))
		return
	}
	c.JSON(http.StatusOK, models.SuccessWithCount(playbooks, len(playbooks)))
}
