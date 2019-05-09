package main

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"./api"
	"./page"
)

func main() {
	engine := gin.New()
	engine.LoadHTMLGlob("views/*.html")
	engine.GET("/", page.Top)
	engine.POST("/", api.NewTopic)
	engine.GET("/punch/:topic", page.Topic)
	engine.POST("/punch/:topic", api.NewKobushi)
	engine.GET("/punch/:topic/:kobushi", page.Kobushi)
	engine.POST("/punch/:topic/:kobushi", api.NewMessage)
	engine.StaticFS("/static/", http.Dir("views/static/"))
	engine.Run()
}
