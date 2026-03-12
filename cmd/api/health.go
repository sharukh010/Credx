package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type HealthResponse struct {
	Status string `json:"status"`
	Version string `json:"version"`
	Environment string `json:"environment"`
}
func getHealthHandler(c *gin.Context){
	r := HealthResponse{
		Status: "Alive",
		Version: version,
		Environment: "development",
	}

	c.JSON(http.StatusOK,r)
}