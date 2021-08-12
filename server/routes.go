package server

import (
	"github.com/fyriyadi/Go-Simple-Crud/controllers"
	"github.com/gin-gonic/gin"
)

func Routes(router *gin.Engine, inDB *controllers.InDB) {
	router.GET("/", inDB.HomeView)

	router.GET("/todolist", inDB.GetToDoList)
	router.GET("/todo/:id", inDB.GetTodoID)
	router.POST("/todo/add", inDB.AddTodo)
	router.PUT("/todo/:id", inDB.UpdateTodoID)
	router.DELETE("/todo/:id", inDB.DeleteTodoID)
}
