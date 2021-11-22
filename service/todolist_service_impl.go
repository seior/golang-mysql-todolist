package service

import (
	"context"
	"fmt"
	"strconv"
	"todolist/model"
	"todolist/repository"
)

type todolistServiceImpl struct {
	repository repository.TodolistRepository
}

func NewTodolistService(todoolist repository.TodolistRepository) TodolistService {
	return &todolistServiceImpl{todoolist}
}

func (service todolistServiceImpl) ShowTodolist(ctx context.Context) {
	get, err := service.repository.Get(ctx)
	if err != nil {
		panic(err)
	}

	fmt.Println("\nTodolist : ")
	for _, todolist := range get {
		fmt.Println(strconv.Itoa(int(todolist.Id)) + ". " + todolist.Name)
	}
}

func (service todolistServiceImpl) CreateTodolist(ctx context.Context, todolist *model.Todolist) {
	_, err := service.repository.Create(ctx, todolist)
	if err != nil {
		panic(err)
	}
}

func (service todolistServiceImpl) DeleteTodolist(ctx context.Context, id *int) {
	_, err := service.repository.Delete(ctx, uint32(*id))
	if err != nil {
		panic(err)
	}
}
