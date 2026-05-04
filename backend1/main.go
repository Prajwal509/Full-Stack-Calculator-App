package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/calculator/backend/handlers"
	"github.com/calculator/backend/middleware"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	mux := http.NewServeMux()
	mux.HandleFunc("/api/calculate", handlers.CalculateHandler)
	mux.HandleFunc("/api/health", handlers.HealthHandler)

	handler := middleware.Logger(middleware.CORS(mux))

	fmt.Printf("Backend listening on :%s\n", port)
	log.Fatal(http.ListenAndServe(":"+port, handler))
}
