package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()

	server := http.Server{
		Addr:    ":8000",
		Handler: router,
	}
	log.Println("Server starting on port 8000 ......")
	err := server.ListenAndServe()
	if err != nil {
		log.Println("Error running the server:", err)

	}

}
