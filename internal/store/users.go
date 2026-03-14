package store

import (
	"context"
	"database/sql"
	"time"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

var UserID int64 = 0

type User struct {
	ID        int64     `json:"id" gorm:"primaryKey;autoIncrement"`
	UserName  string    `json:"user_name" gorm:"index unique"`
	Name      Name      `json:"name" gorm:"embedded"`
	Gender    string    `json:"gender"`
	Email     string    `json:"email" gorm:"unique"`
	DOB       string    `json:"dob"`
	Password  string    `json:"-"`
	CreditCards []Card `json:"credit_cards" gorm:"foreignKey:UserID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	CreatedAt time.Time `json:"created_at" gorm:"autoCreateTime:nano"`
	UpdatedAt time.Time `json:"updated_at" gorm:"autoUpdateTime:nano"`
}

type Name struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

type UserStore struct {
	db *gorm.DB
}

func HashPassword(plain string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(plain), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	hashStr := string(hash)
	return hashStr, nil
}

func CompareHashAndPassword(hash, plain string) error {
	return bcrypt.CompareHashAndPassword([]byte(hash), []byte(plain))
}


func (s *UserStore) Create(ctx context.Context, u *User) error {

	result := s.db.Create(u)
	if result.Error != nil {
		return result.Error
	}
	return nil

}

func (s *UserStore) GetByUserName(ctx context.Context, uname string) (*User, error) {
	
	user := &User{}

	result := s.db.Where("user_name = ?",uname).Find(user)
	if result.Error != nil {
		return nil,result.Error 
	}
	if result.RowsAffected == 0 {
		return nil,sql.ErrNoRows
	}
	return user,nil  
}
