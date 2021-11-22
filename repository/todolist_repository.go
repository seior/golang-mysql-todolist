package repository

import (
	"context"
	"todolist/model"
)

type TodolistRepository interface {
	Create(ctx context.Context, todolist *model.Todolist) (int, error)
	Get(ctx context.Context) ([]model.Todolist, error)
	GetByIndex(ctx context.Context, id uint32) (model.Todolist, error)
	Delete(ctx context.Context, id uint32) (int, error)
	DeleteAll(ctx context.Context) (int, error)
}
