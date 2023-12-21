package usecase

import (
	"github.com/ranerane0101/domain/entity"
	"github.com/ranerane0101/domain/valueobject"
)

// TodoUsecaseInterface はToDoリストのユースケースのインタフェースです。
type TodoUsecaseInterface interface {
	ListTodos(ID valueobject.TodoID) ([]entity.Todo, error)
}
