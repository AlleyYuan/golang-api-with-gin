package main

import (
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"myproject/api"
	"myproject/middleware"
	"myproject/model"
)

func main() {
	model.InitDatabase()
	log.Info("this is a log")

	router := gin.Default()
	v1 := router.Group("/api/v1")
	v1.POST("/register", api.Register)
	v1.POST("/login", api.Login)
	v1.Use(middleware.Auth())
	{
		v1.POST("/createtodos", api.CreateTodo)
		v1.GET("/gettodos", api.FetchAllTodo)
		v1.PUT("/updatetodos", api.UpdateTodo)
		v1.DELETE("/deletetodos", api.DeleteTodo)
	}
	router.Run(":8000")
}
