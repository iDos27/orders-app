package controllers

import (
	"context"
	"encoding/json"
	"math/rand"
	"net/http"
	"time"

	"orders-app/db"
	"orders-app/models"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func CreateOrder(w http.ResponseWriter, r *http.Request) {
	var order models.Order
	json.NewDecoder(r.Body).Decode(&order)

	order.Status = "pending"
	order.OrderNumber = generateOrderNumber()

	collection := db.DB.Collection("orders")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err := collection.InsertOne(ctx, order)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(order)
}

func GetOrders(w http.ResponseWriter, r *http.Request) {
	collection := db.DB.Collection("orders")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	cursor, _ := collection.Find(ctx, bson.M{})
	var orders []models.Order
	cursor.All(ctx, &orders)

	json.NewEncoder(w).Encode(orders)
}

func UpdateStatus(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := primitive.ObjectIDFromHex(vars["id"])

	var payload struct {
		Status string `json:"status"`
	}
	json.NewDecoder(r.Body).Decode(&payload)

	collection := db.DB.Collection("orders")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err := collection.UpdateOne(ctx, bson.M{"_id": id}, bson.M{"$set": bson.M{"status": payload.Status}})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func generateOrderNumber() string {
	return "ZAM-" + time.Now().Format("20060102") + "-" + string(rand.Intn(1000))
}
