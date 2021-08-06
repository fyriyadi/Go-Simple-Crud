package main

import (
	// "net/http"
	"gocrud/controllers"
	"gocrud/models"

	// "fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {

	dsn := "host=localhost user=golanguser password=passwordgolanguser dbname=golangdb port=5432 sslmode=disable TimeZone=Asia/Jakarta"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("failed to connect database")
	}
	db.AutoMigrate(&models.Todo{})

	r := gin.Default()

	r.GET("/", controllers.HomeView)
	r.GET("/todolist", controllers.GetToDoList)
	r.GET("/todo/:id", controllers.GetTodoID)
	r.POST("/todo/add", controllers.AddTodo)
	r.PUT("/todo/:id", controllers.UpdateTodoID)
	r.DELETE("/todo/:id", controllers.DeleteTodoID)
	r.Run()

}
