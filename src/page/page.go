package page

import (
	"strconv"

	"github.com/gin-gonic/gin"

	"../db"
)

type kobushi struct {
	Path  string
	Title string
	Desc  string
}

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
	offset := parsePage(c)
	ks, e := db.GetKobushis(c.Param("ring_id"), offset)
	if e != nil {
		c.Status(404)
		return
	}
	var fks []kobushi
	for _, v := range ks {
		fk := kobushi{
			Path:  "./" + v.RingID.ToDec() + "/" + v.ID.ToDec(),
			Title: v.Title,
			Desc:  v.Desc,
		}
		fks = append(fks, fk)
	}
	c.HTML(200, "ring.html", gin.H{"Kobushi": r.Title + " " + r.Author, "Kobushis": fks})
}

// Kobushi page
func Kobushi(c *gin.Context) {
	offset := parsePage(c)
	m, e := db.GetMessages(c.Param("kobushi_id"), offset)
	if e != nil {
		c.Status(404)
		return
	}
	c.HTML(200, "kobushi.html", gin.H{"Messages": m})
}

func parsePage(c *gin.Context) int {
	page := c.Query("p")
	offset, e := strconv.ParseUint(page, 10, 64)
	if len(page) != 0 && e != nil {
		return 0
	}
	offset-- // 0,1,2... => 1,2,3...
	return int(offset)
}
