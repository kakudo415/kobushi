package page

import (
	"github.com/gin-gonic/gin"

	"../db"
)

// Top page
func Top(c *gin.Context) {
	c.HTML(200, "top.html", gin.H{})
}

// Ring page
func Ring(c *gin.Context) {
	r, e := db.GetRing(c.Param("ring_id"))
	if e != nil {
		c.Status(404)
		return
	}
	c.HTML(200, "ring.html", gin.H{"aiueo": r.Title + " " + r.Author})
}

// Kobushi page
func Kobushi(c *gin.Context) {
	// k, e := db.GetKobushi(c.Param("ring_id"), c.Param("kobushi_id"))
}
