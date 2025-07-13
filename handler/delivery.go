package handler

import (
	"net/http"
	"targeting-engine/repo"
	"targeting-engine/service"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine) {
	v1 := r.Group("/v1")
	{
		v1.GET("/delivery", deliveryHandler())
	}
}

func deliveryHandler() gin.HandlerFunc {
	repository := repo.NewMemoryRepository()
	svc := service.NewTargetingService(repository)

	return func(c *gin.Context) {
		app := c.Query("app")
		country := c.Query("country")
		os := c.Query("os")

		if app == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "missing app param"})
			return
		}
		if country == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "missing country param"})
			return
		}
		if os == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "missing os param"})
			return
		}

		results := svc.GetEligibleCampaigns(app, country, os)
		if len(results) == 0 {
			c.Status(http.StatusNoContent)
			return
		}
		c.JSON(http.StatusOK, results)
	}
}
