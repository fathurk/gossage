package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// User (collection)

// {
// 	"_id"	: "ObjectId()"
// 	"email"	: "example@example.com",
// 	"phone_number"	: "08238374928",
// 	"password"	: "hashedpassword",
// 	"profile_picture"	: "http://example.com",
// 	"contact": [
// 		{
// 			"_id"	: "ObjectId()"
// 			"phone_number"	: "08238374928",
// 			"profile_picture"	: "http://example.com",
// 		},
// 		{
// 			"_id"	: "ObjectId()"
// 			"phone_number"	: "08238374928",
// 			"profile_picture"	: null,
// 		},
// 	]
// }

var (
	SettingsPublic  = "PUBLIC"
	SettingsPrivate = "PRIVATE"
)

type Contact struct {
	PhoneNumber    string `bson:"phone_number"`
	ProfilePicture string `bson:"profile_picture"`
}

type UserSetting struct {
	OnlineStatus       string `bson:"online_status"`
	ProfilePictureSeen string `bson:"profile_picture_seen"`
	ChatByOther        string `bson:"chat_by_other"`
}
type User struct {
	ID             primitive.ObjectID `bson:"_id"`
	CreatedAt      time.Time          `bson:"created_at"`
	LastOnline     time.Time          `bson:"last_online"`
	UserName       string             `bson:"user_name"`
	Bio            string             `bson:"bio"`
	Email          string             `bson:"email"`
	PhoneNumber    string             `bson:"phone_number"`
	Password       string             `bson:"password"`
	ProfilePicture string             `bson:"profile_picture"`
	Contact        []Contact          `bson:"contact"`
	UserSetting    UserSetting        `bson:"user_setting"`
}
