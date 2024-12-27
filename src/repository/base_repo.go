// repository/base_repository.go
package repository

import (
	"context"

	"gorm.io/gorm"
)

type BaseRepo[T any] struct {
	Database *gorm.DB
}

func NewBaseRepository[T any](db *gorm.DB) *BaseRepo[T] {
	return &BaseRepo[T]{Database: db}
}

func (r *BaseRepo[T]) FindByID(ctx context.Context, id uint) (*T, error) {
	var entity T
	err := r.Database.WithContext(ctx).First(&entity, id).Error
	if err != nil {
		return nil, err
	}
	return &entity, nil
}

func (r *BaseRepo[T]) Create(ctx context.Context, entity *T) (T, error) {
	err := r.WithTransaction(ctx, func(tx *gorm.DB) error {
		if err := tx.Create(entity).Error; err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return *entity, err
	}
	return *entity, nil
}

func (r *BaseRepo[T]) Update(ctx context.Context, entity *T) error {
	return r.Database.WithContext(ctx).Save(entity).Error
}

func (r *BaseRepo[T]) Delete(ctx context.Context, entity *T) error {
	return r.Database.WithContext(ctx).Delete(entity).Error
}

func (r *BaseRepo[T]) WithTransaction(ctx context.Context, fn func(tx *gorm.DB) error) error {
	tx := r.Database.WithContext(ctx).Begin()
	if err := fn(tx); err != nil {
		tx.Rollback()
		return err
	}
	tx.Commit()
	return nil
}

func (r *BaseRepo[T]) GetAll(ctx context.Context) ([]T, error) {
	var entities []T
	err := r.Database.WithContext(ctx).Find(&entities).Error
	if err != nil {
		return nil, err
	}
	return entities, nil
}
