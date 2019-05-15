package db

import (
	"os"

	"github.com/jinzhu/gorm"
	"github.com/rs/xid"

	// MySQL
	_ "github.com/go-sql-driver/mysql"
)

// Rings model
type Rings struct {
	ID     string
	Title  string
	Author string
	Desc   string // Description
}

// Messages model
type Messages struct {
	ID      string
	RingID  string
	Content string
}

var db *gorm.DB

// NewRing from Title and Author
func NewRing(title, author, desc string) string {
	t := new(Rings)
	t.ID = xid.New().String()
	t.Title = title
	t.Author = author
	t.Desc = desc
	db.FirstOrCreate(t)
	return t.ID
}

// GetRing by ID
func GetRing(id string) (Rings, error) {
	t := new(Rings)
	t.ID = id
	res := db.Take(t)
	if e := res.Error; e != nil {
		return *t, e
	}
	return *t, nil
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
	db.LogMode(true)
	db.AutoMigrate(new(Rings))
	db.AutoMigrate(new(Messages))
}
