package main

import "github.com/gin-gonic/gin"

type DataResponse struct {
	Data any `json:"data"`
}
func jsonResponse(c *gin.Context, status int, data any) {
	r := DataResponse{
		Data: data,
	}
	c.JSON(status,r)
}