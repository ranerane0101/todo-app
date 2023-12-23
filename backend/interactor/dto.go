package interactor

import (
	"github.com/ranerane0101/domain/entity"
	"github.com/ranerane0101/domain/valueobject"
)

type TodoDTO struct {
	ID   valueobject.TodoID `json:"id"`
	Text string             `json:"text"`
	Done bool               `json:"done"`
}

// EntityからDTOへの変換用の関数
func ConvertToTodoDTO(entity entity.Todo) TodoDTO {
	//entityの値をDTOに設定
	return TodoDTO{
		ID:   entity.ID,
		Text: entity.Text,
		Done: entity.Done, //達成済みか否か
	}
}
