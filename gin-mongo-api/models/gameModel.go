package models

import (
    "time"
    "context"
    "gin-mongo-api/configs"
    // "github.com/go-playground/validator/v10"
    "go.mongodb.org/mongo-driver/mongo"
    // "go.mongodb.org/mongo-driver/mongo/options"
    "go.mongodb.org/mongo-driver/bson"
    "go.mongodb.org/mongo-driver/bson/primitive"
)


type Game struct {
    Id          primitive.ObjectID      `json:"id,omitempty"           bson:"_id,omitempty"`
    // basic info
    Title       string                  `json:"title,omitempty"        bson:"title,omitempty"		validate:"required"`
    Description string                  `json:"description,omitempty"  bson:"description,omitempty" validate:"required"`
    Date        string                  `json:"date,omitempty"         bson:"date,omitempty" 		validate:"required"`
    Genre       []string                `json:"genre,omitempty"        bson:"genre,omitempty"`
    Info        string                  `json:"info,omitempty"         bson:"info,omitempty"`
    Length      int                     `json:"length,omitempty"       bson:"length,omitempty"`
    // requirements
    MinReq      string                  `json:"minReq,omitempty"       bson:"minReq,omitempty"`
    RecReq      string                  `json:"recReq,omitempty"       bson:"recReq,omitempty"`
	// pupularity
    Popularity  float64                 `json:"popularity,omitempty"   bson:"popularity,omitempty"`
    VoteCount   int                     `json:"voteCount,omitempty"    bson:"voteCount, omitempty"`
    VoteRating  float64                 `json:"voteRating,omitempty"   bson:"voteRating,omitempty"`
    // related media
    Image       []string                `json:"image,omitempty"        bson:"image,omitempty"`
    // ids
    ExternalIds map[string]string       `json:"externalIds,omitempty"  bson:"externalIds,omitempty"`
    Platforms   []string                `json:"platforms,omitempty"    bson:"platforms,omitempty"`
}

var gameCollection *mongo.Collection = configs.GetCollection(configs.DB, "game")

func (o *Game) Collection() *mongo.Collection {return gameCollection}

func (o *Game) Write() mongo.WriteModel {
    updateModel := mongo.NewUpdateOneModel()
    updateModel.SetFilter(bson.M{"title": o.Title, "date": o.Date}) 
    updateModel.SetUpdate(bson.D{{"$set", *o}})
    updateModel.SetUpsert(true)

    return updateModel
}


func Findgame(filter bson.D) ([]Game, error) {
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
