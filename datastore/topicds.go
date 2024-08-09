package datastore

import (
	"context"
	"errors"
	"log"
	"strings"
	"time"

	"gorm.io/gorm"
)

/////////////////////// Topic DB

func (tdb *OneHubDB) SaveTopic(ctx context.Context, topic *Topic) (err error) {
	_, span := Tracer.Start(ctx, "db.SaveTopic")
	defer span.End()
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

func (tdb *OneHubDB) DeleteTopic(ctx context.Context, topicId string) (err error) {
	_, span := Tracer.Start(ctx, "db.DeleteTopic")
	defer span.End()
	err = tdb.storage.Where("topic_id = ?", topicId).Delete(&Message{}).Error
	if err == nil {
		err = tdb.storage.Where("id = ?", topicId).Delete(&Topic{}).Error
	}
	return
}

func (tdb *OneHubDB) GetTopic(ctx context.Context, id string) (*Topic, error) {
	_, span := Tracer.Start(ctx, "db.GetTopic")
	defer span.End()
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

func (tdb *OneHubDB) ListTopics(ctx context.Context, pageKey string, pageSize int) (out []*Topic, err error) {
	_, span := Tracer.Start(ctx, "db.ListTopics")
	defer span.End()
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
