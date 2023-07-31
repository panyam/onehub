package services

import (
	"time"

	"github.com/lib/pq"
)

type BaseModel struct {
	CreatedAt time.Time
	UpdatedAt time.Time
	Id        string `gorm:"primaryKey"`
	Version   int    // used for optimistic locking
}

type User struct {
	BaseModel
	Name        string
	Avatar      string
	ProfileData map[string]interface{} `gorm:"type:json"`
}

type Topic struct {
	BaseModel
	CreatorId string
	Name      string         `gorm:"index:SortedByName"`
	Users     pq.StringArray `gorm:"type:text[]"`
}

type Message struct {
	BaseModel

	ParentId    string
	TopicId     string    `gorm:"index:SortedByTopicAndCreation,priority:1"`
	CreatedAt   time.Time `gorm:"index:SortedByTopicAndCreation,priority:2"`
	SourceId    string
	UserId      string
	ContentType string
	ContentText string
	ContentData map[string]interface{} `gorm:"type:json"`
}
