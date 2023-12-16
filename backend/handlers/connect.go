package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/ranerane0101/domain/entity"
	"github.com/ranerane0101/domain/valueobject"
)

func GetTodos(w http.ResponseWriter, r *http.Request) {
	todos := []entity.Todo{
		{ID: valueobject.NewTodoID("1"), Text: "buy groceries", Done: false},
		{ID: valueobject.NewTodoID("2"), Text: "Read a book", Done: true},
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(todos)
}
