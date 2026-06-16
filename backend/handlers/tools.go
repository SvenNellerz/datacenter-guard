package handlers

import (
	"net/http"
	"strings"

	"github.com/SvenNellerz/datacenter-guard/backend/models"
	"github.com/SvenNellerz/datacenter-guard/backend/services"
	"github.com/gin-gonic/gin"
)

func GetTools(c *gin.Context) {
	tools, err := services.LoadTools()
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.Error("Failed to load tools"))
		return
	}

	category := c.Query("category")
	pricing := c.Query("pricing")
	search := c.Query("search")

	filtered := tools
	if category != "" {
		filtered = filterByCategory(filtered, category)
	}
	if pricing != "" {
		filtered = filterByPricing(filtered, pricing)
	}
	if search != "" {
		filtered = filterBySearch(filtered, search)
	}

	c.JSON(http.StatusOK, models.SuccessWithCount(filtered, len(filtered)))
}

func GetToolByID(c *gin.Context) {
	id := c.Param("id")
	tools, err := services.LoadTools()
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.Error("Failed to load tools"))
		return
	}

	for _, tool := range tools {
		if tool.ID == id {
			c.JSON(http.StatusOK, models.Success(tool))
			return
		}
	}

	c.JSON(http.StatusNotFound, models.Error("Tool not found"))
}

func filterByCategory(tools []models.Tool, category string) []models.Tool {
	var result []models.Tool
	for _, t := range tools {
		if strings.EqualFold(t.Category, category) {
			result = append(result, t)
		}
	}
	return result
}

func filterByPricing(tools []models.Tool, pricing string) []models.Tool {
	var result []models.Tool
	for _, t := range tools {
		if strings.EqualFold(t.Pricing, pricing) {
			result = append(result, t)
		}
	}
	return result
}

func filterBySearch(tools []models.Tool, search string) []models.Tool {
	search = strings.ToLower(search)
	var result []models.Tool
	for _, t := range tools {
		if strings.Contains(strings.ToLower(t.Name), search) ||
			strings.Contains(strings.ToLower(t.Description), search) {
			result = append(result, t)
		}
	}
	return result
}
