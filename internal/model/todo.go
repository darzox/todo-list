package model

import "errors"

type TodoList struct {
	Id          int    `json:"id"`
	Title       string `json:"title" binding:"required"`
	Description string `json:"description"`
}

type UserList struct {
	Id     int
	UserId int
	ListId int
}

type TodoItem struct {
	Id          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Done        bool   `json:"done"`
}

type ListsItem struct {
	Id     int
	ListId int
	ItemId int
}

type UpdateList struct {
	Title       *string `json:"title"`
	Description *string `json:"description"`
}

func (u UpdateList) Validate() error {
	if u.Title == nil && u.Description == nil {
		return errors.New("update structure has no values")
	}
	return nil
}

type UpdateItem struct {
	Title       *string `json:"title"`
	Description *string `json:"description"`
	Done        *bool   `json:"done"`
}

func (u UpdateItem) ValidateItem() error {
	if u.Title == nil && u.Description == nil && u.Done == nil {
		return errors.New("update structure has no values")
	}
	return nil
}
