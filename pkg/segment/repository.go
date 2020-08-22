package segment

import (
	"context"
)

type Repository interface {
	GetByID(ctx context.Context, id int) (*Segment, error)
	PutIfExits(ctx context.Context, segment *Segment) error
	Create(ctx context.Context, segment *Segment) error
	Delete(ctx context.Context, id int) error
	GetAll(ctx context.Context) ([]*Segment, error)
}
