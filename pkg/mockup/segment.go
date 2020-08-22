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

func (mockup segmentMockup) GetByID(ctx context.Context, id int) (*segment.Segment, error) {
	mockup.mu.Lock()
	defer mockup.mu.Unlock()
	s, ok := mockup.store[id]
	if !ok {
		return nil, errors.New("segment not found")
	} else {
		return s, nil
	}
}

func (mockup segmentMockup) PutIfExits(ctx context.Context, segment *segment.Segment) error {
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

func (mockup segmentMockup) Create(ctx context.Context, segment *segment.Segment) error {
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

func (mockup segmentMockup) Delete(ctx context.Context, id int) error {
	mockup.mu.Lock()
	defer mockup.mu.Unlock()
	_, ok := mockup.store[id]
	if !ok {
		mockup.store[id] = nil
		return nil
	} else {
		return errors.New("segment not found")
	}
}

func (mockup segmentMockup) GetAll(ctx context.Context) ([]*segment.Segment, error) {
	a := []*segment.Segment{}
	for _, v := range mockup.store {
		a = append(a, v)
	}
	return a, nil
}
