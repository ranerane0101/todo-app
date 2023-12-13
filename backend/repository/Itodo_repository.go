package repository

import "github.com/ranerane0101/domain/model"

// TodoRepositoryInterface はTodoモデルのリポジトリのインタフェースです。
type ITodoRepository interface {
	FindAllTodos(ID string) ([]model.Todo, error)
	FindTodoByID(todoID string) (*model.Todo, error)
}
