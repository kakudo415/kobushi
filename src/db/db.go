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

const ringPageSize = 25
const kobushiPageSize = 25
const messagePageSize = 25

var db *gorm.DB

// NewRing from Title and Author (Desc)
func NewRing(title, author, desc string) (Rings, error) {
	r := new(Rings)
	if len(title) == 0 || len(author) == 0 {
		return *r, errors.New("EMPTY TITLE OR AUTHOR")
	}
	r.ID = kid.New(0)
	r.Title = title
	r.Author = author
	r.Desc = desc
	r.Time = time.Now()
	db.Create(r)
	return *r, nil
}

// GetRing by ID
func GetRing(id string) (Rings, error) {
	r := new(Rings)
	r.ID = kid.Parse(id)
	if r.ID.IsError() {
		return *r, errors.New("INVALID ID")
	}
	res := db.Take(r)
	if e := res.Error; e != nil {
		return *r, e
	}
	return *r, nil
}

// GetRings (max = ringsPageSize)
func GetRings(offset int) ([]Rings, error) {
	rs := new([]Rings)
	db.Order("time desc").Limit(ringPageSize).Offset(offset * ringPageSize).Find(rs)
	return *rs, nil
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

// GetKobushis (max = kobushiPageSize)
func GetKobushis(ringID string, offset int) ([]Kobushis, error) {
	ks := new([]Kobushis)
	rID := kid.Parse(ringID)
	if rID.IsError() {
		return *ks, errors.New("INVALID RING ID")
	}
	db.Order("time desc").Limit(kobushiPageSize).Offset(offset*kobushiPageSize).Find(ks, "ring_id=?", rID)
	return *ks, nil
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

// GetMessages (max = messagePageSize)
func GetMessages(kobushiID string, offset int) ([]Messages, error) {
	ms := new([]Messages)
	kID := kid.Parse(kobushiID)
	if kID.IsError() {
		return *ms, errors.New("INVALID KOBUSHI ID")
	}
	db.Order("time desc").Limit(messagePageSize).Offset(offset*messagePageSize).Find(ms, "kobushi_id=?", kID)
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
	db.AutoMigrate(new(Rings))
	db.AutoMigrate(new(Kobushis)).AddForeignKey("ring_id", "rings(id)", "CASCADE", "CASCADE")
	db.AutoMigrate(new(Messages)).AddForeignKey("kobushi_id", "kobushis(id)", "CASCADE", "CASCADE")
}
