package main

import (
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/sharukh010/credx/internal/store"
)


var(
	version = "1.0.0"
	DatabaseOperationsTimeOut = 15 * time.Second
)

const (
	USERTOKEN = "User-Token"
)


type application struct {
	config config 
	store store.Storage
}

type config struct {
	Addr string 
	JWTSecret []byte 
	dbConfig dbConfig
	env string 
}

type dbConfig struct {
	Addr string 
}
func (app *application) mount() http.Handler{
	r := gin.Default()

	v1 := r.Group("/v1")
	// health route
	v1.GET("/health",app.getHealthHandler)

	card := v1.Group("/cards")
	card.Use(app.AuthMiddleware())
	// credit card routes 
	card.GET("/",app.getCardsHandler) // get all the cards
	card.GET("/:id",app.getCardByIDHandler) // get card by id 
	card.POST("/",app.addCardHandler) // add card 
	card.PATCH("/:id",app.updateCardHandler) // update card details
	card.DELETE("/:id",app.deleteCardHandler) // delete card by id 

	// authentication routes
	auth := v1.Group("/auth")

	auth.POST("/register",app.userRegistrationHandler) // rotue to register user
	auth.GET("/log-in",app.userLoginHandler)


	return r 
}

func (app *application) run(mux http.Handler) error {
	log.Printf("Server started running on Port %s\n",app.config.Addr)
	return http.ListenAndServe(app.config.Addr,mux)
}