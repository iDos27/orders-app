package main

import (
	"log"
	"net/http"
	"orders-app/controllers"
	"orders-app/db"

	"github.com/gorilla/mux"
)

func main() {
	db.Connect()
	r := mux.NewRouter()
	r.HandleFunc("/orders", controllers.CreateOrder).Methods("POST")
	r.HandleFunc("/orders", controllers.GetOrders).Methods("GET")
	r.HandleFunc("/orders/{id}/status", controllers.UpdateStatus).Methods("PATCH")

	log.Println("Server is running on port 8080")
	http.ListenAndServe(":8080", r)
}
