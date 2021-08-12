package main

import (
	// "net/http"
	"github.com/fyriyadi/Go-Simple-Crud/controllers"
	"github.com/fyriyadi/Go-Simple-Crud/database"
	"github.com/fyriyadi/Go-Simple-Crud/server"

	// "fmt"
	"github.com/gin-gonic/gin"
)

func main() {

	// dsn := "host=localhost user=golanguser password=passwordgolanguser dbname=golangdb port=5432 sslmode=disable TimeZone=Asia/Jakarta"
	// db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	// if err != nil {
	// 	panic("failed to connect database")
	// }
	// db.AutoMigrate(&models.Todo{})

	r := gin.Default()
	db := database.PostgresqlConnect()
	inDB := &controllers.InDB{DB: db}
	server.Routes(r, inDB)
	r.Run()
	// r.GET("/todolist", inDB.GetToDoList)

}
