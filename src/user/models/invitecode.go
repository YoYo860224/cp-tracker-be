package models

type InviteCode struct {
	Code  string `bson:"code" json:"code"`
	Email string `bson:"email" json:"email"`
}
