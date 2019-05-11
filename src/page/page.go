package page

import (
	"github.com/gin-gonic/gin"
)

// Top page
func Top(c *gin.Context) {
	c.HTML(200, "top.html", gin.H{})
}

// Topic page
func Topic(c *gin.Context) {

}

// Kobushi page
func Kobushi(c *gin.Context) {

}
