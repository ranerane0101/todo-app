package usecase

import (
	"github.com/ranerane0101/domain/valueobject"
	"github.com/ranerane0101/dto"
)

// TodoUsecaseInterface はToDoリストのユースケースのインタフェースです。
type TodoUsecaseInterface interface {
	ListTodos(TodoID valueobject.TodoID) ([]dto.TodoDTO, error)
}
