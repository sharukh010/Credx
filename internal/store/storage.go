package store

import "context"

type Storage struct {
	Card interface {
		Add(context.Context,*Card) error 
	}
}

func NewStorage() Storage {
	return Storage{
		Card: &CardStore{},
	}
}