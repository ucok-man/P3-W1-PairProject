package repo

import (
	"go.mongodb.org/mongo-driver/mongo"
)

type UserService struct {
	coll *mongo.Collection
}

// func (s *UserService) setup() {
// 	indexmodel := mongo.IndexModel{
// 		Keys:    bson.D{{"email", -1}},
// 		Options: options.Index().SetUnique(true),
// 	}

// 	if _, err := s.coll.Indexes().CreateOne(context.Background(), indexmodel); err != nil {
// 		panic(fmt.Sprintf("[user.setup()] ERROR setup index: %v\n", err))
// 	}
// }

// func (s *UserService) GetByEmail(email string) (*entity.User, error) {
// 	// User := entity.User{}

// 	// err := s.db.Where("email = $1", email).First(&User).Error
// 	// if err != nil {
// 	// 	switch {
// 	// 	case errors.Is(err, gorm.ErrRecordNotFound):
// 	// 		return nil, ErrRecordNotFound
// 	// 	default:
// 	// 		return nil, err
// 	// 	}
// 	// }
// 	// return &User, nil

// 	return nil, nil
// }

// func (s *UserService) GetByID(id int) (*entity.User, error) {
// 	// User := entity.User{}

// 	// err := s.db.Where("user_id = $1", id).First(&User).Error
// 	// if err != nil {
// 	// 	switch {
// 	// 	case errors.Is(err, gorm.ErrRecordNotFound):
// 	// 		return nil, ErrRecordNotFound
// 	// 	default:
// 	// 		return nil, err
// 	// 	}
// 	// }
// 	// return &User, nil
// 	return nil, nil
// }

// func (s *UserService) Insert(user *entity.User) (primitive.ObjectID, error) {
// 	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
// 	defer cancel()

// 	result, err := s.coll.InsertOne(ctx, user)
// 	if err != nil {
// 		switch {
// 		case mongo.IsDuplicateKeyError(err):
// 			return primitive.ObjectID{}, ErrDuplicateRecord
// 		default:
// 			return primitive.ObjectID{}, err
// 		}
// 	}

// 	objid, ok := result.InsertedID.(primitive.ObjectID)
// 	if !ok {
// 		return primitive.ObjectID{}, fmt.Errorf("id returned from driver is not object id: %v", err)
// 	}

// 	return objid, nil
// }
