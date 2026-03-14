package store

import (
	"context"
	"database/sql"
	"strings"
	"time"

	"gorm.io/gorm"
)

var cardID int64 = 0
type Card struct {
	ID     int64      `json:"id"`
	UserID int64	  `json:"user_id"`
	Name   string     `json:"name"`
	Number string `json:"number"`
	ExpireAt string `json:"expire_at"`
	Version int `json:"version"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (c *Card) SetID(ID int64){
	c.ID = ID
}

func (c *Card) SetCreatedAt(t time.Time){
	c.CreatedAt = t
}

func (c *Card) SetUpdatedAt(t time.Time){
	c.UpdatedAt = t 
}

func (c *Card) updateVersion(){
	c.Version += 1
}

func (c *Card) MaskNumber(number string){
	base := "XXXX"
	c.Number = strings.Repeat(base,3)
	c.Number += number
}


// func generateCardID() int64 {
// 	cardID += 1 
// 	return cardID
// }

type CardStore struct {
	db *gorm.DB

}

func (s *CardStore) Add(ctx context.Context,card *Card) error {

	setID(card,&cardID)
	setCreatedAt(card)
	setUpdatedAt(card)
	card.updateVersion()

	Cards = append(Cards, *card)

	return nil 
}

func (s *CardStore) GetAll(ctx context.Context) ([]Card,error){
	return Cards,nil 
}

func (s *CardStore) GetByID(ctx context.Context,ID int64) (*Card,error){
	for _,card := range Cards {
		if card.ID == ID {
			return &card,nil
		}
	}
	return nil,sql.ErrNoRows
}

func (s *CardStore) Update(ctx context.Context,c *Card) error {
	var idx int
	idx = getCardIndex(c.ID)
	if idx == -1 {
		return sql.ErrNoRows
	}
	c.updateVersion()
	Cards[idx] = *c 
	return nil 
}

func (s *CardStore) Delete(ctx context.Context,ID int64) error {
	var idx int
	idx = getCardIndex(ID)
	if idx == -1 {
		return sql.ErrNoRows
	}
	Cards = append(Cards[:idx],Cards[idx+1:]...)
	return nil 
}

func getCardIndex(ID int64) int {
	var exists bool 
	var idx int
	for i,card := range Cards {
		if card.ID == ID {
			exists = true
			idx = i
			break
		}
	}
	if !exists {
		return -1
	}
	return idx 
}