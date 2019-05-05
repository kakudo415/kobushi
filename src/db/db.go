package db

import (
	"github.com/jinzhu/gorm"
	"github.com/rs/xid"

	// MySQL
	_ "github.com/go-sql-driver/mysql"
)

// Topics model
type Topics struct {
	ID     xid.ID
	Title  string
	Author string
	Desc   string // Description
}

// Messages model
type Messages struct {
	ID      xid.ID
	TopicID xid.ID
	Content string
}

var db *gorm.DB

func conn() {
	var err error
	db, err = gorm.Open("mysql", "")
	if err != nil {
		panic(err)
	}
}

func init() {
	conn()
}
