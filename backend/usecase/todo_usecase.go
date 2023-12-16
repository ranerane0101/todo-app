package usecase

import "github.com/ranerane0101/domain/entity"

// TodoUsecaseInterface はToDoリストのユースケースのインタフェースです。
type TodoUsecaseInterface interface {
	GetTodoList(userID string) ([]entity.Todo, error)
}
