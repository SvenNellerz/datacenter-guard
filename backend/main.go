package main

import (
	"log"
	"os"

	"github.com/SvenNellerz/datacenter-guard/backend/handlers"
	"github.com/SvenNellerz/datacenter-guard/backend/middleware"
	"github.com/gin-gonic/gin"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	r := gin.Default()
	r.Use(middleware.CORS())

	api := r.Group("/api")
	{
		api.GET("/health", handlers.Health)
		api.GET("/tools", handlers.GetTools)
		api.GET("/tools/:id", handlers.GetToolByID)
		api.GET("/frameworks", handlers.GetFrameworks)
		api.GET("/frameworks/:id/checklist", handlers.GetFrameworkChecklist)
		api.GET("/configs/templates", handlers.GetConfigTemplates)
		api.POST("/configs/generate", handlers.GenerateConfig)
		api.GET("/playbooks", handlers.GetPlaybooks)
		api.GET("/architectures", handlers.GetArchitectures)
		api.POST("/export", handlers.Export)
	}

	log.Printf("DataCenter Guard API starting on port %s", port)
	if err := r.Run(":" + port); err != nil {
		log.Fatal(err)
	}
}
