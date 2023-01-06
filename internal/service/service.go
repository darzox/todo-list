package service

import (
	"todo-app/internal/model"
	"todo-app/internal/repository"
)

type Authorization interface {
	CreateUser(user model.User) (int, error)
	GenerateToken(username, password string) (string, error)
	ParseToken(token string) (int, error)
}

type TodoList interface {
	Create(userId int, list model.TodoList) (int, error)
	GetAllLists(userId int) ([]model.TodoList, error)
	GetListById(userId, listId int) (model.TodoList, error)
	DeleteList(userId, listId int) error
	UpdateList(userId, listId int, updateList model.UpdateList) error
}

type TodoItem interface {
	CreateItem(userId, listId int, item model.TodoItem) (int, error)
	GetAllItems(userId, listId int) ([]model.TodoItem, error)
	GetItem(userId, listId, itemId int) (model.TodoItem, error)
	UpdateItem(userId, listId, itemId int, input model.UpdateItem) error
	DeleteItem(userId, listId, itemId int) error
}

type Service struct {
	Authorization
	TodoItem
	TodoList
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repos.Authorization),
		TodoItem:      NewTodoItemService(repos.TodoItem, repos.TodoList),
		TodoList:      NewTodoListService(repos.TodoList),
	}
}
