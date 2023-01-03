package service

import (
	"todo-app/internal/model"
	"todo-app/internal/repository"
)

type TodoListService struct {
	repo repository.TodoList
}

func NewTodoListService(repo repository.TodoList) *TodoListService {
	return &TodoListService{
		repo: repo,
	}
}

func (t *TodoListService) Create(userId int, list model.TodoList) (int, error) {
	return t.repo.Create(userId, list)
}

func (t *TodoListService) GetAllLists(userId int) ([]model.TodoList, error) {
	return t.repo.GetAllLists(userId)
}

func (t *TodoListService) GetListById(userId, listId int) (model.TodoList, error) {
	return t.repo.GetListById(userId, listId)
}

func (t *TodoListService) DeleteList(userId, listId int) error {
	return t.repo.DeleteList(userId, listId)
}

func (t *TodoListService) UpdateList(userId, listId int, updateList model.UpdateList) error {
	if err := updateList.Validate(); err != nil {
		return err
	}
	return t.repo.UpdateList(userId, listId, updateList)
}
