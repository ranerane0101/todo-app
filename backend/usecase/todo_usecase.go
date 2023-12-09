package usecase

import "github.com/ranerane0101/domain/model"

// TodoUsecaseInterface はToDoリストのユースケースのインタフェースです。
type TodoUsecaseInterface interface {
	GetTodoList(userID string) ([]model.Todo, error)
}
