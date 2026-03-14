package store

import (
	"context"

	"gorm.io/gorm"
)

type Storage struct {
	Cards interface {
		Add(context.Context,*Card) error 
		GetAll(context.Context) ([]Card,error)
		GetByID(context.Context,int64) (*Card,error)
		Update(context.Context,*Card) error 
		Delete(context.Context,int64) error 
	}
	Users interface {
		Create(context.Context,*User) error 
		GetByUserName(context.Context, string) (*User,error)
	}
}

func NewStorage(db *gorm.DB) Storage {
	return Storage{
		Cards: &CardStore{db:db},
		Users: &UserStore{db:db},
	}
}