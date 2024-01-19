package contract

import "go.mongodb.org/mongo-driver/bson/primitive"

type ReqTransactionCreate struct {
	Description string  `json:"description" validate:"required"`
	Amount      float64 `json:"amount" validate:"required"`
}

type ReqTransactionUpdate struct {
	Id          primitive.ObjectID `param:"id" validate:"required"`
	Description string             `json:"description" validate:"required"`
	Amount      float64            `json:"amount" validate:"required"`
}

type ReqTransactionDelete struct {
	Id primitive.ObjectID `param:"id" validate:"required"`
}

type ReqTransactionInsert struct {
	Description string  `json:"description" validate:"required"`
	Amount      float64 `json:"amount" validate:"required"`
}
