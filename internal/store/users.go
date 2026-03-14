package store

import (
	"context"
	"database/sql"
	"time"

	"golang.org/x/crypto/bcrypt"
)

var UserID int64 = 0 
type User struct {
	ID       int64    `json:"id"`
	UserName string   `json:"user_name"`
	Name     Name     `json:"name"`
	Gender   string   `json:"gender"`
	Email    string   `json:"email"`
	DOB      string   `json:"dob"`
	Password string `json:"-"`
	Version int `json:"version"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type Name struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

type UserStore struct {

}

func HashPassword(plain string) (string,error) {
	hash,err := bcrypt.GenerateFromPassword([]byte(plain),bcrypt.DefaultCost)
	if err != nil {
		return "",err 
	}
	hashStr := string(hash)
	return hashStr,nil 
}

func CompareHashAndPassword(hash,plain string) error {
	return bcrypt.CompareHashAndPassword([]byte(hash),[]byte(plain))
}

func (u *User) SetID(ID int64){
	u.ID = ID 
}

func (u *User) SetCreatedAt(t time.Time){
	u.CreatedAt = t
}

func (u *User) SetUpdatedAt(t time.Time){
	u.UpdatedAt = t 
}

func (u *User) updateVersion(){
	u.Version += 1 
}

func (s *UserStore) Create(ctx context.Context,u *User) error {
	setID(u,&UserID)
	setCreatedAt(u)
	setUpdatedAt(u)
	u.updateVersion()
	Users = append(Users, *u)
	return nil 

}

func (s *UserStore) GetByUserName(ctx context.Context,uname string) (*User,error) {
	for _,user := range Users {
		if user.UserName == uname {
			return &user,nil 
		}
	}

	return nil,sql.ErrNoRows
}