package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/ranerane0101/domain/model"
	"github.com/rs/cors"
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/api/todos", GetTodos).Methods("GET")

	handler := cors.Default().Handler(router)

	// サーバー起動前にログ出力
	fmt.Println("Starting server on :5000...")
	http.ListenAndServe(":5000", handler)
}

func GetTodos(w http.ResponseWriter, r *http.Request) {
	// 新しいモデルを使用する
	todos := []model.Todo{
		{ID: 1, Text: "buy groceries", Done: false},
		{ID: 2, Text: "Read a book", Done: true},
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(todos)
}
