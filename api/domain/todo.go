// domain/todo.go
package domain

type Todo struct {
	ID        int    `json:"id"`
	Text      string `json:"text"`
	Completed bool   `json:"completed"`
}

type TodoUsecase interface {
	AllGet() ([]Todo, error)
	StatusUpdate(id int) error
	Store(todo Todo) error
	Delete(id int) error
}

type TodoRepository interface {
	AllGet() ([]Todo, error)
	StatusUpdate(id int) error
	Store(todo Todo) error
	Delete(id int) error
}
