package user

import (
	"context"
	"github.com/Angry-Lab/api/pkg/entity"
)

type Repository interface {
	GetOneByLogin(ctx context.Context, login string) (*entity.User, error)
	Update(ctx context.Context, user *entity.User) error
}
