package interactor

import (
	"github.com/ranerane0101/domain/entity"
	"github.com/ranerane0101/dto"
)

// EntityからDTOへの変換用の関数
func ConvertToTodoDTO(entity entity.Todo) dto.TodoDTO {
	//entityの値をDTOに設定
	return dto.TodoDTO{
		TodoID: entity.TodoID,
		Text:   entity.Text,
		Done:   entity.Done, //達成済みか否か
	}
}
