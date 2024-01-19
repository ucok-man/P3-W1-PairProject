package repo

import (
	"errors"

	"go.mongodb.org/mongo-driver/mongo"
)

var (
	ErrRecordNotFound  = errors.New("no record found")
	ErrDuplicateRecord = errors.New("record duplicate on unique constraint")
)

type Services struct {
	User        UserService
	Transaction TransactionService
}

func New(db *mongo.Database) *Services {
	return &Services{
		User:        UserService{coll: db.Collection("user")},
		Transaction: TransactionService{coll: db.Collection("transaction")},
	}
}

func (r *Services) InitSetup() *Services {
	r.User.setup()
	r.Transaction.setup()
	return r
}
