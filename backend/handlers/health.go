package handlers

import (
	"net/http"
	"runtime"
	"time"

	"github.com/gin-gonic/gin"
)

var startTime = time.Now()

func Health(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status":  "healthy",
		"service": "datacenter-guard-api",
		"version": "1.0.0",
		"uptime":  time.Since(startTime).String(),
		"go":      runtime.Version(),
	})
}
