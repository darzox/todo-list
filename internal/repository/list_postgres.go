package repository

import (
	"database/sql"
	"fmt"
	"strings"

	"todo-app/internal/model"
)

type TodoListPostgres struct {
	db *sql.DB
}

func NewTodoListPostgres(db *sql.DB) *TodoListPostgres {
	return &TodoListPostgres{
		db: db,
	}
}

func (t *TodoListPostgres) Create(userId int, list model.TodoList) (int, error) {
	tx, err := t.db.Begin()
	if err != nil {
		return 0, err
	}
	var id int

	createListQuery := fmt.Sprintf("INSERT INTO %s (title, description) VALUES ($1, $2) RETURNING id", todoListTable)
	row := tx.QueryRow(createListQuery, list.Title, list.Description)
	if err := row.Scan(&id); err != nil {
		tx.Rollback()
		return 0, err
	}
	createUsersListQuery := fmt.Sprintf("INSERT INTO %s (user_id, list_id) VALUES ($1, $2) RETURNING id", usersListsTable)
	_, err = tx.Exec(createUsersListQuery, userId, id)
	if err != nil {
		tx.Rollback()
		return 0, err
	}

	return id, tx.Commit()
}

func (t *TodoListPostgres) GetAllLists(userId int) ([]model.TodoList, error) {
	var lists []model.TodoList
	query := fmt.Sprintf("SELECT lt.id, lt.title, lt.description FROM %s lt JOIN  %s ult ON lt.id = ult.list_id WHERE ult.user_id = $1", todoListTable, usersListsTable)
	rows, err := t.db.Query(query, userId)
	if err != nil {
		return []model.TodoList{}, err
	}
	for rows.Next() {
		var tempList model.TodoList
		rows.Scan(&tempList.Id, &tempList.Title, &tempList.Description)
		lists = append(lists, tempList)
	}
	return lists, nil
}

func (t *TodoListPostgres) GetListById(userId, listId int) (model.TodoList, error) {
	var list model.TodoList

	query := fmt.Sprintf("SELECT tl.id, tl.title, tl.description FROM %s tl JOIN %s ult ON tl.id = ult.list_id WHERE ult.user_id = $1 AND tl.id = $2", todoListTable, usersListsTable)
	row := t.db.QueryRow(query, userId, listId)
	err := row.Scan(&list.Id, &list.Title, &list.Description)
	if err != nil {
		return model.TodoList{}, err
	}
	return list, nil
}

func (t *TodoListPostgres) DeleteList(userId, listId int) error {
	query := fmt.Sprintf("DELETE FROM %s tl USING %s ult WHERE tl.id = ult.list_id AND ult.user_id = $1 AND ult.list_id = $2", todoListTable, usersListsTable)
	_, err := t.db.Exec(query, userId, listId)
	return err
}

func (t *TodoListPostgres) UpdateList(userId, listId int, updateList model.UpdateList) error {
	setValues := make([]string, 0)
	args := make([]interface{}, 0)
	argId := 1
	if updateList.Title != nil {
		setValues = append(setValues, fmt.Sprintf("title=$%d", argId))
		args = append(args, *updateList.Title)
		argId++
	}
	if updateList.Description != nil {
		setValues = append(setValues, fmt.Sprintf("description=$%d", argId))
		args = append(args, *updateList.Description)
		argId++
	}
	setQuery := strings.Join(setValues, ", ")
	query := fmt.Sprintf("UPDATE %s tl SET %s FROM %s ul WHERE tl.id = ul.list_id AND ul.list_id = $%d AND ul.user_id = $%d", todoListTable, setQuery, usersListsTable, argId, argId+1)
	args = append(args, listId, userId)

	_, err := t.db.Exec(query, args...)
	return err
}
