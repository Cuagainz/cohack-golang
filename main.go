package main

import (
	"github.com/cohack-golang/api/todo"
	"github.com/cohack-golang/models"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func init() {
	models.InitModel()
}

func main() {
	defer models.CloseDB()

	g := gin.New()
	g.Use(gin.Logger())
	g.Use(gin.Recovery())

	todoApi := g.Group("/api/todo")

	todoApi.Use()
	{
		todoApi.POST("resolve", todo.ResolveTodo)
		todoApi.POST("", todo.CreateTodo)
		todoApi.GET("", todo.GetTodo)
		todoApi.PUT("", todo.UpdateTodo)
		todoApi.DELETE("", todo.DeleteTodo)
	}

	// run on 8888 for the servers
	err := g.Run(":8888")
	if err != nil {
		logrus.Fatal(err)
	}
}
