package main

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

type Todo struct {
	ID   int    `json:"ID"`
	Text string `json:"Text"`
	Done bool   `json:"Done"`
}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/api/todos", GetTodos).Methods("GET")

	handler := cors.Default().Handler(router) // CORS設定を適用したハンドラを作成

	http.Handle("/", handler)
	http.ListenAndServe(":5000", nil)
}

func GetTodos(w http.ResponseWriter, r *http.Request) {
	todos := []Todo{
		{ID: 1, Text: "buy groceries", Done: false},
		{ID: 2, Text: "Read a book", Done: true},
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(todos)
}