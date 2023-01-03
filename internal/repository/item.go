package repository

import (
	"database/sql"
	"fmt"

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
	query := fmt.Sprintf("SELECT il.id, il.title, il.description, done FROM %s il JOIN %s lit ON il.id = lit.item_id WHERE lit.list_id = $1", todoItemsTable, listsItemsTable)
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
