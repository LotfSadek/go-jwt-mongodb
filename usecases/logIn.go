package usecases

import (

	//"golang-jwt/controllers"
	"context"
	"errors"
	"fmt"
	"golang-jwt/database"
	helper "golang-jwt/helpers"
	"golang-jwt/models"

	"go.mongodb.org/mongo-driver/bson"
)

func (c *_us) Login(user *models.User) (*models.User, error) {

	var foundUser *models.User
	err := database.UserCollection.FindOne(context.TODO(), bson.M{"email": user.Email}).Decode(&foundUser)
	if err != nil {
		return nil, err
	}
	passwordIsValid, msg := helper.VerifyPassword(*user.Password, *foundUser.Password)
	fmt.Println(msg)
	if !passwordIsValid {
		return nil, errors.New("pass not valid")
	}
	if foundUser.Email == nil {
		return nil, errors.New("email not valid")
	}
	token, refreshToken, _ := helper.GenerateAllTokens(*foundUser.Email, *foundUser.First_name, *foundUser.Last_name, *foundUser.User_type, foundUser.User_id)
	helper.UpdateAllTokens(token, refreshToken, foundUser.User_id)
	return foundUser, nil

}
