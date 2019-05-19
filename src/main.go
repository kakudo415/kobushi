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
	engine.POST("/", api.NewRing)
	engine.GET("/ring/:ring_id", page.Ring)
	engine.POST("/ring/:ring_id", api.NewKobushi)
	engine.GET("/ring/:ring_id/:kobushi_id", page.Kobushi)
	engine.POST("/ring/:ring_id/:kobushi_id", api.NewMessage)
	engine.StaticFS("/static/", http.Dir("views/static/"))
	engine.Run()
}
