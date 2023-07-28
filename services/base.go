package services

import (
	"fmt"
	"log"
	"sort"
	"time"

	tspb "google.golang.org/protobuf/types/known/timestamppb"
)

type EntityStore[T any] struct {
	IDCount  int
	Entities map[string]*T

	// Getters/Setters for ID
	IDSetter func(entity *T, id string)
	IDGetter func(entity *T) string

	// Getters/Setters for created timestamp
	CreatedAtSetter func(entity *T, ts *tspb.Timestamp)
	CreatedAtGetter func(entity *T) *tspb.Timestamp

	// Getters/Setters for udpated timestamp
	UpdatedAtSetter func(entity *T, ts *tspb.Timestamp)
	UpdatedAtGetter func(entity *T) *tspb.Timestamp
}

func NewEntityStore[T any]() *EntityStore[T] {
	return &EntityStore[T]{
		Entities: make(map[string]*T),
	}
}

func (s *EntityStore[T]) Create(entity *T) *T {
	s.IDCount++
	newid := fmt.Sprintf("%d", s.IDCount)
	s.Entities[newid] = entity
	s.IDSetter(entity, newid)
	s.CreatedAtSetter(entity, tspb.New(time.Now()))
	s.UpdatedAtSetter(entity, tspb.New(time.Now()))
	return entity
}

func (s *EntityStore[T]) Get(id string) *T {
	if entity, ok := s.Entities[id]; ok {
		return entity
	}
	return nil
}

func (s *EntityStore[T]) BatchGet(ids []string) map[string]*T {
	out := make(map[string]*T)
	for _, id := range ids {
		if entity, ok := s.Entities[id]; ok {
			out[id] = entity
		}
	}
	return out
}

// Updates specific fields of an Entity
func (s *EntityStore[T]) Update(entity *T) *T {
	s.UpdatedAtSetter(entity, tspb.New(time.Now()))
	return entity
}

// Deletes an entity from our system.
func (s *EntityStore[T]) Delete(id string) bool {
	_, ok := s.Entities[id]
	if ok {
		delete(s.Entities, id)
	}
	return ok
}

// Finds and retrieves entity matching the particular criteria.
func (s *EntityStore[T]) List(ltfunc func(t1, t2 *T) bool, filterfunc func(t *T) bool) (out []*T) {
	log.Println("E: ", s.Entities)
	for _, ent := range s.Entities {
		if filterfunc == nil || filterfunc(ent) {
			out = append(out, ent)
		}
	}
	// Sort in reverse order of name
	sort.Slice(out, func(idx1, idx2 int) bool {
		ent1 := out[idx1]
		ent2 := out[idx2]
		return ltfunc(ent1, ent2)
	})
	return
}
