package store

import (
	"context"
	"strings"
	"time"
)

var cardID int64 = 0
type Card struct {
	ID     int64      `json:"id"`
	Name   string     `json:"name"`
	Number cardNumber 
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

type cardNumber struct {
	Number string `json:"number"`
}

func (cn *cardNumber) GenerateNumber(num string) {
	base := "XXXX"
	cn.Number = strings.Repeat(base,3)
	cn.Number += num 
}


// func generateCardID() int64 {
// 	cardID += 1 
// 	return cardID
// }

type CardStore struct {

}

func (s *CardStore) Add(ctx context.Context,card Card) error {

	setID(&card,&cardID)
	setCreatedAt(&card)
	setUpdatedAt(&card)
	card.updateVersion()

	return nil 
}