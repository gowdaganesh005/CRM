package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/user", Getallusers).Methods("GET")
	router.HandleFunc("/user", CreateUser).Methods("POST")
	router.HandleFunc("/user/{id}", GetUserByID).Methods("GET")
	router.HandleFunc("/user/{id}", Delete).Methods("DELETE")
	router.HandleFunc("/user/update/{id}", Update).Methods("POST")

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
