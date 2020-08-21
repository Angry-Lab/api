package segment

import (
	"context"
	"github.com/Angry-Lab/api/pkg/entity"
)

type Repository interface {
	GetByID(ctx context.Context, id int) (*entity.Segment, error)
	PutIfExits(ctx context.Context, segment *entity.Segment) error
	Create(ctx context.Context, segment *entity.Segment) error
}
