package dto

import "github.com/ranerane0101/domain/valueobject"

type TodoDTO struct {
	TodoID   valueobject.TodoID `json:"id"`
	Text string             `json:"text"`
	Done bool               `json:"done"`
}
