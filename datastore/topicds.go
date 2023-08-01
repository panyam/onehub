package datastore

import (
	"errors"
	"log"
	"strings"
	"time"

	"gorm.io/gorm"
)

/////////////////////// Topic DB

func (tdb *OneHubDB) SaveTopic(topic *Topic) (err error) {
	db := tdb.storage
	topic.UpdatedAt = time.Now()
	if strings.Trim(topic.Id, " ") == "" {
		return InvalidIDError
		// create a new one
	}

	result := db.Save(topic)
	err = result.Error
	if err == nil && result.RowsAffected == 0 {
		topic.CreatedAt = time.Now()
		err = tdb.storage.Create(topic).Error
	}
	return
}

func (tdb *OneHubDB) DeleteTopic(topicId string) (err error) {
	err = tdb.storage.Where("topic_id = ?", topicId).Delete(&Message{}).Error
	if err == nil {
		err = tdb.storage.Where("id = ?", topicId).Delete(&Topic{}).Error
	}
	return
}

func (tdb *OneHubDB) GetTopic(id string) (*Topic, error) {
	var out Topic
	err := tdb.storage.First(&out, "id = ?", id).Error
	if err != nil {
		log.Println("GetTopic Error: ", id, err)
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		} else {
			return nil, err
		}
	}

	return &out, err
}

func (tdb *OneHubDB) ListTopics(pageKey string, pageSize int) (out []*Topic, err error) {
	query := tdb.storage.Model(&Topic{}).Order("name asc")
	if pageKey != "" {
		count := 0
		query = query.Offset(count)
	}
	if pageSize <= 0 || pageSize > tdb.MaxPageSize {
		pageSize = tdb.MaxPageSize
	}
	query = query.Limit(pageSize)
	err = query.Find(&out).Error
	return out, err
}
