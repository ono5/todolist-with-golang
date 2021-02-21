// repository/todo_map.go
package repository

import (
	"errors"
	"sync"
	"todo/domain"
)

type todoRepository struct {
	m sync.Map
}

func NewSyncMapTodoRepository() domain.TodoRepository {
	return &todoRepository{}
}

// 全てのTodoを取得
func (t *todoRepository) AllGet() ([]domain.Todo, error) {
	var todos []domain.Todo

	t.m.Range(func(key interface{}, value interface{}) bool {
		todos = append(
			todos,
			// interface型をTodo型に変換
			value.(domain.Todo),
		)
		return true
	})

	return todos, nil
}

// Todoのステータスを更新
func (t *todoRepository) StatusUpdate(id int) error {
	r, ok := t.m.LoadAndDelete(id)
	if !ok {
		return errors.New("Fail Load Data")
	}

	newTodo := r.(domain.Todo)
	newTodo.Completed = true

	t.Store(newTodo)
	return nil
}

func (t *todoRepository) Store(todo domain.Todo) error {
	t.m.Store(todo.ID, todo)
	return nil
}

func (t *todoRepository) Delete(id int) error {
	t.m.Delete(id)
	return nil
}
