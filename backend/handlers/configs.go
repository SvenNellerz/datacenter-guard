package handlers

import (
	"net/http"

	"github.com/SvenNellerz/datacenter-guard/backend/models"
	"github.com/SvenNellerz/datacenter-guard/backend/services"
	"github.com/gin-gonic/gin"
)

func GetConfigTemplates(c *gin.Context) {
	templates, err := services.LoadTemplates()
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.Error("Failed to load templates"))
		return
	}
	c.JSON(http.StatusOK, models.SuccessWithCount(templates, len(templates)))
}

func GenerateConfig(c *gin.Context) {
	var req models.GenerateRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, models.Error("Invalid request: "+err.Error()))
		return
	}

	content, err := services.GenerateConfig(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.Error("Failed to generate config: "+err.Error()))
		return
	}

	c.JSON(http.StatusOK, models.Success(content))
}

func Export(c *gin.Context) {
	c.JSON(http.StatusOK, models.Success(gin.H{"message": "Export functionality coming soon"}))
}
