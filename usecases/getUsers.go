package usecases

import (
	"context"
	"errors"
	"golang-jwt/database"
	"log"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func (c *_us) GetUsers(allusers []primitive.M) (primitive.M, error) {

	var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
	ginCon := gin.Context{}
	recordPerPage, err := strconv.Atoi(ginCon.Query("recordPerPage"))
	if err != nil || recordPerPage < 1 {
		recordPerPage = 10
	}
	page, err1 := strconv.Atoi(ginCon.Query("page"))
	if err1 != nil || page < 1 {
		page = 1
	}

	startIndex := (page - 1) * recordPerPage
	startIndex, err = strconv.Atoi(ginCon.Query("startIndex"))

	matchStage := bson.D{{"$match", bson.D{{}}}}
	groupStage := bson.D{{"$group", bson.D{
		{"_id", bson.D{{"_id", "null"}}},
		{"total_count", bson.D{{"$sum", 1}}},
		{"data", bson.D{{"$push", "$$ROOT"}}}}}}
	projectStage := bson.D{
		{"$project", bson.D{
			{"_id", 0},
			{"total_count", 1},
			{"user_items", bson.D{{"$slice", []interface{}{"$data", startIndex, recordPerPage}}}}}}}
	result, err := database.UserCollection.Aggregate(ctx, mongo.Pipeline{
		matchStage, groupStage, projectStage})
	defer cancel()
	if err != nil {
		return nil, errors.New("error occured while listing user items")
	}
	//var allusers []bson.M
	if err = result.All(ctx, &allusers); err != nil {
		log.Fatal(err)
	}
	return allusers[0], nil
}
