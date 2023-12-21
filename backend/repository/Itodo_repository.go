package repository

import (
	"github.com/ranerane0101/domain/entity"
	"github.com/ranerane0101/domain/valueobject"
)

// TodoRepositoryInterface はTodoモデルのリポジトリのインタフェースです。
type ITodoRepository interface {
	FindAllTodos(ID valueobject.TodoID) ([]entity.Todo, error)
	FindTodoByID(ID valueobject.TodoID) (*entity.Todo, error)
}
