package db 
import (
  "gorm.io/driver/postgres"
  "gorm.io/gorm"
)

func New(dsn string) (*gorm.DB,error){
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	return db,err 
}