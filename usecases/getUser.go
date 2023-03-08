package usecases

import (
	"context"
	"golang-jwt/models"
	"time"
)

func (c *_us) GetUser(user_id string) (models.User, error) {
	var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
	user := models.User{}
	// to be delegated to the repo layer
	err := c.ur.FetchById(ctx, cancel, user_id, &user)
	if err != nil {
		return user, err
	}
	return user, nil
}
