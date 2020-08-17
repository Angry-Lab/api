package mockup

import (
	"context"
	"github.com/Angry-Lab/api/pkg/entity"
	"github.com/Angry-Lab/api/pkg/user"
	"github.com/pkg/errors"
	"sync"
)

type userMockup struct {
	store map[string]*entity.User
	mu    *sync.RWMutex
}

func Users() user.Repository {
	return &userMockup{
		store: make(map[string]*entity.User),
		mu:    &sync.RWMutex{},
	}
}

func (repo *userMockup) GetOneByLogin(ctx context.Context, login string) (*entity.User, error) {
	repo.mu.Lock()
	defer repo.mu.Unlock()

	u, ok := repo.store[login]
	if !ok {
		return nil, errors.New("user not found")
	}

	return u, nil
}

func (repo *userMockup) Update(ctx context.Context, u *entity.User) error {
	repo.mu.Lock()
	defer repo.mu.Unlock()

	repo.store[u.Login] = u

	return nil
}
