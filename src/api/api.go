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

// NewRing API
func NewRing(c *gin.Context) {
	j := new(NewRingJSON)
	c.BindJSON(j)
	id := db.NewRing(j.Title, j.Author, j.Desc)
	c.JSON(200, NewRingIDJSON{RingID: id})
}

// NewKobushi API
func NewKobushi(c *gin.Context) {

}

// NewMessage API
func NewMessage(c *gin.Context) {

}
