package mockup

import (
	"context"
	"github.com/Angry-Lab/api/pkg/entity"
	"github.com/Angry-Lab/api/pkg/token"
	"github.com/pkg/errors"
	"sync"
)

type tokenMockup struct {
	store map[string]*entity.Token
	mu    *sync.RWMutex
}

func Tokens() token.Repository {
	return &tokenMockup{
		store: make(map[string]*entity.Token),
		mu:    &sync.RWMutex{},
	}
}

func (repo *tokenMockup) Get(ctx context.Context, value string) (*entity.Token, error) {
	repo.mu.Lock()
	defer repo.mu.Unlock()

	t, ok := repo.store[value]
	if !ok {
		return nil, errors.New("token not found")
	}

	return t, nil
}

func (repo *tokenMockup) Set(ctx context.Context, t *entity.Token) error {
	repo.mu.Lock()
	defer repo.mu.Unlock()

	repo.store[t.Value] = t

	return nil
}
