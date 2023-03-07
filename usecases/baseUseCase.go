package usecases

import (
	"golang-jwt/models"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type Base interface {
	SignUp(user *models.User) (*mongo.InsertOneResult, error)
	Login(user *models.User) (*models.User, error)
	GetUsers(allusers []primitive.M) (primitive.M, error)
	GetUser() gin.HandlerFunc
}
type _us struct {
}

func NewBaseRepo() Base {
	return &_us{}
}
