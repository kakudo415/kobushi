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

// NewMessageJSON - POST API
type NewMessageJSON struct {
	Body string `json:"body"`
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
	c.JSON(200, NewRingIDJSON{RingID: r.ID.ToDec()})
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
	c.JSON(200, NewKobushiIDJSON{KobushiID: k.ID.ToDec()})
}

// NewMessage API
func NewMessage(c *gin.Context) {
	j := new(NewMessageJSON)
	c.BindJSON(j)
	_, e := db.NewMessage(c.Param("kobushi_id"), j.Body)
	if e != nil {
		c.Status(404)
		return
	}
	c.Status(200)
}
