package repository

import (
	"database/sql"

	"todo-app/internal/model"
)

const (
	userTable       = "users"
	todoListTable   = "lists"
	usersListsTable = "users_lists"
	todoItemsTable  = "items"
	listsItemsTable = "lists_items"
)

type Authorization interface {
	CreateUser(user model.User) (int, error)
	GetUser(username, password string) (model.User, error)
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
	GetAllItems(listId int) ([]model.TodoItem, error)
}

type Repository struct {
	Authorization
	TodoItem
	TodoList
}

func NewRepository(db *sql.DB) *Repository {
	return &Repository{
		Authorization: NewAuthPostgres(db),
		TodoList:      NewTodoListPostgres(db),
		TodoItem:      NewTodoItemPostgres(db),
	}
}
