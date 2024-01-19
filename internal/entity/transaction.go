package entity

import "go.mongodb.org/mongo-driver/bson/primitive"

type Transaction struct {
	Id          primitive.ObjectID `json:"id"`
	Description string             `json:"description"`
	Amount      float64            `json:"amount"`
}
