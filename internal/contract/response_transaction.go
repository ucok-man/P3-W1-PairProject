package contract

import "go.mongodb.org/mongo-driver/bson/primitive"

type ResTransactionInsert struct {
	Data struct {
		Id          primitive.ObjectID `json:"id"`
		Description string             `json:"description"`
		Amount      float64            `json:"amount"`
	} `json:"data"`
	Message string `json:"message"`
}

type ResTransactionGetAll struct {
	Data []struct {
		Id          primitive.ObjectID `json:"id"`
		Description string             `json:"description"`
		Amount      float64            `json:"amount"`
	} `json:"data"`
}

type ResTransactionGetByID struct {
	Data struct {
		Id          primitive.ObjectID `json:"id"`
		Description string             `json:"description"`
		Amount      float64            `json:"amount"`
	} `json:"data"`
}
