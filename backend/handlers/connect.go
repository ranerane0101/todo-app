package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/ranerane0101/domain/valueobject"
	"github.com/ranerane0101/usecase"
)

func GetTodos(w http.ResponseWriter, r *http.Request, todoUsecase usecase.TodoUsecaseInterface) {
	// nil チェックを行う
	if todoUsecase == nil {
		http.Error(w, "Todo usecase is not initialized", http.StatusInternalServerError)
		return
	}

	todos, err := todoUsecase.ListTodos(valueobject.NewTodoID("1"))
	if err != nil {
		http.Error(w, "Failed to get todos", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(todos)
}
