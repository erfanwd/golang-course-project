package repo_interfaces

import "context"

type BaseRepositoryInterface[T any] interface {
	Create(ctx context.Context, entity *T) (T, error)
	FindByID(ctx context.Context, id uint) (*T, error)
	GetAll(ctx context.Context) ([]T, error)
	Update(ctx context.Context, entity *T) error
	Delete(ctx context.Context, entity *T) error
}
