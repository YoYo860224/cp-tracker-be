package models

type UserData struct {
	Uid   string                 `bson:"uid" json:"uid"`
	Items map[string]interface{} `bson:"items" json:"items"`
}
