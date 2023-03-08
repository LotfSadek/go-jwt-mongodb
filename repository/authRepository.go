package repository

import (
	"context"
	"golang-jwt/database"
	"golang-jwt/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type AuthRepo interface {
	Create(ctx context.Context, user models.User) (*mongo.InsertOneResult, error)
	Fetch(userEmail string, foundUser *models.User) error
}
type _authRepo struct {
}

func NewAuthRepo() AuthRepo {
	return &_authRepo{}
}

func (c *_authRepo) Create(ctx context.Context, user models.User) (*mongo.InsertOneResult, error) {
	resultInsertionNumber, insertErr := database.UserCollection.InsertOne(ctx, user)
	return resultInsertionNumber, insertErr
}

// no need to return a *user because we are working with pointers?
func (c *_authRepo) Fetch(userEmail string, foundUser *models.User) error {
	err := database.UserCollection.FindOne(context.TODO(), bson.M{"email": userEmail}).Decode(&foundUser)
	return err
}
