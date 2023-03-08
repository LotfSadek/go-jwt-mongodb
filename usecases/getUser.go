package usecases

import (
	"context"
	"golang-jwt/database"
	"golang-jwt/models"
	"time"

	"go.mongodb.org/mongo-driver/bson"
)

func (c *_us) GetUser(user_id string) (models.User, error) {
	var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
	var user models.User
	err := database.UserCollection.FindOne(ctx, bson.M{"user_id": user_id}).Decode(&user)
	defer cancel()
	if err != nil {
		return user, err
	}
	return user, nil
}
