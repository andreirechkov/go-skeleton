package domain

import "context"

type Repository interface {
	List(ctx context.Context) ([]Organisation, error)
}