package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/ranerane0101/handlers"
	"github.com/rs/cors"
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/api/todos", handlers.GetTodos).Methods("GET")

	handler := cors.Default().Handler(router)

	server := &http.Server{
		Addr:    ":52048",
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
