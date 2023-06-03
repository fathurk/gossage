package handlers

import (
	"context"
	"fmt"
	"gossage/models"
	"gossage/utils"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type CreateUserRequestBody struct {
	UserName    string `json:"user_name" binding:"required"`
	Email       string `json:"email" binding:"required"`
	PhoneNumber string `json:"phone_number" binding:"required,lte=15"`
	Password    string `json:"password" binding:"required"`
	Bio         string `json:"bio" `
}

var (
	DefaultUserSettings = &models.UserSetting{OnlineStatus: "PUBLIC", ProfilePictureSeen: "PUBLIC", ChatByOther: "TRUE"}
)

func CreateUser(c *gin.Context) {

	var requestBody CreateUserRequestBody

	if err := c.BindJSON(&requestBody); err != nil {
		Responser(c, 400, err.Error())
		return
	}

	client, err := utils.ConnectDb()
	defer utils.DisconnectDb(client)

	if err != nil {
		Responser(c, 500, err.Error())
		return
	}

	hashed, err := utils.HashPassword(requestBody.Password)

	if err != nil {
		Responser(c, 500, err.Error())
	}

	now := time.Now()

	userPayload := &models.User{
		ID:          primitive.NewObjectID(),
		CreatedAt:   now,
		LastOnline:  now,
		UserName:    requestBody.UserName,
		Email:       requestBody.Email,
		PhoneNumber: requestBody.PhoneNumber,
		Password:    hashed,
		Bio:         requestBody.Bio,
		UserSetting: *DefaultUserSettings,
	}

	result, err := client.Database("gossage").Collection("user").InsertOne(context.TODO(), userPayload)

	fmt.Printf("Inserted document with _id: %v\n", result.InsertedID)

	if err != nil {
		Responser(c, 500, err.Error())
		return
	}

	Responser(c, 201, "User with esuccessfully created")
}
