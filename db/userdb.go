package db

import (
	"errors"
	"log"
	"strings"
	"time"

	"gorm.io/gorm"
)

func (tdb *OneHubDB) SaveUser(topic *User) (err error) {
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

func (tdb *OneHubDB) DeleteUser(topicId string) (err error) {
	err = tdb.storage.Where("topic_id = ?", topicId).Delete(&Message{}).Error
	if err == nil {
		err = tdb.storage.Where("id = ?", topicId).Delete(&User{}).Error
	}
	return
}

func (tdb *OneHubDB) GetUser(id string) (*User, error) {
	var out User
	err := tdb.storage.First(&out, "id = ?", id).Error
	if err != nil {
		log.Println("GetUser Error: ", id, err)
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		} else {
			return nil, err
		}
	}

	return &out, err
}

func (tdb *OneHubDB) ListUsers(pageKey string, pageSize int) (out []*User, err error) {
	query := tdb.storage.Model(&User{}).Order("name asc")
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
