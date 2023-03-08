package usecases

import (
	"context"
	"golang-jwt/database"
	helper "golang-jwt/helpers"
	"golang-jwt/models"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func (c *_us) SignUp(user *models.User) (*mongo.InsertOneResult, error) {
	var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Minute)
	validationErr := helper.Validate.Struct(user)
	if validationErr != nil {
		return nil, validationErr
	}
	count, err := database.UserCollection.CountDocuments(ctx, bson.M{"email": user.Email})
	defer cancel()
	if err != nil {
		log.Panic(err)
	}
	password := helper.HashPassword(*user.Password)
	user.Password = &password
	count, err = database.UserCollection.CountDocuments(ctx, bson.M{"phone": user.Phone})
	defer cancel()
	if err != nil {
		return nil, err
	}
	if count > 0 {
		return nil, err
	}
	user.Created_at, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
	user.Updated_at, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
	user.ID = primitive.NewObjectID()
	user.User_id = user.ID.Hex()
	token, refreshToken, _ := helper.GenerateAllTokens(*user.Email, *user.First_name, *user.Last_name, *user.User_type, *&user.User_id)
	user.Token = &token
	user.Refresh_token = &refreshToken
	// to be delegated to the repo layer
	resultInsertionNumber, insertErr := c.ar.Create(ctx, *user)
	if insertErr != nil {
		return nil, insertErr
	}
	defer cancel()
	return resultInsertionNumber, nil
}
