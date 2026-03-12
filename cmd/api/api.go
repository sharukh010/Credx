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
	// health route
	v1.GET("/health",getHealthHandler)

	card := v1.Group("/card")
	// credit card routes 
	card.GET("/",getCardsHandler)
	card.GET("/:id",getCardByIDHandler)
	// route to add card 
	card.POST("/",addCardHandler)
	card.PATCH("/:id",updateCardHandler)
	card.DELETE("/:id",deleteCardHandler)


	// authentication routes
	return r 
}

func (app *application) run(mux http.Handler) error {
	log.Printf("Server started running on Port %s\n",app.Config.Addr)
	return http.ListenAndServe(app.Config.Addr,mux)
}