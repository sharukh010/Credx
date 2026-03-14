package db

import (
	"fmt"

	"github.com/sharukh010/credx/internal/store"
	"gorm.io/gorm"
)

type TestFixtures struct {
	Users []store.User
	Cards []store.Card
}

// PrepareTestDB migrates the schema, clears old rows, and inserts seed data.
func PrepareTestDB(db *gorm.DB) (*TestFixtures, error) {
	if err := db.AutoMigrate(&store.User{}, &store.Card{}); err != nil {
		return nil, err
	}

	if err := resetTestDB(db); err != nil {
		return nil, err
	}

	return seedTestData(db)
}

func resetTestDB(db *gorm.DB) error {
	if err := db.Exec("TRUNCATE TABLE cards, users RESTART IDENTITY CASCADE").Error; err != nil {
		return fmt.Errorf("reset test db: %w", err)
	}

	return nil
}

func seedTestData(db *gorm.DB) (*TestFixtures, error) {
	alicePassword, err := store.HashPassword("pass1234")
	if err != nil {
		return nil, fmt.Errorf("hash alice password: %w", err)
	}

	bobPassword, err := store.HashPassword("pass5678")
	if err != nil {
		return nil, fmt.Errorf("hash bob password: %w", err)
	}

	users := []store.User{
		{
			UserName: "alice_user",
			Name: store.Name{
				FirstName: "Alice",
				LastName:  "Shaw",
			},
			Gender:   "female",
			Email:    "alice@example.com",
			DOB:      "01/01/1998",
			Password: alicePassword,
		},
		{
			UserName: "bob_user",
			Name: store.Name{
				FirstName: "Bobby",
				LastName:  "Stone",
			},
			Gender:   "male",
			Email:    "bob@example.com",
			DOB:      "02/02/1996",
			Password: bobPassword,
		},
	}

	if err := db.Create(&users).Error; err != nil {
		return nil, fmt.Errorf("seed users: %w", err)
	}

	cards := []store.Card{
		{
			UserID:   users[0].ID,
			Name:     "Alice Primary",
			Number:   "XXXXXXXXXXXX1234",
			ExpireAt: "12/28",
		},
		{
			UserID:   users[0].ID,
			Name:     "Alice Backup",
			Number:   "XXXXXXXXXXXX5678",
			ExpireAt: "10/29",
		},
		{
			UserID:   users[1].ID,
			Name:     "Bob Travel",
			Number:   "XXXXXXXXXXXX9876",
			ExpireAt: "09/27",
		},
	}

	if err := db.Create(&cards).Error; err != nil {
		return nil, fmt.Errorf("seed cards: %w", err)
	}

	return &TestFixtures{
		Users: users,
		Cards: cards,
	}, nil
}
