package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/ranerane0101/domain/model"
)

func GetTodos(w http.ResponseWriter, r *http.Request) {
	todos := []model.Todo{
		{ID: 1, Text: "buy groceries", Done: false},
		{ID: 2, Text: "Read a book", Done: true},
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(todos)
}
