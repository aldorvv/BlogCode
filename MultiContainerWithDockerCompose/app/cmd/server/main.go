package main

import (
	"github.com/BlogCode/MultiContainer/internal/views"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.GET("/todos", views.GetTodos)
	router.GET("/todos/:id", views.GetTodo)
	router.POST("/todos", views.PostTodo)
	router.PUT("/todos/:id", views.UpdateTodo)
	router.DELETE("/todos/:id", views.SoftDeleteTodo)

	router.Run("0.0.0.0:8080")
}
