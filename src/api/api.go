package api

import (
	"github.com/gin-gonic/gin"

	"../db"
)

// NewTopicJSON - POST API
type NewTopicJSON struct {
	Title  string
	Author string
	Desc   string // Description
}

// NewTopic API
func NewTopic(c *gin.Context) {
	j := new(NewTopicJSON)
	c.BindJSON(j)
	db.NewTopic(j.Title, j.Author, j.Desc)
}

// NewKobushi API
func NewKobushi(c *gin.Context) {

}

// NewMessage API
func NewMessage(c *gin.Context) {

}
