package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Order struct {
	ID          primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	OrderNumber string             `json:"order_number"`
	Options     []string           `json:"options"`
	Status      string             `json:"status"`
}
