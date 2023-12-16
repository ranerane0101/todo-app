package repository

import "github.com/ranerane0101/domain/entity"

// TodoRepositoryInterface はTodoモデルのリポジトリのインタフェースです。
type ITodoRepository interface {
	FindAllTodos(ID string) ([]entity.Todo, error)
	FindTodoByID(todoID string) (*entity.Todo, error)
}
