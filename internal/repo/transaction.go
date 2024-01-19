package repo

import (
	"context"
	"fmt"

	"github.com/ucok-man/P3-W1-PairProject/internal/entity"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"gopkg.in/mgo.v2/bson"
)

type TransactionService struct {
	coll *mongo.Collection
}

func (s *TransactionService) setup() {
	indexmodel := mongo.IndexModel{
		Keys:    bson.D{{"id", -1}},
		Options: options.Index().SetUnique(true),
	}

	if _, err := s.coll.Indexes().CreateOne(context.Background(), indexmodel); err != nil {
		panic(fmt.Sprintf("[repo.setup()] ERROR setup index: %v\n", err))
	}
}

func (s *TransactionService) Update(ctx context.Context, req entity.Transaction) (*mongo.UpdateResult, error) {
	res, err := s.coll.UpdateByID(ctx, req.Id, req)
	if err != nil {
		return res, err
	}

	return res, nil
}

func (s *TransactionService) Delete(ctx context.Context, req entity.Transaction) (*mongo.DeleteResult, error) {
	res, err := s.coll.DeleteOne(ctx, req)
	if err != nil {
		return res, err
	}

	return res, nil
}
