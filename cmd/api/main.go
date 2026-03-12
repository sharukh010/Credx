package main

import "log"

func main() {
	cfg := config{
		Addr: ":8080",
	}
	
	api := &application{
		Config: cfg,
	}

	mux := api.mount()

	if err := api.run(mux); err != nil {
		log.Fatalf("Server stopped running Error: %v\n",err.Error())
	}
}