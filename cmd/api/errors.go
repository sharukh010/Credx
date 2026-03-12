package main

import (
	"errors"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

var (
	ErrNotImplemented = errors.New("not implemented")
	ErrBadRequest = errors.New("bad request")
	ErrInternalServer = errors.New("something went wrong")
	ErrNotFound = errors.New("not found")
	ErrInvalidUpdateRequest = errors.New("invalid update request")
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

func badRequestResponse(c *gin.Context,err error){
	response := ErrorResponse{
		Error: ErrBadRequest.Error(),
	}
	log.Printf("Bad Request Error: %v\n",err.Error())
	c.JSON(http.StatusBadRequest,response)
}

func internalServerErrorResponse(c *gin.Context,err error){
	response := ErrorResponse{
		Error: ErrInternalServer.Error(),
	}
	log.Printf("Internal Server Error: %v\n",err.Error())
	c.JSON(http.StatusInternalServerError,response)
}

func notFoundResponse(c *gin.Context,err error){
	response := ErrorResponse{
		Error : ErrNotFound.Error(),
	}
	log.Printf("Not Found Error: %v\n",err.Error())
	c.JSON(http.StatusNotFound,response)
}