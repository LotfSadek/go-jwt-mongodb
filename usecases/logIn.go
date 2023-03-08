package usecases

import (
	"errors"
	"fmt"
	helper "golang-jwt/helpers"
	"golang-jwt/models"
)

func (c *_us) Login(user *models.User) (*models.User, error) {

	foundUser := models.User{}
	// delegated to repo layer
	err := c.ar.Fetch(*user.Email, &foundUser)
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
	return &foundUser, nil

}
