package services

import (
	"cp_tracker/user/models"

	"context"
	"crypto/sha256"
	"encoding/hex"
	"errors"

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
		collection: db.Collection("user"),
	}, nil
}

func (s *UserService) CreateUser(user *models.User) error {
	ctx := context.Background()

	// 檢查 email 是否已存在
	existingUser := &models.User{}
	filter := bson.M{"email": user.Email}
	if err := s.collection.FindOne(ctx, filter).Decode(existingUser); err == nil {
		return errors.New("duplicate email")
	}

	// 對密碼進行 hash
	hashedPassword := sha256.Sum256([]byte(user.Password))
	user.Password = hex.EncodeToString(hashedPassword[:])

	// 插入新用戶
	if res, err := s.collection.InsertOne(ctx, user); err != nil {
		return err
	} else {
		user.Uid = res.InsertedID.(primitive.ObjectID)
		return nil
	}
}

func (s *UserService) GetUserByUid(uid string) (*models.User, error) {
	ctx := context.Background()

	// 將 uid 轉換為 ObjectID
	var objectId primitive.ObjectID
	if res, err := primitive.ObjectIDFromHex(uid); err != nil {
		return nil, err
	} else {
		objectId = res
	}

	// 查找用戶
	var user models.User
	filter := bson.M{"_id": objectId}
	if err := s.collection.FindOne(ctx, filter).Decode(&user); err != nil {
		return nil, err
	} else {
		return &user, nil
	}
}

func (s *UserService) UpdateUser(user *models.User) error {
	ctx := context.Background()

	filter := bson.M{"_id": user.Uid}
	update := bson.M{
		"$set": bson.M{
			"displayName": user.DisplayName,
		},
	}

	_, err := s.collection.UpdateOne(ctx, filter, update)
	return err
}

func (s *UserService) UpdatePassword(uid primitive.ObjectID, newPassword string) error {
	ctx := context.Background()

	// 對新密碼進行 hash
	hashedPassword := sha256.Sum256([]byte(newPassword))

	filter := bson.M{"_id": uid}
	update := bson.M{
		"$set": bson.M{
			"password": hex.EncodeToString(hashedPassword[:]),
		},
	}

	_, err := s.collection.UpdateOne(ctx, filter, update)
	return err

}

func (s *UserService) DeleteUser(id string) error {
	ctx := context.Background()

	filter := bson.M{"_id": id}
	_, err := s.collection.DeleteOne(ctx, filter)
	return err
}

func (s *UserService) GetUserByEmail(email string) (*models.User, error) {
	ctx := context.Background()
	var user models.User
	filter := bson.M{"email": email}
	if err := s.collection.FindOne(ctx, filter).Decode(&user); err != nil {
		return nil, err
	}
	return &user, nil
}

func (s *UserService) CheckPassword(user *models.User, password string) bool {
	hashedPassword := sha256.Sum256([]byte(password))
	return user.Password == hex.EncodeToString(hashedPassword[:])
}
