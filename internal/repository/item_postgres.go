package repository

import (
	"database/sql"
	"fmt"
	"strings"

	"todo-app/internal/model"
)

type TodoItemPostgress struct {
	db *sql.DB
}

func NewTodoItemPostgres(db *sql.DB) *TodoItemPostgress {
	return &TodoItemPostgress{
		db: db,
	}
}

func (t *TodoItemPostgress) CreateItem(userId, listId int, item model.TodoItem) (int, error) {
	tx, err := t.db.Begin()
	if err != nil {
		return 0, err
	}
	query := fmt.Sprintf(`INSERT INTO %s (title, description) VALUES ($1, $2) RETURNING id`, todoItemsTable)
	var itemId int

	row := tx.QueryRow(query, item.Title, item.Description)
	err = row.Scan(&itemId)
	if err != nil {
		tx.Rollback()
		return 0, err
	}

	query = fmt.Sprintf(`INSERT INTO %s (list_id, item_id) VALUES ($1, $2)`, listsItemsTable)
	_, err = tx.Exec(query, listId, itemId)
	if err != nil {
		tx.Rollback()
		return 0, err
	}
	return itemId, tx.Commit()
}

func (t *TodoItemPostgress) GetAllItems(listId int) ([]model.TodoItem, error) {
	query := fmt.Sprintf("SELECT il.id, il.title, il.description, il.done FROM %s il JOIN %s lit ON il.id = lit.item_id WHERE lit.list_id = $1", todoItemsTable, listsItemsTable)
	var items []model.TodoItem
	rows, err := t.db.Query(query, listId)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var tempItem model.TodoItem
		rows.Scan(&tempItem.Id, &tempItem.Title, &tempItem.Description, &tempItem.Done)
		items = append(items, tempItem)
	}

	return items, nil
}

func (t *TodoItemPostgress) GetItem(itemId int) (model.TodoItem, error) {
	query := fmt.Sprintf("SELECT il.id, il.title, il.description, il.done FROM %s il WHERE il.id = $1", todoItemsTable)
	var item model.TodoItem
	err := t.db.QueryRow(query, itemId).Scan(&item.Id, &item.Title, &item.Description, &item.Done)
	if err != nil {
		return model.TodoItem{}, err
	}
	return item, nil
}

func (t *TodoItemPostgress) UpdateItem(itemId int, input model.UpdateItem) error {
	setValues := make([]string, 0)
	args := make([]interface{}, 0)
	argId := 1
	if input.Title != nil {
		setValues = append(setValues, fmt.Sprintf("title=$%d", argId))
		args = append(args, *input.Title)
		argId++
	}
	if input.Description != nil {
		setValues = append(setValues, fmt.Sprintf("description=$%d", argId))
		args = append(args, *input.Description)
		argId++
	}
	if input.Done != nil {
		setValues = append(setValues, fmt.Sprintf("done=$%d", argId))
		args = append(args, *input.Done)
		argId++
	}
	setQuery := strings.Join(setValues, ", ")
	query := fmt.Sprintf("UPDATE %s il SET %s WHERE il.id = $%d", todoItemsTable, setQuery, argId)
	fmt.Println(query)
	args = append(args, itemId)
	fmt.Println(args)

	_, err := t.db.Exec(query, args...)
	return err
}

func (t *TodoItemPostgress) DeleteItem(itemId int) error {
	query := fmt.Sprintf("DELETE FROM %s il WHERE il.id = $1", todoItemsTable)
	_, err := t.db.Exec(query, itemId)
	return err
}
