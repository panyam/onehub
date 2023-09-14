package datastore

import (
	"errors"
	"strings"
	"time"

	"gorm.io/gorm"
)

func (tdb *OneHubDB) GetMessages(topic_id string, user_id string, pageKey string, pageSize int) (out []*Message, err error) {
	user_id = strings.Trim(user_id, " ")
	topic_id = strings.Trim(topic_id, " ")
	if user_id == "" && topic_id == "" {
		return nil, errors.New("Either topic_id or user_id or both must be provided")
	}
	query := tdb.storage
	if topic_id != "" {
		query = query.Where("topic_id = ?", topic_id)
	}
	if user_id != "" {
		query = query.Where("user_id = ?", user_id)
	}
	if pageKey != "" {
		offset := 0
		query = query.Offset(offset)
	}
	if pageSize <= 0 || pageSize > 10000 {
		pageSize = 10000
	}
	query = query.Limit(pageSize)

	err = query.Find(&out).Error
	return out, err
}

// Get messages in a topic paginated and ordered by creation time stamp
func (tdb *OneHubDB) ListMessagesInTopic(topic_id string, pageKey string, pageSize int) (out []*Topic, err error) {
	err = tdb.storage.Where("topic_id= ?", topic_id).Find(&out).Error
	return
}

func (tdb *OneHubDB) GetMessage(msgid string) (*Message, error) {
	var out Message
	err := tdb.storage.First(&out, "id = ?", msgid).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		} else {
			return nil, err
		}
	}
	return &out, err
}

func (tdb *OneHubDB) ListMessages(topic_id string, pageKey string, pageSize int) (out []*Message, err error) {
	query := tdb.storage.Where("topic_id = ?").Order("created_at asc")
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

func (tdb *OneHubDB) CreateMessages(msgs []*Message) (err error) {
	for _, msg := range msgs {
		msg.CreatedAt = time.Now()
		msg.UpdatedAt = time.Now()
	}
	result := tdb.storage.Model(&Message{}).Create(msgs)
	err = result.Error
	return
}

func (tdb *OneHubDB) CreateMessage(msg *Message) (err error) {
	msg.CreatedAt = time.Now()
	msg.UpdatedAt = time.Now()
	result := tdb.storage.Model(&Message{}).Create(msg)
	err = result.Error
	return
}

func (tdb *OneHubDB) DeleteMessage(msgId string) (err error) {
	err = tdb.storage.Where("id = ?", msgId).Delete(&Message{}).Error
	return
}

func (tdb *OneHubDB) SaveMessage(msg *Message) (err error) {
	db := tdb.storage
	q := db.Model(msg).Where("id = ? and version = ?", msg.Id, msg.Version)
	msg.UpdatedAt = time.Now()
	result := q.UpdateColumns(map[string]interface{}{
		"updated_at":   msg.UpdatedAt,
		"content_type": msg.ContentType,
		"content_text": msg.ContentText,
		"content_data": msg.ContentData,
		"user_id":      msg.SourceId,
		"source_id":    msg.SourceId,
		"parent_id":    msg.ParentId,
		"version":      msg.Version + 1,
	})

	err = result.Error
	if err == nil && result.RowsAffected == 0 {
		// Must have failed due to versioning
		err = MessageUpdateFailed
	}
	return
}
