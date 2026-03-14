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

// getHealthHandler godoc
// @Summary Health check
// @Description Returns API health information.
// @Tags health
// @Produce json
// @Success 200 {object} HealthResponseDoc
// @Router /health [get]
func (app *application) getHealthHandler(c *gin.Context){
	r := HealthResponse{
		Status: "Alive",
		Version: version,
		Environment: app.config.env,
	}

	c.JSON(http.StatusOK,r)
}
