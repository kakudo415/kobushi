package page

import (
	"os"
	"strconv"

	"github.com/gin-gonic/gin"

	"../db"
)

type ring struct {
	Path  string
	Title string
	Desc  string
}

type kobushi struct {
	Path  string
	Title string
	Desc  string
}

var baseURL = "http://localhost:8080/"

// Top page
func Top(c *gin.Context) {
	offset := parsePage(c)
	rs, e := db.GetRings(offset)
	if e != nil {
		c.Status(404)
		return
	}
	var frs []ring
	for _, v := range rs {
		frs = append(frs, ring{
			Path:  "ring/" + v.ID.ToDec(),
			Title: v.Title,
			Desc:  v.Desc,
		})
	}
	c.HTML(200, "top.html", gin.H{"BaseURL": baseURL, "Rings": frs})
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
		fks = append(fks, kobushi{
			Path:  "ring/" + v.RingID.ToDec() + "/" + v.ID.ToDec(),
			Title: v.Title,
			Desc:  v.Desc,
		})
	}
	c.HTML(200, "ring.html", gin.H{"BaseURL": baseURL, "Ring": r, "Kobushis": fks})
}

// Kobushi page
func Kobushi(c *gin.Context) {
	r, e := db.GetRing(c.Param("ring_id"))
	if e != nil {
		c.Status(404)
		return
	}
	k, e := db.GetKobushi(c.Param("kobushi_id"))
	if e != nil {
		c.Status(404)
		return
	}
	offset := parsePage(c)
	ms, e := db.GetMessages(c.Param("kobushi_id"), offset)
	if e != nil {
		c.Status(404)
		return
	}
	c.HTML(200, "kobushi.html", gin.H{"BaseURL": baseURL, "Ring": r, "Kobushi": k, "Messages": ms})
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

func init() {
	if os.Getenv("GIN_MODE") == "release" {
		baseURL = "https://kakudo.app/kobushi/"
	}
}
