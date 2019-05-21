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
	Title  string `sql:"type:char(200)"`
	Author string `sql:"type:char(200)"`
	Desc   string `sql:"type:varchar(1000)"` // Description
	Time   time.Time
}

// Kobushis model
type Kobushis struct {
	ID     kid.ID `sql:"type:bigint unsigned" gorm:"primary_key"`
	RingID kid.ID `sql:"type:bigint unsigned"`
	Title  string `sql:"type:char(200)"`
	Desc   string `sql:"type:varchar(1000)"` // Description
	Time   time.Time
}

// Messages model
type Messages struct {
	ID        kid.ID `sql:"type:bigint unsigned" gorm:"primary_key"`
	KobushiID kid.ID `sql:"type:bigint unsigned"`
	Body      string `sql:"type:varchar(2000)"`
	Time      time.Time
}

const ringPageSize = 25
const kobushiPageSize = 25
const messagePageSize = 25

var db *gorm.DB

// NewRing from Title and Author (Desc)
func NewRing(title, author, desc string) (Rings, error) {
	r := new(Rings)
	if len(title) == 0 {
		return *r, errors.New("EMPTY TITLE")
	}
	r.ID = kid.New(0)
	r.Title = title
	r.Author = author
	r.Desc = desc
	r.Time = time.Now()
	res := db.Create(r)
	if e := res.Error; e != nil {
		return *r, e
	}
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
	res := db.Order("time desc").Limit(ringPageSize).Offset(offset * ringPageSize).Find(rs)
	if e := res.Error; e != nil {
		return *rs, e
	}
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
	res := db.Create(k)
	if e := res.Error; e != nil {
		return *k, e
	}
	return *k, nil
}

// GetKobushi by ID
func GetKobushi(id string) (Kobushis, error) {
	k := new(Kobushis)
	k.ID = kid.Parse(id)
	if k.ID.IsError() {
		return *k, errors.New("INVALID ID")
	}
	res := db.Take(k)
	if e := res.Error; e != nil {
		return *k, e
	}
	return *k, nil
}

// GetKobushis (max = kobushiPageSize)
func GetKobushis(ringID string, offset int) ([]Kobushis, error) {
	ks := new([]Kobushis)
	rID := kid.Parse(ringID)
	if rID.IsError() {
		return *ks, errors.New("INVALID RING ID")
	}
	res := db.Order("time desc").Limit(kobushiPageSize).Offset(offset*kobushiPageSize).Find(ks, "ring_id=?", rID)
	if e := res.Error; e != nil {
		return *ks, e
	}
	return *ks, nil
}

// NewMessage from KobushiID and Body text
func NewMessage(kobushiID, body string) (Messages, error) {
	m := new(Messages)
	if len(body) == 0 {
		return *m, errors.New("EMPTY BODY")
	}
	m.ID = kid.New(0)
	m.KobushiID = kid.Parse(kobushiID)
	if m.KobushiID.IsError() {
		return *m, errors.New("INVALID KOBUSHI ID")
	}
	m.Body = body
	m.Time = time.Now()
	res := db.Create(m)
	if e := res.Error; e != nil {
		return *m, e
	}
	return *m, nil
}

// GetMessages (max = messagePageSize)
func GetMessages(kobushiID string, offset int) ([]Messages, error) {
	ms := new([]Messages)
	kID := kid.Parse(kobushiID)
	if kID.IsError() {
		return *ms, errors.New("INVALID KOBUSHI ID")
	}
	res := db.Order("time desc").Limit(messagePageSize).Offset(offset*messagePageSize).Find(ms, "kobushi_id=?", kID)
	if e := res.Error; e != nil {
		return *ms, e
	}
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
