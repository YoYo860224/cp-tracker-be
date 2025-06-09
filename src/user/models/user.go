package models

type User struct {
	Uid         string `bson:"uid" json:"uid"`
	Username    string `bson:"username" json:"username"`
	Email       string `bson:"email" json:"email"`
	Password    string `bson:"password" json:"password"`
	DisplayName string `bson:"displayName" json:"displayName"`
	Role        string `bson:"role" json:"role"`
}
