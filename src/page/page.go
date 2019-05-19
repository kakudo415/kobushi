package page

import (
	"strconv"

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
	page := c.Query("p")
	offset, e := strconv.ParseUint(page, 10, 64)
	if len(page) != 0 && e != nil {
		c.Status(404)
		return
	}
	offset-- // 0,1,2... => 1,2,3...
	m, e := db.GetMessages(c.Param("kobushi_id"), uint(offset))
	if e != nil {
		c.Status(404)
		return
	}
	c.HTML(200, "kobushi.html", gin.H{"Messages": m})
}
