package handlers

import (
	"context"
	"fmt"
	"gossage/models"
	"gossage/utils"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type AddContactRequestBody struct {
	IdRequester    string `json:"id_requester"`
	EmailRequester string `json:"email_requester"`
	EmailTarget    string `json:"email_target"`
}

func AddContact(c *gin.Context) {
	var requestBody AddContactRequestBody

	if err := c.BindJSON(&requestBody); err != nil {
		Responser(c, 400, err.Error(), nil)
		return
	}

	client, err := utils.ConnectDb()

	if err != nil {
		Responser(c, 500, err.Error(), nil)
		return
	}

	defer utils.DisconnectDb(client)

	// Find user target
	var findTarget *models.User

	fiterTarget := bson.M{"email": requestBody.EmailTarget, "user_setting.chat_by_other": models.SettingsPublic}
	err = client.Database("gossage").Collection("user").FindOne(context.TODO(), fiterTarget).Decode(&findTarget)

	if err != nil {
		if err == mongo.ErrNoDocuments {
			Responser(c, 404, err.Error(), nil)
			return
		}
		Responser(c, 500, err.Error(), nil)
		return
	}

	// Check are user target private or not
	if findTarget.UserSetting.ChatByOther == models.SettingsPrivate {
		Responser(c, 404, err.Error(), nil)
		return
	}

	// Check user target are requester blocked or not
	for _, contact := range findTarget.Contact {
		if contact.Email == requestBody.EmailRequester && contact.IsBlocked {
			Responser(c, 404, err.Error(), nil)
			return
		}
	}

	// Parse ID
	objectId, err := primitive.ObjectIDFromHex(requestBody.IdRequester)

	if err != nil {
		Responser(c, 400, err.Error(), nil)
		return
	}

	var findRequester *models.User

	filterRequester := bson.M{"_id": objectId}
	err = client.Database("gossage").Collection("user").FindOne(context.TODO(), filterRequester).Decode(&findRequester)

	if err != nil {
		if err == mongo.ErrNoDocuments {
			Responser(c, 404, err.Error(), nil)
			return
		}
		Responser(c, 500, err.Error(), nil)
		return
	}

	for _, contact := range findTarget.Contact {
		if contact.Email == requestBody.EmailTarget {
			Responser(c, 200, "Add contact successfully", nil)
			return
		}
	}

	newContact := &models.Contact{
		Email:      requestBody.EmailTarget,
		IsFavorite: false,
		IsBlocked:  false,
	}

	updatePayload := bson.M{
		"$push": bson.M{"contact": newContact}}

	result, err := client.Database("gossage").Collection("user").UpdateOne(context.TODO(), filterRequester, updatePayload)

	if err != nil {
		Responser(c, 500, err.Error(), nil)
		return
	}

	fmt.Printf("Inserted document with _id: %v\n", result.UpsertedID)

	Responser(c, 200, "Add contact successfully", nil)

}
