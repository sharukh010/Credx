include .env

.PHONY:up-db
up-db: 
	docker compose up --build

.PHONY:down-db 
down-db: 
	docker compose down 

run:
	air