package segment

import "context"

type Segment struct {
	ID          int      `json:"id"`
	Name        string   `json:"name"`
	Description string   `json:"description"`
	Metadata    []string `json:"metadata"`
	UserList    []int    `json:"user_list"`
}

type UseCase interface {
	Validate(user *Segment, create bool) error

	GetByID(ctx context.Context, id int) (*Segment, error)
	PutIfExits(ctx context.Context, segment *Segment) error
	Create(ctx context.Context, segment *Segment) error
	Delete(ctx context.Context, id int) error
	GetAll(ctx context.Context) ([]*Segment, error)

	Repository() Repository
}
