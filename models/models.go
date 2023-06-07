package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

var (
	SettingsPublic  = "PUBLIC"
	SettingsPrivate = "PRIVATE"
)

type Contact struct {
	Email      string `bson:"email"`
	IsFavorite bool   `bson:"is_favorite"`
	IsBlocked  bool   `bson:"is_blocked"`
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
