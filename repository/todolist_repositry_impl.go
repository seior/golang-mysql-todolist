package repository

import (
	"context"
	"database/sql"
	"errors"
	"todolist/model"
)

type todolistRepositoryImpl struct {
	DB *sql.DB
}

func NewTodolistRepository(db *sql.DB) TodolistRepository {
	return &todolistRepositoryImpl{DB: db}
}

// Create mengembalikan nilai id jika berhasil. jika gagal, maka akan mengembalikan 0
func (repository *todolistRepositoryImpl) Create(ctx context.Context, todolist *model.Todolist) (int, error) {
	script := "INSERT INTO todolist(name, author) value (?, ?)"
	result, err := repository.DB.ExecContext(ctx, script, todolist.Name, todolist.Author)
	if err != nil {
		return 0, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return int(id), nil
}

// Get mengembalikan semua data yang ada di database todolist
func (repository *todolistRepositoryImpl) Get(ctx context.Context) ([]model.Todolist, error) {
	script := "SELECT id, name, author FROM todolist"
	var todolists []model.Todolist

	result, err := repository.DB.QueryContext(ctx, script)
	defer result.Close()
	if err != nil {
		return todolists, err
	}

	for result.Next() {
		todolist := model.Todolist{}

		err := result.Scan(&todolist.Id, &todolist.Name, &todolist.Author)
		if err != nil {
			return todolists, err
		}
		todolists = append(todolists, todolist)
	}

	return todolists, nil
}

// GetByIndex Mendapatkan todolist berdasarkan id
func (repository *todolistRepositoryImpl) GetByIndex(ctx context.Context, id uint32) (model.Todolist, error) {
	script := "SELECT id, name, author FROM todolist WHERE id = ?"
	var todolist model.Todolist

	result, err := repository.DB.QueryContext(ctx, script, id)
	defer result.Close()
	if err != nil {
		return todolist, err
	}

	for result.Next() {
		err := result.Scan(&todolist.Id, &todolist.Name, &todolist.Author)
		if err != nil {
			return todolist, err
		}

		if len(todolist.Name) < 1 {
			return todolist, errors.New("Todolist tidak ditemukan")
		}
	}

	return todolist, nil
}

// Delete menghapus todolist dari database
func (repository *todolistRepositoryImpl) Delete(ctx context.Context, id uint32) (int, error) {
	script := "DELETE FROM todolist WHERE id = ?"
	_, err := repository.DB.ExecContext(ctx, script, id)
	if err != nil {
		return 0, err
	}

	return int(id), nil
}

// DeleteAll menghapus semua data dari todolist
func (repository *todolistRepositoryImpl) DeleteAll(ctx context.Context) (int, error) {
	script := "DELETE FROM todolist"
	_, err := repository.DB.ExecContext(ctx, script)
	if err != nil {
		return 0, err
	}

	return 1, nil
}
