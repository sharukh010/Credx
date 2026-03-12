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

	card := v1.Group("/cards")
	// credit card routes 
	card.GET("/",getCardsHandler) // get all the cards
	card.GET("/:id",getCardByIDHandler) // get card by id 
	card.POST("/",addCardHandler) // add card 
	card.PATCH("/:id",updateCardHandler) // update card details
	card.DELETE("/:id",deleteCardHandler) // delete card by id 

	// authentication routes
	auth := v1.Group("/auth")

	auth.POST("/register",userRegistrationHandler) // rotue to register user
	auth.GET("/log-in",userLoginHandler)


	return r 
}

func (app *application) run(mux http.Handler) error {
	log.Printf("Server started running on Port %s\n",app.Config.Addr)
	return http.ListenAndServe(app.Config.Addr,mux)
}