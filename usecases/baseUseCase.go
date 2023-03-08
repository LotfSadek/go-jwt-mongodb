package usecases

import (
	"golang-jwt/models"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type Base interface {
	SignUp(user *models.User) (*mongo.InsertOneResult, error)
	Login(user *models.User) (*models.User, error)
	GetUsers() ([]primitive.M, error)
	GetUser(user_id string) (models.User, error)
}
type _us struct {
}

func NewBaseRepo() Base {
	return &_us{}
}
