package datastore

import (
	"errors"
	"log"
	"math"
	"math/rand"
	"strconv"
	"time"

	"gorm.io/gorm"
)

type OneHubDB struct {
	storage     *gorm.DB
	MaxPageSize int
}

type GenId struct {
	Class     string `gorm:"primaryKey"`
	Id        string `gorm:"primaryKey"`
	CreatedAt time.Time
}

var InvalidIDError = errors.New("ID is invalid or empty")
var MessageUpdateFailed = errors.New("Update failed concurrency check")
var TopicUpdateFailed = errors.New("Update failed concurrency check")
var UserUpdateFailed = errors.New("Update failed concurrency check")

func NewOneHubDB(gormdb *gorm.DB) *OneHubDB {
	gormdb.AutoMigrate(&GenId{})
	gormdb.AutoMigrate(&User{})
	gormdb.AutoMigrate(&Topic{})
	gormdb.AutoMigrate(&Message{})
	return &OneHubDB{
		storage:     gormdb,
		MaxPageSize: 1000,
	}
}

func randid() string {
	max_id := int64(math.Pow(36, 8))
	randval := rand.Int63() % max_id
	return strconv.FormatInt(randval, 36)
}

// Generate 1 New ID
func (tdb *OneHubDB) NewID(cls string) string {
	for {
		gid := GenId{Id: randid(), Class: cls, CreatedAt: time.Now()}
		err := tdb.storage.Create(gid).Error
		if err == nil {
			return gid.Id
		} else {
			log.Println("ID Create Error: ", err)
		}
	}
}

/**
 * Create N IDs in batch.
 */
func (tdb *OneHubDB) NewIDs(cls string, numids int) (out []string) {
	for i := 0; i < numids; i++ {
		for {
			gid := GenId{Id: randid(), Class: cls, CreatedAt: time.Now()}
			err := tdb.storage.Create(gid).Error
			if err != nil {
				log.Println("ID Create Error: ", i, err)
			} else {
				out = append(out, gid.Id)
				break
			}
		}
	}
	return
}
