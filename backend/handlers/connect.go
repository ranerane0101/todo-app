package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/ranerane0101/domain/valueobject"
	"github.com/ranerane0101/dto"
)

func GetTodos(w http.ResponseWriter, r *http.Request) {
	todos := []dto.TodoDTO{
		{ID: valueobject.NewTodoID("1"), Text: "買い物をする", Done: false},
		{ID: valueobject.NewTodoID("2"), Text: "本を読む", Done: true},
		{ID: valueobject.NewTodoID("3"), Text: "おにぎりを食べる", Done: true},
		{ID: valueobject.NewTodoID("4"), Text: "転職活動をする", Done: false},
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(todos)
}
