package main

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sharukh010/credx/internal/store"
)

type AddCardRequest struct {
	Name string `json:"name" binding:"required,min=5"`
	Number string `json:"number" binding:"required,min=4,max=4"`
	ExpireAt string `json:"expire_at" binding:"required,min=5,max=5"`

}
func (app *application) getCardsHandler(c *gin.Context){
	notImplementedError(c)
}

func (app *application) getCardByIDHandler(c *gin.Context){
	notImplementedError(c)
}

func (app *application) addCardHandler(c *gin.Context){
	r := &AddCardRequest{}
	if err := c.BindJSON(r); err != nil {
		badRequestResponse(c,err)
		return 
	}
	
	card := &store.Card{
		Name: r.Name,
		ExpireAt: r.ExpireAt,
	}
	card.MaskNumber(r.Number)

	ctx,cancel := context.WithTimeout(c,DatabaseOperationsTimeOut)
	defer cancel()

	if err := app.store.Card.Add(ctx,card); err != nil {
		internalServerErrorResponse(c,err)
		return 
	}

	jsonResponse(c,http.StatusCreated,card)
}

func (app *application) updateCardHandler(c *gin.Context){
	notImplementedError(c)
}

func (app *application) deleteCardHandler(c *gin.Context){
	notImplementedError(c)
}