package main

import "github.com/sharukh010/credx/internal/store"

type HealthResponseDoc struct {
	Status      string `json:"status"`
	Version     string `json:"version"`
	Environment string `json:"environment"`
}

type CardDataResponse struct {
	Data store.Card `json:"data"`
}

type CardsDataResponse struct {
	Data []store.Card `json:"data"`
}

type UserDataResponse struct {
	Data store.User `json:"data"`
}

type LoginDataResponse struct {
	Data userLoginResponse `json:"data"`
}
