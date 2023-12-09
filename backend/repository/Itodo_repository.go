package repository

import "github.com/ranerane0101/domain/model"

// TodoRepositoryInterface はTodoモデルのリポジトリのインタフェースです。
type ITodoRepository interface {
	GetAllTodos(ID string) ([]model.Todo, error)
}
