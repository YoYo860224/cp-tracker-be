package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct {
	Uid         primitive.ObjectID `bson:"_id,omitempty" json:"uid"`
	Email       string             `bson:"email" json:"email"`
	Password    string             `bson:"password" json:"password"`
	DisplayName string             `bson:"displayName" json:"displayName"`
}

type UserDto struct {
	Uid         primitive.ObjectID `bson:"_id,omitempty" json:"uid"`
	Email       string             `bson:"email" json:"email"`
	DisplayName string             `bson:"displayName" json:"displayName"`
	Token       string             `bson:"token" json:"token,omitempty"`
}

func (u *User) ToDto() UserDto {
	return UserDto{
		Uid:         u.Uid,
		Email:       u.Email,
		DisplayName: u.DisplayName,
	}
}
