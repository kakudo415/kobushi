package db

import (
	"errors"
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

// Kobushis model
type Kobushis struct {
	ID     string
	RingID string
	Title  string
	Desc   string // Description
}

// Messages model
type Messages struct {
	ID        string
	KobushiID string
	Content   string
}

var db *gorm.DB

// NewRing from Title and Author (Desc)
func NewRing(title, author, desc string) (Rings, error) {
	t := new(Rings)
	if len(title) == 0 || len(author) == 0 {
		return *t, errors.New("EMPTY TITLE OR AUTHOR")
	}
	t.ID = xid.New().String()
	t.Title = title
	t.Author = author
	t.Desc = desc
	db.FirstOrCreate(t)
	return *t, nil
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

// NewKobushi from Title and RingID (Desc)
func NewKobushi(title, ringID, desc string) (Kobushis, error) {
	k := new(Kobushis)
	if len(title) == 0 {
		return *k, errors.New("EMPTY TITLE")
	}
	k.ID = xid.New().String()
	k.RingID = ringID
	k.Title = title
	k.Desc = desc
	db.FirstOrCreate(k)
	return *k, nil
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
	db.AutoMigrate(new(Kobushis)).AddForeignKey("ring_id", "rings(id)", "CASCADE", "CASCADE")
	db.AutoMigrate(new(Messages)).AddForeignKey("kobushi_id", "kobushis(id)", "CASCADE", "CASCADE")
}
