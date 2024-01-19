package repo

import (
	"context"
	"errors"
	"fmt"

	"github.com/ucok-man/P3-W1-PairProject/internal/entity"
	"go.mongodb.org/mongo-driver/bson/primitive"
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

func (s *TransactionService) Insert(ctx context.Context, input *entity.Transaction) error {
	result, err := s.coll.InsertOne(ctx, input)
	if err != nil {
		return err
	}

	objid, ok := result.InsertedID.(primitive.ObjectID)
	if !ok {
		return fmt.Errorf("[Repo.TransactionService.Insert] object id is not primitive.ObjectID")
	}
	input.Id = objid

	return nil
}

func (s *TransactionService) GetAll(ctx context.Context) ([]*entity.Transaction, error) {
	cursor, err := s.coll.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}

	var transactions []*entity.Transaction
	err = cursor.All(ctx, &transactions)
	if err != nil {
		return nil, err
	}

	if len(transactions) == 0 {
		return make([]*entity.Transaction, 0), nil
	}

	return transactions, nil
}

func (s *TransactionService) GetByID(ctx context.Context, id primitive.ObjectID) (*entity.Transaction, error) {
	var transaction *entity.Transaction
	err := s.coll.FindOne(ctx, bson.M{"_id": id}).Decode(&transaction)
	if err != nil {
		switch {
		case errors.Is(err, mongo.ErrNoDocuments):
			return nil, ErrRecordNotFound
		default:
			return nil, err
		}
	}
	return transaction, nil
}
