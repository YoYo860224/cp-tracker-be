package services

import (
	"context"
	"cp_tracker/user/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserService struct {
	db         *mongo.Database
	collection *mongo.Collection
}

func NewUserService(db *mongo.Database) (*UserService, error) {
	if db == nil {
		return nil, mongo.ErrNilDocument
	}

	return &UserService{
		db:         db,
		collection: db.Collection("users"),
	}, nil
}

// CreateUser implements UserService.
func (s *UserService) CreateUser(user *models.User) error {
	ctx := context.Background()

	// 如果沒有 ID 則產生一個新的
	if user.Uid == "" {
		objectID := primitive.NewObjectID()
		user.Uid = objectID.Hex()
	}

	_, err := s.collection.InsertOne(ctx, user)
	return err
}

// GetUserByID implements UserService.
func (s *UserService) GetUserByUid(uid string) (*models.User, error) {
	ctx := context.Background()

	var user models.User
	filter := bson.M{"uid": uid}

	err := s.collection.FindOne(ctx, filter).Decode(&user)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

// UpdateUser implements UserService.
func (s *UserService) UpdateUser(user *models.User) error {
	ctx := context.Background()

	filter := bson.M{"uid": user.Uid}
	update := bson.M{
		"$set": bson.M{
			"username":    user.Username,
			"email":       user.Email,
			"password":    user.Password,
			"displayName": user.DisplayName,
			"role":        user.Role,
		},
	}

	_, err := s.collection.UpdateOne(ctx, filter, update)
	return err
}

// DeleteUser implements UserService.
func (s *UserService) DeleteUser(id string) error {
	ctx := context.Background()

	filter := bson.M{"_id": id}
	_, err := s.collection.DeleteOne(ctx, filter)
	return err
}
