package token

import (
	"context"
	"github.com/Angry-Lab/api/pkg/entity"
)

type Repository interface {
	Get(ctx context.Context, value string) (*entity.Token, error)
	Set(ctx context.Context, token *entity.Token) error
}
