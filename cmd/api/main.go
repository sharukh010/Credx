package main

import (
	"log"

	"github.com/sharukh010/credx/internal/store"
)

func main() {
	cfg := config{
		Addr: ":8080",
	}
	
	api := &application{
		config: cfg,
		store: store.NewStorage(),
	}

	mux := api.mount()

	if err := api.run(mux); err != nil {
		log.Fatalf("Server stopped running Error: %v\n",err.Error())
	}
}