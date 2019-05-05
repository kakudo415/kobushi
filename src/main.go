package main

import (
	"github.com/gin-gonic/gin"

	"./api"
	"./page"
)

func main() {
	engine := gin.New()
	engine.GET("/", page.Top)
	engine.POST("/", api.NewTopic)
	engine.GET("/:topic", page.Topic)
	engine.POST("/:topic", api.NewKobushi)
	engine.GET("/:topic/:kobushi", page.Kobushi)
	engine.POST("/:topic/:kobushi", api.NewMessage)
	engine.Run()
}
