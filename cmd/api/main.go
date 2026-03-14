// @title Credx API
// @version 1.0
// @description Swagger documentation for the Credx APIs.
// @BasePath /v1
// @schemes http https
// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
package main

import (
	"log"

	"github.com/joho/godotenv"
	"github.com/sharukh010/credx/internal/db"
	"github.com/sharukh010/credx/internal/env"
	"github.com/sharukh010/credx/internal/store"
	_ "github.com/sharukh010/credx/docs"
)

func main() {
	if err := godotenv.Load();err != nil {
		log.Fatalln("Error loading .env file")
	}

	cfg := config{
		Addr: env.GetString("SERVER_ADDR",":8080"),
		dbConfig: dbConfig{
			Addr: env.GetString("DB_ADDR","host=localhost user=admin password=adminpassword dbname=credx port=5432 sslmode=disable"),
		},
		env: env.GetString("ENV","development"),
		JWTSecret: []byte(env.GetString("JWT_SECRET","MY_SECRET")),
	}
	
	db,err := db.New(cfg.dbConfig.Addr)
	if err != nil {
		log.Fatalf("Unable to Connect to DB Error: %v\n",err)
	}

	//migration 
	db.AutoMigrate(&store.Card{})
	db.AutoMigrate(&store.User{})

	store := store.NewStorage(db)

	api := &application{
		config: cfg,
		store: store,
	}

	mux := api.mount()

	if err := api.run(mux); err != nil {
		log.Fatalf("Server stopped running Error: %v\n",err.Error())
	}
}
