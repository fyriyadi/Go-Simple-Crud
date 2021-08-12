package controllers

import (
	"net/http"

	"github.com/fyriyadi/Go-Simple-Crud/models"

	"github.com/gin-gonic/gin"
)

func (inDB *InDB) GetToDoList(c *gin.Context) {
	var results []map[string]interface{}
	err := inDB.DB.Table("todos").Find(&results).Error
	if err == nil {
		c.JSON(http.StatusOK, gin.H{
			"data": results,
		})
	}

}

func (inDB *InDB) GetTodoID(c *gin.Context) {
	id := c.Param("id")

	var results []map[string]interface{}
	inDB.DB.Raw("SELECT title, body FROM todos WHERE id = ? ", id).Scan(&results)

	c.JSON(http.StatusOK, gin.H{
		"data": results,
	})

}

func (inDB *InDB) AddTodo(c *gin.Context) {
	title := c.PostForm("title")
	body := c.PostForm("body")

	todo := models.Todo{Title: title, Body: body}
	inDB.DB.Create(&todo)

	c.JSON(http.StatusCreated, gin.H{
		"message": "data added",
	})

}

func (inDB *InDB) UpdateTodoID(c *gin.Context) {
	id := c.Param("id")
	title := c.PostForm("title")
	body := c.PostForm("body")

	var todo models.Todo
	inDB.DB.First(&todo, id)
	todo.Title = title
	todo.Body = body
	inDB.DB.Save(&todo)

	c.JSON(http.StatusOK, gin.H{
		"message": "data updated",
	})

}

func (inDB *InDB) DeleteTodoID(c *gin.Context) {
	id := c.Param("id")

	inDB.DB.Raw("SELECT title, body FROM todos WHERE id = ? ", id)

	var todo models.Todo
	inDB.DB.First(&todo, id)
	inDB.DB.Delete(&todo)

	c.JSON(http.StatusOK, gin.H{
		"message": "data deleted",
	})

}
