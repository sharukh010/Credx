package store

import "context"

type Storage struct {
	Cards interface {
		Add(context.Context,*Card) error 
		GetAll(context.Context) ([]Card,error)
	}
}

func NewStorage() Storage {
	return Storage{
		Cards: &CardStore{},
	}
}