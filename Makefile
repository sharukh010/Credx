include .env

.PHONY: up-db down-db run swagger docs

up-db: 
	docker compose up --build

down-db: 
	docker compose down 

run: swagger
	air

swagger:
	swag init -g cmd/api/main.go -o docs

docs: swagger
