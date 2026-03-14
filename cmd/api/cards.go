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

// getCardsHandler godoc
// @Summary List cards
// @Description Returns all cards for the authenticated user context used by the application.
// @Tags cards
// @Security BearerAuth
// @Produce json
// @Success 200 {object} CardsDataResponse
// @Failure 500 {object} ErrorResponse
// @Router /cards/ [get]
func (app *application) getCardsHandler(c *gin.Context){
	userID := c.Value("userID").(int64)
	ctx,cancel := context.WithTimeout(c,DatabaseOperationsTimeOut)
	defer cancel()
	ctx = context.WithValue(ctx,"UserID",userID)

	cards,err := app.store.Cards.GetAll(ctx)
	if err != nil {
		internalServerErrorResponse(c,err)
		return 
	}
	jsonResponse(c,http.StatusOK,cards)
}

// getCardByIDHandler godoc
// @Summary Get card by ID
// @Description Returns a card by its ID.
// @Tags cards
// @Security BearerAuth
// @Produce json
// @Param id path int true "Card ID"
// @Success 200 {object} CardDataResponse
// @Failure 404 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /cards/{id} [get]
func (app *application) getCardByIDHandler(c *gin.Context){
	userID := c.Value("userID").(int64)
	id,err := strconv.ParseInt(c.Param("id"),10,64)
	if err != nil {
		internalServerErrorResponse(c,err)
		return 
	}

	ctx,cancel := context.WithTimeout(c,DatabaseOperationsTimeOut)
	defer cancel()
	ctx = context.WithValue(ctx,"UserID",userID)
	
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

// addCardHandler godoc
// @Summary Add card
// @Description Creates a new card.
// @Tags cards
// @Security BearerAuth
// @Accept json
// @Produce json
// @Param request body AddCardRequest true "Card payload"
// @Success 201 {object} CardDataResponse
// @Failure 400 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /cards/ [post]
func (app *application) addCardHandler(c *gin.Context){
	userID := c.Value("userID").(int64)
	
	r := &AddCardRequest{}
	if err := c.BindJSON(r); err != nil {
		badRequestResponse(c,err)
		return 
	}
	
	card := &store.Card{
		Name: r.Name,
		ExpireAt: r.ExpireAt,
		UserID: userID,
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

// updateCardHandler godoc
// @Summary Update card
// @Description Updates card details by ID.
// @Tags cards
// @Security BearerAuth
// @Accept json
// @Produce json
// @Param id path int true "Card ID"
// @Param request body UpdateCardRequest true "Updated card payload"
// @Success 200 {object} CardDataResponse
// @Failure 400 {object} ErrorResponse
// @Failure 404 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /cards/{id} [patch]
func (app *application) updateCardHandler(c *gin.Context){
	userID := c.Value("userID").(int64)
	id,err := strconv.ParseInt(c.Param("id"),10,64)
	if err != nil {
		internalServerErrorResponse(c,err)
		return 
	}

	ctx,cancel := context.WithTimeout(c,DatabaseOperationsTimeOut)
	defer cancel()
	ctx = context.WithValue(ctx,"UserID",userID)

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
	ctx = context.WithValue(ctx,"UserID",userID)

	if err := app.store.Cards.Update(ctx,card); err != nil {
		internalServerErrorResponse(c,err)
		return 
	}
	jsonResponse(c,http.StatusOK,card)
}

// deleteCardHandler godoc
// @Summary Delete card
// @Description Deletes a card by ID.
// @Tags cards
// @Security BearerAuth
// @Produce json
// @Param id path int true "Card ID"
// @Success 204 {object} DataResponse
// @Failure 404 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /cards/{id} [delete]
func (app *application) deleteCardHandler(c *gin.Context){
	userID := c.Value("userID").(int64)
	id,err := strconv.ParseInt(c.Param("id"),10,64)
	if err != nil {
		internalServerErrorResponse(c,err)
		return 
	}

	ctx,cancel := context.WithTimeout(c,DatabaseOperationsTimeOut)
	defer cancel()
	ctx = context.WithValue(ctx,"UserID",userID)

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
