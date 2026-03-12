package main

import (
	"context"
	"database/sql"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/sharukh010/credx/internal/store"
)

type AddCardRequest struct {
	Name string `json:"name" binding:"required,min=5"`
	Number string `json:"number" binding:"required,min=4,max=4"`
	ExpireAt string `json:"expire_at" binding:"required,min=5,max=5"`
}

type UpdateCardRequest struct {
	Name *string `json:"name" binding:"omitempty,min=5"`
	Number *string `json:"number" binding:"omitempty,min=4,max=4"`
	ExpireAt *string `json:"expire_at" binding:"omitempty,min=5,max=5"`
}

func (app *application) getCardsHandler(c *gin.Context){
	ctx,cancel := context.WithTimeout(c,DatabaseOperationsTimeOut)
	defer cancel()

	cards,err := app.store.Cards.GetAll(ctx)
	if err != nil {
		internalServerErrorResponse(c,err)
		return 
	}
	jsonResponse(c,http.StatusOK,cards)
}

func (app *application) getCardByIDHandler(c *gin.Context){
	id,err := strconv.ParseInt(c.Param("id"),10,64)
	if err != nil {
		internalServerErrorResponse(c,err)
		return 
	}

	ctx,cancel := context.WithTimeout(c,DatabaseOperationsTimeOut)
	defer cancel()
	
	card, err := app.store.Cards.GetByID(ctx,id)
	if err != nil {
		switch err {
		case sql.ErrNoRows:
			notFoundResponse(c,err)
		default:
			internalServerErrorResponse(c,err)
		}
		return 
	}
	jsonResponse(c,http.StatusOK,card)
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

	if err := app.store.Cards.Add(ctx,card); err != nil {
		internalServerErrorResponse(c,err)
		return 
	}

	jsonResponse(c,http.StatusCreated,card)
}

func (app *application) updateCardHandler(c *gin.Context){
	id,err := strconv.ParseInt(c.Param("id"),10,64)
	if err != nil {
		internalServerErrorResponse(c,err)
		return 
	}

	ctx,cancel := context.WithTimeout(c,DatabaseOperationsTimeOut)
	defer cancel()

	card,err := app.store.Cards.GetByID(ctx,id)
	if err != nil {
		switch err {
		case sql.ErrNoRows:
			notFoundResponse(c,err)
		default:
			internalServerErrorResponse(c,err)
		}
		return 
	}

	r := &UpdateCardRequest{}
	if err := c.BindJSON(r); err != nil {
		badRequestResponse(c,err)
		return 
	}

	if r.Name == nil && r.Number == nil && r.ExpireAt == nil {
		badRequestResponse(c,ErrInvalidUpdateRequest)
		return 
	}

	if r.Name != nil {
		card.Name  = *r.Name
	}

	if r.ExpireAt != nil {
		card.ExpireAt = *r.ExpireAt
	}
	if r.Number != nil {
		card.MaskNumber(*r.Number)
	}

	ctx,cancel = context.WithTimeout(c,DatabaseOperationsTimeOut)
	defer cancel()

	if err := app.store.Cards.Update(ctx,card); err != nil {
		internalServerErrorResponse(c,err)
		return 
	}
	jsonResponse(c,http.StatusOK,card)
}

func (app *application) deleteCardHandler(c *gin.Context){
	id,err := strconv.ParseInt(c.Param("id"),10,64)
	if err != nil {
		internalServerErrorResponse(c,err)
		return 
	}

	ctx,cancel := context.WithTimeout(c,DatabaseOperationsTimeOut)
	defer cancel()

	if err := app.store.Cards.Delete(ctx,id); err != nil {
		switch err {
		case sql.ErrNoRows:
			notFoundResponse(c,err)
		default:
			internalServerErrorResponse(c,err)
		}
		return 
	}

	jsonResponse(c,http.StatusNoContent,nil)
}