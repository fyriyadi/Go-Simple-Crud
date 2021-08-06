package controllers

import (
	"gocrud/models"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func GetToDoList(c *gin.Context) {
	dsn := "host=localhost user=golanguser password=passwordgolanguser dbname=golangdb port=5432 sslmode=disable TimeZone=Asia/Bangkok"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	var results []map[string]interface{}
	db.Table("todos").Find(&results)

	c.JSON(200, gin.H{
		"data": results,
	})

}

func GetTodoID(c *gin.Context) {
	id := c.Param("id")

	// dsn := "host=localhost user=golanguser password=passwordgolanguser dbname=golangdb port=5432 sslmode=disable TimeZone=Asia/Bangkok"
	db, err := gorm.Open(postgres.Open("host=localhost user=golanguser password=passwordgolanguser dbname=golangdb port=5432 sslmode=disable TimeZone=Asia/Bangkok"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	var results []map[string]interface{}
	db.Raw("SELECT title, body FROM todos WHERE id = ? ", id).Scan(&results)

	c.JSON(200, gin.H{
		"data": results,
	})

}

func AddTodo(c *gin.Context) {
	title := c.PostForm("title")
	body := c.PostForm("body")

	dsn := "host=localhost user=golanguser password=passwordgolanguse dbname=golangdb port=5432 sslmode=disable TimeZone=Asia/Bangkok"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		// panic("failed to connect database")
		c.JSON(200, gin.H{
			"message": "failed to connect to database",
		})
	}else{
		todo := models.Todo{Title: title, Body: body}
		db.Create(&todo)

		c.JSON(200, gin.H{
			"message": "data added",
		})

	}

	
}

func UpdateTodoID(c *gin.Context) {
	id := c.Param("id")
	title := c.PostForm("title")
	body := c.PostForm("body")

	dsn := "host=localhost user=golanguser password=passwordgolanguser dbname=golangdb port=5432 sslmode=disable TimeZone=Asia/Bangkok"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	var todo models.Todo
	db.First(&todo, id)
	todo.Title = title
	todo.Body = body
	db.Save(&todo)

	c.JSON(200, gin.H{
		"message": "data updated",
	})

}

func DeleteTodoID(c *gin.Context) {
	id := c.Param("id")

	dsn := "host=localhost user=golanguser password=passwordgolanguser dbname=golangdb port=5432 sslmode=disable TimeZone=Asia/Bangkok"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	db.Raw("SELECT title, body FROM todos WHERE id = ? ", id)

	var todo models.Todo
	db.First(&todo, id)
	db.Delete(&todo)

	c.JSON(200, gin.H{
		"message": "data deleted",
	})

}
