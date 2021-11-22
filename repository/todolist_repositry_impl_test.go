package repository

import (
	"context"
	"github.com/stretchr/testify/assert"
	"testing"
	"todolist/model"
	"todolist/utils"
)

func TestMain(m *testing.M) {
	ctx := context.Background()
	repository := NewTodolistRepository(utils.GetDatabase())

	m.Run()

	repository.DeleteAll(ctx)
}

func TestTodolistRepositoryImpl_Create(t *testing.T) {
	ctx := context.Background()
	repository := NewTodolistRepository(utils.GetDatabase())

	todolist := model.Todolist{Name: "Learn Golang", Author: "Raden Mohamad Rishawn"}

	result, err := repository.Create(ctx, &todolist)
	if err != nil {
		panic(err)
	}

	assert.NotNil(t, result)
}

func TestTodolistRepositoryImpl_Get(t *testing.T) {
	ctx := context.Background()
	repository := NewTodolistRepository(utils.GetDatabase())

	todolist := model.Todolist{Name: "Learn Golang", Author: "Raden Mohamad Rishawn"}

	_, err := repository.Create(ctx, &todolist)
	if err != nil {
		panic(err)
	}

	_, err = repository.Create(ctx, &todolist)
	if err != nil {
		panic(err)
	}

	result, err := repository.Get(ctx)
	if err != nil {
		panic(err)
	}

	assert.True(t, len(result) > 0)
}

func TestTodolistRepositoryImpl_GetByIndex(t *testing.T) {
	ctx := context.Background()
	repository := NewTodolistRepository(utils.GetDatabase())

	todolist := model.Todolist{Name: "Learn Golang", Author: "Raden Mohamad Rishawn"}

	result, err := repository.Create(ctx, &todolist)
	if err != nil {
		panic(err)
	}

	newTodolist, _ := repository.GetByIndex(ctx, uint32(result))

	assert.Equal(t, result, int(newTodolist.Id))
	assert.Equal(t, todolist.Name, newTodolist.Name)
	assert.Equal(t, todolist.Author, newTodolist.Author)
}

func TestTodolistRepositoryImpl_Delete(t *testing.T) {
	ctx := context.Background()
	repository := NewTodolistRepository(utils.GetDatabase())

	todolist := model.Todolist{Name: "Learn Golang", Author: "Raden Mohamad Rishawn"}

	result, err := repository.Create(ctx, &todolist)
	if err != nil {
		panic(err)
	}

	_, err = repository.Delete(ctx, uint32(result))
	if err != nil {
		panic(err)
	}

	_, err = repository.GetByIndex(ctx, uint32(result))

	assert.Nil(t, err)
}