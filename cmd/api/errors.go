package main

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

var (
	ErrNotImplemented = errors.New("not implemented")
)
type ErrorResponse struct {
	Error string `json:"error"`
}
func notImplementedError(c *gin.Context) {
	response := ErrorResponse{
		Error: ErrNotImplemented.Error(),
	}
	c.JSON(http.StatusInternalServerError,response)
}