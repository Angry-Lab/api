package mockup

import (
	"context"
	"errors"
	"github.com/Angry-Lab/api/pkg/entity"
	"github.com/Angry-Lab/api/pkg/segment"
	"sync"
)

type segmentMockup struct {
	store map[int]*entity.Segment
	mu    *sync.RWMutex
}

func Segments() segment.Repository {
	return &segmentMockup{
		store: make(map[int]*entity.Segment),
		mu:    &sync.RWMutex{},
	}
}

func (mockup segmentMockup) GetByID(ctx context.Context, id int) (*entity.Segment, error) {
	mockup.mu.Lock()
	defer mockup.mu.Unlock()
	s, ok := mockup.store[id]
	if !ok {
		return nil, errors.New("segment not found")
	} else {
		return s, nil
	}
}

func (mockup segmentMockup) PutIfExits(ctx context.Context, segment *entity.Segment) error {
	mockup.mu.Lock()
	defer mockup.mu.Unlock()
	var id = segment.ID
	_, ok := mockup.store[id]
	if ok {
		mockup.store[id] = segment
		return nil
	} else {
		return errors.New("segment not found")
	}
}

func (mockup segmentMockup) Create(ctx context.Context, segment *entity.Segment) error {
	mockup.mu.Lock()
	defer mockup.mu.Unlock()
	var id = segment.ID
	_, ok := mockup.store[id]
	if !ok {
		mockup.store[id] = segment
		return nil
	} else {
		return errors.New("unable to add new segment. segment found")
	}
}
