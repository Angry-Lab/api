package city

import "context"

type Repository interface {
	Cities(ctx context.Context, postcodes ...int) (map[int]*City, error)
}
