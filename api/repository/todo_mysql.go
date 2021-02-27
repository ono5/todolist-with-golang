// repository/todo_mysql.go
package repository

import (
	"errors"
	"fmt"
	"todo/domain"

	"gorm.io/gorm"
)

type todoRepositoryMySQL struct {
	db *gorm.DB
}

func NewTodoRepositoryMySQL(db *gorm.DB) domain.TodoRepository {
	return &todoRepositoryMySQL{
		db: db,
	}
}

// 全Todo取得
func (t *todoRepositoryMySQL) AllGet() ([]domain.Todo, error) {
	var todos []domain.Todo
	// https://gorm.io/ja_JP/docs/query.html#Retrieving-all-objects
	result := t.db.Find(&todos)
	if result.Error != nil {
		return nil, errors.New("Unexpected Error in retrieving all the records ")
	}
	return todos, nil
}

// Todoステータス更新
func (t *todoRepositoryMySQL) StatusUpdate(id int) error {
	var todo domain.Todo
	result := t.db.First(&todo, id)
	if result.Error != nil {
		return errors.New("Unexpected Error in getting the todo")
	}

	// ステータス更新
	if todo.Completed {
		todo.Completed = false
	} else {
		todo.Completed = true
	}

	result = t.db.Save(&todo)
	if result.Error != nil {
		return errors.New("Unexpected Error in updating the todo")
	}

	return nil
}

// Todo保存
func (t *todoRepositoryMySQL) Store(todo domain.Todo) error {
	result := t.db.Create(todo)
	if result.Error != nil {
		return errors.New("Unexpected Error in creating the todo")
	}
	return nil
}

// Todo削除
func (t *todoRepositoryMySQL) Delete(id int) error {
	result := t.db.Delete(&domain.Todo{}, id)
	if result.Error != nil {
		return errors.New("Unexpected Error in deliting the todo")
	}
	return nil
}

// Todo検索
func (t *todoRepositoryMySQL) Search(key string) ([]domain.Todo, error) {
	var todos []domain.Todo
	result := t.db.Where("text LIKE ?", fmt.Sprintf("%%%s%%", key)).Find(&todos)
	if result.Error != nil {
		return nil, errors.New("Unexpected Error in retrieving search todos ")
	}
	return todos, nil
}
