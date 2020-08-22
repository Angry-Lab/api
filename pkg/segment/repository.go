package segment

import (
	"context"
)

type Repository interface {
	List(ctx context.Context) ([]*Segment, error)
	GetByID(ctx context.Context, id int) (*Segment, error)
	PutIfExits(ctx context.Context, segment *Segment) error
	Create(ctx context.Context, segment *Segment) error
}
