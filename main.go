package main

import (
	"log"
	"net/http"
	"time"
	"github.com/gorilla/mux"
	"example.com/poc/handler"
)

func main() {
	// Initialize the router
	router := mux.NewRouter()
	// Set up routes and corresponding handlers
	router.HandleFunc("/user/submitcard", handler.SubmitCreditCardUserRole).Methods("POST")
	router.HandleFunc("/admin/submitcards", handler.SubmitCreditCardAdminRole).Methods("POST")
	router.HandleFunc("/getallcards", handler.GetAllCreditCards).Methods("GET")
	router.HandleFunc("/getsinglecard/{id}", handler.GetCreditCard).Methods("GET")
	// Start the server

	srv := &http.Server{
		Addr:         ":8000",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
		Handler: router,
	}
	log.Println("Server started. Listening on port 8000...")

	log.Fatal(srv.ListenAndServe())

}

