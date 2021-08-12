package controllers

import (
	"github.com/fyriyadi/Go-Simple-Crud/database"
	"gorm.io/gorm"
)

type InDB struct {
	DB *gorm.DB
}

func ConnectDB() *InDB {
	db := database.PostgresqlConnect()
	inDB := &InDB{DB: db}
	return inDB
}
