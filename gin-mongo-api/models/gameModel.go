package models

import (
	"context"
	"gin-mongo-api/configs"
	"time"

	// "github.com/go-playground/validator/v10"
	"go.mongodb.org/mongo-driver/mongo"
	// "go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Game struct {
	Id primitive.ObjectID `json:"id,omitempty"           bson:"_id,omitempty"`
	// basic info
	Title       string   `json:"title"        bson:"title,omitempty"		validate:"required"`
	Description string   `json:"description"  bson:"description,omitempty" validate:"required"`
	Date        string   `json:"date,omitempty"         bson:"date,omitempty" 		validate:"required"`
	Genre       []string `json:"genre,omitempty"        bson:"genre,omitempty"`
	Info        string   `json:"info,omitempty"         bson:"info,omitempty"`
	// requirements
	MinReq string `json:"minReq,omitempty"       bson:"minReq,omitempty"`
	RecReq string `json:"recReq,omitempty"       bson:"recReq,omitempty"`
	// pupularity
	Reviews map[string]string `json:"reviews,omitempty"      bson:"reviews,omitempty" `
	// related media
	Image []string `json:"image,omitempty"        bson:"image,omitempty"`
	// ids
	ExternalIds map[string]string `json:"externalIds,omitempty"  bson:"externalIds,omitempty"`
	Platforms   []string          `json:"platforms,omitempty"    bson:"platforms,omitempty"`
}

var gameCollection *mongo.Collection = configs.GetCollection(configs.DB, "game")

func WriteGame(models []Game) error {
	if len(models) == 0 {
		return nil
	}
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var writeObjs []mongo.WriteModel
	for i := 0; i < len(models); i++ {
		updateModel := mongo.NewUpdateOneModel()
		updateModel.SetFilter(bson.M{"title": models[i].Title, "date": models[i].Date})
		updateModel.SetUpdate(bson.D{{"$set", models[i]}})
		updateModel.SetUpsert(true)

		writeObjs = append(writeObjs, updateModel)
	}

	_, err := gameCollection.BulkWrite(ctx, writeObjs)

	return err
}

func FindGame(filter bson.D) ([]Game, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	var games []Game
	defer cancel()

	results, err := gameCollection.Find(ctx, filter)
	if err != nil {
		return games, err
	}

	//reading from the db in an optimal way
	defer results.Close(ctx)
	for results.Next(ctx) {
		var singleGame Game
		if err = results.Decode(&singleGame); err != nil {
			return games, err
		}
		games = append(games, singleGame)
	}

	return games, nil
}
