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
	User UserService
}

func New(db *mongo.Database) *Services {
	return &Services{
		User: UserService{coll: db.Collection("user")},
	}
}

func (r *Services) InitSetup() *Services {
	r.User.setup()
	return r
}
