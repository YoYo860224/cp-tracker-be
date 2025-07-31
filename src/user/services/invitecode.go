package services

import (
	"cp_tracker/user/models"

	"context"
	"errors"
	"regexp"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type InviteCodeService struct {
	db         *mongo.Database
	collection *mongo.Collection
}

func NewInviteCodeService(db *mongo.Database) (*InviteCodeService, error) {
	if db == nil {
		return nil, mongo.ErrNilDocument
	}

	return &InviteCodeService{
		db:         db,
		collection: db.Collection("invite_code"),
	}, nil
}

func (s *InviteCodeService) ValidateInviteCode(inviteCode string, email string) error {
	ctx := context.Background()

	var entry models.InviteCode
	err := s.collection.FindOne(ctx, bson.M{"code": inviteCode}).Decode(&entry)
	if err != nil {
		return errors.New("no invite code found")
	}

	if matched, err := regexp.MatchString(entry.Email, email); err != nil {
		return errors.New("error matching email with invite code")
	} else if !matched {
		return errors.New("email does not match invite code")
	}

	return nil
}
