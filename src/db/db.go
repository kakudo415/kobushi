package db

import (
	"errors"
	"os"
	"time"

	"github.com/jinzhu/gorm"
	"github.com/kakudo415/kid"

	// MySQL
	_ "github.com/go-sql-driver/mysql"
)

// Rings model
type Rings struct {
	ID     kid.ID `sql:"type:bigint unsigned" gorm:"primary_key"`
	Title  string
	Author string
	Desc   string // Description
	Time   time.Time
}

// Kobushis model
type Kobushis struct {
	ID     kid.ID `sql:"type:bigint unsigned" gorm:"primary_key"`
	RingID kid.ID `sql:"type:bigint unsigned"`
	Title  string
	Desc   string // Description
	Time   time.Time
}

// Messages model
type Messages struct {
	ID        kid.ID `sql:"type:bigint unsigned" gorm:"primary_key"`
	KobushiID kid.ID `sql:"type:bigint unsigned"`
	Body      string
	Time      time.Time
}

const pageSize = 25 // Messages page size

var db *gorm.DB

// NewRing from Title and Author (Desc)
func NewRing(title, author, desc string) (Rings, error) {
	t := new(Rings)
	if len(title) == 0 || len(author) == 0 {
		return *t, errors.New("EMPTY TITLE OR AUTHOR")
	}
	t.ID = kid.New(0)
	t.Title = title
	t.Author = author
	t.Desc = desc
	t.Time = time.Now()
	db.Create(t)
	return *t, nil
}

// GetRing by ID
func GetRing(id string) (Rings, error) {
	t := new(Rings)
	t.ID = kid.Parse(id)
	if t.ID.IsError() {
		return *t, errors.New("INVALID ID")
	}
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
	k.ID = kid.New(0)
	k.RingID = kid.Parse(ringID)
	if k.RingID.IsError() {
		return *k, errors.New("INVALID RING ID")
	}
	k.Title = title
	k.Desc = desc
	k.Time = time.Now()
	db.Create(k)
	return *k, nil
}

// NewMessage from KobushiID and Body text
func NewMessage(kobushiID, body string) (Messages, error) {
	m := new(Messages)
	m.ID = kid.New(0)
	m.KobushiID = kid.Parse(kobushiID)
	if m.KobushiID.IsError() {
		return *m, errors.New("INVALID KOBUSHI ID")
	}
	m.Body = body
	m.Time = time.Now()
	db.Create(m)
	return *m, nil
}

// GetMessages (max = pageSize)
func GetMessages(kobushiID string, offset uint) ([]Messages, error) {
	ms := new([]Messages)
	kID := kid.Parse(kobushiID)
	if kID == 0 {
		return *ms, errors.New("INVALID ID")
	}
	db.Order("time desc").Limit(pageSize).Offset(offset*pageSize).Find(ms, "kobushi_id=?", kID)
	return *ms, nil
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
