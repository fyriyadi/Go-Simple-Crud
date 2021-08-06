package controllers

import (
	"github.com/gin-gonic/gin"
)

func HomeView(c *gin.Context) {
	c.JSON(200, gin.H{
		"data": "Welcome Home",
	})
}
