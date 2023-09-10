package dbsync

import "time"

const (
	SYNC_SESSION_CREATED   = 0
	SYNC_SESSION_STARTED   = 1
	SYNC_SESSION_PAUSED    = 2
	SYNC_SESSION_DESTROYED = 3
)

const (
	EVENT_TYPE_WATERMARK = 0
	EVENT_TYPE_CREATE    = 1
	EVENT_TYPE_UPDATE    = 2
	EVENT_TYPE_DELETE    = 3
)

type WaterMark = string

type Row interface {
	Key() interface{}
	Data() interface{}
}

type DBLogEvent struct {
	Type      int
	LSN       string
	Timestamp time.Time
	Source    string
	Key       interface{}
	Data      interface{}
}

type DBItem struct {
	Key  interface{}
	Data interface{}
}

func (e *DBLogEvent) IsWatermark() bool {
	return e.Type == EVENT_TYPE_WATERMARK
}
