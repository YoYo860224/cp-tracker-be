package services

import (
	"cp_tracker/user/models"

	"context"
	"errors"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type UserDataService struct {
	collection *mongo.Collection
}

func NewUserDataService(db *mongo.Database) *UserDataService {
	return &UserDataService{
		collection: db.Collection("user_data"),
	}
}

func (s *UserDataService) GetUserData(uid string) (models.UserData, error) {
	ctx := context.Background()

	filter := bson.M{"uid": uid}
	var userData models.UserData
	err := s.collection.FindOne(ctx, filter).Decode(&userData)
	if errors.Is(err, mongo.ErrNoDocuments) {
		// 若查無資料，自動建立一筆空資料
		emptyData := models.UserData{
			Uid:   uid,
			Items: map[string]interface{}{},
		}
		_, err := s.collection.InsertOne(ctx, emptyData)
		if err != nil {
			return models.UserData{}, err
		}
		return emptyData, nil
	} else if err != nil {
		return models.UserData{}, err
	}
	return userData, nil
}

func (s *UserDataService) UpdateUserData(userData *models.UserData) error {
	ctx := context.Background()

	filter := bson.M{"uid": userData.Uid}
	update := bson.M{
		"$set": bson.M{
			"items": userData.Items,
		},
	}

	opts := options.Update().SetUpsert(true)
	res, err := s.collection.UpdateOne(ctx, filter, update, opts)
	if err != nil {
		return err
	}
	if res.MatchedCount == 0 && res.UpsertedCount == 0 {
		return errors.New("user data not found")
	}
	return nil
}
