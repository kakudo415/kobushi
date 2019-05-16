package api

import (
	"github.com/gin-gonic/gin"

	"../db"
)

// NewRingJSON - POST API
type NewRingJSON struct {
	Title  string `json:"title"`
	Author string `json:"author"`
	Desc   string `json:"description"`
}

// NewRingIDJSON - POST API Response
type NewRingIDJSON struct {
	RingID string `json:"ring_id"`
}

// NewKobushiJSON - POST API
type NewKobushiJSON struct {
	Title string `json:"title"`
	Desc  string `json:"description"`
}

// NewKobushiIDJSON - POST API Response
type NewKobushiIDJSON struct {
	KobushiID string `json:"kobushi_id"`
}

// NewRing API
func NewRing(c *gin.Context) {
	j := new(NewRingJSON)
	c.BindJSON(j)
	r, e := db.NewRing(j.Title, j.Author, j.Desc)
	if e != nil {
		c.Status(404)
		return
	}
	c.JSON(200, NewRingIDJSON{RingID: r.ID})
}

// NewKobushi API
func NewKobushi(c *gin.Context) {
	j := new(NewKobushiJSON)
	c.BindJSON(j)
	k, e := db.NewKobushi(j.Title, c.Param("ring_id"), j.Desc)
	if e != nil {
		c.Status(404)
		return
	}
	c.JSON(200, NewKobushiIDJSON{KobushiID: k.ID})
}

// NewMessage API
func NewMessage(c *gin.Context) {

}
