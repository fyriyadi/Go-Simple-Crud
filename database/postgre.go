package database

import (
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func PostgresqlConnect() *gorm.DB {
	database := "host=localhost user=golanguser password=passwordgolanguser dbname=golangdb port=5432 sslmode=disable TimeZone=Asia/Bangkok"
	db, err := gorm.Open(postgres.Open(database), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	// Get generic database object sql.DB to use its functions
	sqlDB, _ := db.DB()

	// SetMaxIdleConns sets the maximum number of connections in the idle connection pool.
	sqlDB.SetMaxIdleConns(10)

	// SetMaxOpenConns sets the maximum number of open connections to the database.
	sqlDB.SetMaxOpenConns(50)

	// SetConnMaxLifetime sets the maximum amount of time a connection may be reused.
	sqlDB.SetConnMaxLifetime(time.Hour / 2)

	// migrate model / entity
	// db.AutoMigrate(&models.Todo{})

	return db
}
