# Credx

Credx is a backend application built for users to track their credit cards in a simple and organized way. It focuses on user authentication, secure card management, and a clean API structure that can keep improving as the project grows.

More than just an API, this project represents a learning journey in backend engineering. I started it from scratch with the goal of understanding how a real-world Go application evolves over time, from simple in-memory data handling to a structured database-backed system.

## The Journey

This project did not begin with PostgreSQL or an ORM.

It started with the basics:

- building the application from scratch in Go
- structuring the code into clear packages and handlers
- using slices as a temporary data store to understand the business logic first
- learning how CRUD flows work before introducing database complexity

That early version helped me move fast and validate the core idea: users should be able to register, log in, and track their credit cards cleanly.

Once the core flows were working, I iteratively improved the project:

1. I moved away from slice-based storage when persistence and scalability started to matter.
2. I introduced PostgreSQL as the real database layer.
3. I used GORM to model entities, manage relationships, and simplify database operations.
4. I added Docker Compose so the database setup became repeatable and easy to run.
5. I improved the API structure with authentication, protected routes, request validation, and Swagger documentation.

This README reflects that journey: starting simple, learning by building, and improving the system step by step instead of trying to make it perfect on day one.

## What Credx Does

Credx allows users to:

- register an account
- log in and receive a JWT token
- add credit cards
- view all saved credit cards
- fetch a specific card by ID
- update card details
- delete cards

Each card is tied to its owning user, and protected routes ensure users only work with their own data.

## Tech Stack

- Go
- Gin
- PostgreSQL
- GORM
- JWT authentication
- Docker Compose
- Swagger

## Current Architecture

The current version uses:

- `Gin` for routing and HTTP handling
- `GORM` for database access and model management
- `PostgreSQL` for persistent storage
- `JWT` for authentication and route protection
- `Swagger` for API documentation

Core entities:

- `User`
- `Card`

The application is organized into packages for API handlers, authentication, database setup, environment management, and storage logic.

## Project Structure

```text
credx/
├── cmd/api/            # API entrypoint, routes, handlers
├── internal/auth/      # JWT helpers
├── internal/db/        # DB bootstrap
├── internal/env/       # Environment helpers
├── internal/store/     # Models and storage layer
├── docs/               # Swagger generated docs
├── docker-compose.yml  # PostgreSQL service
├── Makefile            # Common commands
└── README.md
```

## Features

- User registration
- User login with JWT token generation
- Protected credit card routes
- Per-user card access
- Card masking support for stored card numbers
- PostgreSQL-backed persistence
- Auto migration using GORM
- Swagger API documentation

## API Overview

Base path:

```text
/v1
```

Public routes:

- `POST /auth/register`
- `POST /auth/log-in`
- `GET /health`

Protected routes:

- `GET /cards/`
- `GET /cards/:id`
- `POST /cards/`
- `PATCH /cards/:id`
- `DELETE /cards/:id`

Swagger UI:

```text
/swagger/index.html
```

## Getting Started

### 1. Clone the project

```bash
git clone https://github.com/sharukh010/credx.git
cd credx
```

### 2. Start PostgreSQL

```bash
docker compose up -d
```

The default database configuration in this project uses:

- database: `credx`
- user: `admin`
- password: `adminpassword`
- port: `5432`

### 3. Create `.env`

Create a `.env` file in the project root with values like:

```env
SERVER_ADDR=:8080
DB_ADDR=host=localhost user=admin password=adminpassword dbname=credx port=5432 sslmode=disable
ENV=development
JWT_SECRET=MY_SECRET
```

### 4. Run the API

```bash
go run ./cmd/api
```

If you use the provided Makefile and already have the required tools installed:

```bash
make swagger
make run
```

## Data Model Snapshot

### User

- `id`
- `user_name`
- `name`
- `gender`
- `email`
- `dob`
- `password`
- timestamps

### Card

- `id`
- `user_id`
- `name`
- `number`
- `expire_at`
- timestamps

## Why This Project Matters To Me

Credx is a practical record of my backend learning process.

I wanted to understand how to:

- start with a simple idea and make it work first
- build CRUD APIs from scratch
- introduce authentication into a real project
- migrate from temporary in-memory storage to a relational database
- use GORM effectively with PostgreSQL
- improve project structure iteratively instead of rewriting everything

This project shows that progression clearly. It began as a small hands-on experiment and grew into a more realistic backend application for managing credit card records.

## Future Improvements

- stronger card-number security and encryption strategy
- refresh token support
- better validation for expiry dates and card formats
- pagination and filtering for card lists
- unit and integration test coverage
- production-ready configuration and deployment

## Closing Note

Credx is a project built by learning in public through iteration. I started from scratch, used slices when simplicity mattered, moved to PostgreSQL when persistence mattered, and adopted GORM to build a cleaner and more maintainable backend.

That journey is the real story of this project.
