package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

var version = "1.0.0"

type application struct {
	Config config 
}

type config struct {
	Addr string 
}

func (app *application) mount() http.Handler{
	r := gin.Default()

	v1 := r.Group("/v1")
	// credit card routes 
	v1.GET("/card",getCardsHandler)
	v1.GET("/card/:id",getCardByIDHandler)
	// route to add card 
	v1.POST("/card",addCardHandler)
	v1.PATCH("/card/:id",updateCardHandler)
	v1.DELETE("/card/:id",deleteCardHandler)



	// authentication routes
	return r 
}

func (app *application) run(mux http.Handler) error {
	log.Printf("Server started running on Port %s\n",app.Config.Addr)
	return http.ListenAndServe(app.Config.Addr,mux)
}