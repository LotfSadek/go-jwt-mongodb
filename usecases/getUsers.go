package usecases

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (c *_us) GetUsers() ([]primitive.M, error) {

	var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
	// to be delegated to repo layer
	allusers, err := c.ur.Fetch(ctx, cancel)
	if err != nil {
		return nil, err
	}
	return allusers, nil
}
