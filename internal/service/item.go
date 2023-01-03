package service

import (
	"todo-app/internal/model"
	"todo-app/internal/repository"
)

type TodoItemService struct {
	repo     repository.TodoItem
	listRepo repository.TodoList
}

func NewTodoItemService(repo repository.TodoItem, listRepo repository.TodoList) *TodoItemService {
	return &TodoItemService{
		repo:     repo,
		listRepo: listRepo,
	}
}

func (t *TodoItemService) CreateItem(userId, listId int, item model.TodoItem) (int, error) {
	_, err := t.listRepo.GetListById(userId, listId)
	if err != nil {
		// list does not exists or does not belongs to user
		return 0, err
	}
	return t.repo.CreateItem(userId, listId, item)
}

func (t *TodoItemService) GetAllItems(userId, listId int) ([]model.TodoItem, error) {
	_, err := t.listRepo.GetListById(userId, listId)
	if err != nil {
		// list does not exists or does not belongs to user
		return nil, err
	}
	return t.repo.GetAllItems(listId)
}
