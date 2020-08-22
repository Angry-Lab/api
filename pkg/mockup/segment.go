package mockup

import (
	"context"
	"errors"
	"github.com/Angry-Lab/api/pkg/segment"
	"sync"
)

type segmentMockup struct {
	store map[int]*segment.Segment
	mu    *sync.RWMutex
}

func Segments() segment.Repository {
	return &segmentMockup{
		store: make(map[int]*segment.Segment),
		mu:    &sync.RWMutex{},
	}
}

func (mock *segmentMockup) List(ctx context.Context) ([]*segment.Segment, error) {
	return nil, nil
}

func (mock *segmentMockup) GetByID(ctx context.Context, id int) (*segment.Segment, error) {
	mock.mu.Lock()
	defer mock.mu.Unlock()
	s, ok := mock.store[id]
	if !ok {
		return nil, errors.New("segment not found")
	} else {
		return s, nil
	}
}

func (mock *segmentMockup) PutIfExits(ctx context.Context, segment *segment.Segment) error {
	mock.mu.Lock()
	defer mock.mu.Unlock()
	var id = segment.ID
	_, ok := mock.store[id]
	if ok {
		mock.store[id] = segment
		return nil
	} else {
		return errors.New("segment not found")
	}
}

func (mock *segmentMockup) Create(ctx context.Context, segment *segment.Segment) error {
	mock.mu.Lock()
	defer mock.mu.Unlock()
	var id = segment.ID
	_, ok := mock.store[id]
	if !ok {
		mock.store[id] = segment
		return nil
	} else {
		return errors.New("unable to add new segment. segment found")
	}
}
