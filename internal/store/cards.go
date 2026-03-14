package store

import (
	"context"
	"database/sql"
	"strings"
	"time"

	"gorm.io/gorm"
)

// var cardID int64 = 0
type Card struct {
	ID     int64      `json:"id" gorm:"primaryKey;autoIncrement"`
	UserID int64	  `json:"user_id" gorm:"index;not null"`
	User User `json:"-" gorm:"foreignKey:UserID;references:ID"`
	Name   string     `json:"name" gorm:"not null"`
	Number string `json:"number" gorm:"not null; unique"`
	ExpireAt string `json:"expire_at" gorm:"not null"`
	CreatedAt time.Time `json:"created_at" gorm:"autoCreateTime:nano"`
	UpdatedAt time.Time `json:"updated_at" gorm:"autoUpdateTime:nano"`
}



func (c *Card) MaskNumber(number string){
	base := "XXXX"
	c.Number = strings.Repeat(base,3)
	c.Number += number
}

type CardStore struct {
	db *gorm.DB

}

func (s *CardStore) Add(ctx context.Context,card *Card) error {

	result := s.db.Create(card)
	if result.Error != nil {
		return result.Error
	}
	return nil 
}

func (s *CardStore) GetAll(ctx context.Context) ([]Card,error){
	userID := ctx.Value("UserID")
	cards := []Card{}
	result := s.db.Where("user_id = ?",userID).Order("id ASC").Find(&cards)
	if result.Error != nil {
		return nil,result.Error 
	}
	return cards,nil  
}

func (s *CardStore) GetByID(ctx context.Context,ID int64) (*Card,error){
	userID := ctx.Value("UserID")
	card := &Card{}
	result := s.db.Where("id = ?",ID).Where("user_id = ?",userID).Find(card)
	if result.Error != nil {
		return nil,result.Error 
	}
	if result.RowsAffected == 0 {
		return nil,sql.ErrNoRows
	}
	return card,nil  

}

func (s *CardStore) Update(ctx context.Context,c *Card) error {

	result := s.db.Where("id = ?",c.ID).Where("user_id = ?",c.UserID).Updates(c)
	if result.Error != nil {
		return result.Error 
	}
	if result.RowsAffected == 0 {
		return sql.ErrNoRows
	}
	return nil 
}

func (s *CardStore) Delete(ctx context.Context,ID int64) error {
	userID := ctx.Value("UserID")
	card := &Card{}
	result := s.db.Where("user_id = ?",userID).Where("id = ?",ID).Delete(card)
	if result.Error != nil {
		return result.Error 
	}
	if result.RowsAffected == 0 {
		return sql.ErrNoRows
	}
	
	return nil 
}