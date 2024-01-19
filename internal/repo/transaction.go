package repo

import (
	"context"

	"github.com/ucok-man/P3-W1-PairProject/internal/entity"
	"go.mongodb.org/mongo-driver/mongo"
	"gopkg.in/mgo.v2/bson"
)

type TransactionService struct {
	coll *mongo.Collection
}

// func (s *TransactionService) setup() {
// 	indexmodel := mongo.IndexModel{
// 		Keys:    bson.D{{"id", -1}},
// 		Options: options.Index().SetUnique(true),
// 	}

// 	if _, err := s.coll.Indexes().CreateOne(context.Background(), indexmodel); err != nil {
// 		panic(fmt.Sprintf("[transaction.setup()] ERROR setup index: %v\n", err))
// 	}
// }

func (s *TransactionService) Create(ctx context.Context, req entity.Transaction) (*mongo.InsertOneResult, error) {
	res, err := s.coll.InsertOne(ctx, req)
	if err != nil {
		return res, err
	}

	return res, nil
}

func (s *TransactionService) Update(ctx context.Context, req entity.Transaction) (*mongo.UpdateResult, error) {
	update := bson.M{
		"$set": bson.M{
			"description": req.Description,
			"amount":      req.Amount,
		},
	}
	res, err := s.coll.UpdateByID(ctx, req.Id, update)
	if err != nil {
		return res, err
	}

	return res, nil
}

func (s *TransactionService) Delete(ctx context.Context, req entity.Transaction) (*mongo.DeleteResult, error) {
	filter := bson.M{"_id": req.Id}
	res, err := s.coll.DeleteOne(ctx, filter)
	if err != nil {
		return res, err
	}

	return res, nil
}
