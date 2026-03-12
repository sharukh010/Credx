package store

import (
	"context"
)

type Storage struct {
	Cards interface {
		Add(context.Context,*Card) error 
		GetAll(context.Context) ([]Card,error)
		GetByID(context.Context,int64) (*Card,error)
		Update(context.Context,*Card) error 
		Delete(context.Context,int64) error 
	}
}

func NewStorage() Storage {
	return Storage{
		Cards: &CardStore{},
	}
}