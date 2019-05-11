package db

import (
	"os"

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

// NewTopic from Title and Author
func NewTopic(title, author, desc string) {
	t := new(Topics)
	t.ID = xid.New()
	t.Title = title
	t.Author = author
	t.Desc = desc
	db.FirstOrCreate(t)
}

func conn() {
	var err error
	db, err = gorm.Open("mysql", os.Getenv("KOBUSHI_DB"))
	if err != nil {
		panic(err)
	}
}

func init() {
	conn()
	db.AutoMigrate(new(Topics))
	db.AutoMigrate(new(Messages))
}
