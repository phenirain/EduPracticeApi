package service

import (
	"api/internal/infrastructure"
	"context"
)

type Repository[TDB infrastructure.DbModel[T], T infrastructure.Model] interface {
	Create(ctx context.Context, model T) (T, error)
	GetAll(ctx context.Context) ([]T, error)
	Update(ctx context.Context, model T) error
	Delete(ctx context.Context, id int32) error
}
