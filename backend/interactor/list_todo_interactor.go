package interactor

import (
	"github.com/ranerane0101/domain/entity"
	"github.com/ranerane0101/domain/valueobject"
	"github.com/ranerane0101/repository"
	"github.com/ranerane0101/usecase"
)

// TodoInteractor はToDoリストのユースケースを表します。
type TodoInteractor struct {
	TodoRepository repository.ITodoRepository
}

// NewTodoInteractor はToDoリストのインタラクターを作成します。
func NewTodoInteractor(todoRepo repository.ITodoRepository) usecase.TodoUsecaseInterface {
	return &TodoInteractor{
		TodoRepository: todoRepo,
	}
}

// GetTodoList は指定したユーザーのToDoリストを取得します。
func (ti *TodoInteractor) ListTodos(ID valueobject.TodoID) ([]entity.Todo, error) {
	todos, err := ti.TodoRepository.FindAllTodos(ID)
	if err != nil {
		return nil, err
	}

	return todos, nil
}
