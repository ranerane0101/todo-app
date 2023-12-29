package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/ranerane0101/handlers"
	"github.com/ranerane0101/usecase"
	"github.com/rs/cors"
)

func main() {
	router := mux.NewRouter()

	// 新しい関数を作成し、GetTodos 関数を呼び出す
	getTodosHandler := func(w http.ResponseWriter, r *http.Request) {
		// ここで usecase を準備して GetTodos を呼び出す
		var todoUsecase usecase.TodoUsecaseInterface // 適切な方法でインスタンスを取得する
		handlers.GetTodos(w, r, todoUsecase)
	}

	router.HandleFunc("/api/todos", getTodosHandler).Methods("GET")

	handler := cors.Default().Handler(router)

	server := &http.Server{
		Addr:    ":5000",
		Handler: handler,
	}

	go func() {
		if err := server.ListenAndServe(); err != nil {
			fmt.Println("Failed to start server:", err)
		}
	}()

	fmt.Println("Server is running on port 8000")
	fmt.Println("Press CTRL+C to stop the server")
	select {}
}
