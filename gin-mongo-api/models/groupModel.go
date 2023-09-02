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


type Group struct {
    Id          primitive.ObjectID      `json:"id,omitempty"           bson:"_id,omitempty"`
    // basic info
    Title       string                  `json:"title,omitempty"        bson:"title,omitempty"       validate:"required"`
    Description string                  `json:"description,omitempty"  bson:"description,omitempty" validate:"required"`
    Genre       []string                `json:"genre,omitempty"        bson:"genre,omitempty"`
    Info        string                  `json:"info,omitempty"         bson:"info,omitempty"`
    // pupularity
    Popularity  float64                 `json:"popularity,omitempty"   bson:"popularity,omitempty"`
    VoteCount   int                     `json:"voteCount,omitempty"    bson:"voteCount, omitempty"`
    VoteRating  float64                 `json:"voteRating,omitempty"   bson:"voteRating,omitempty"`
	// media
    Movies      []primitive.ObjectID    `json:"movies,omitempty"       bson:"movies,omitempty"`
    Shows       []primitive.ObjectID    `json:"shows,omitempty"        bson:"shows,omitempty"`
    Books       []primitive.ObjectID    `json:"books,omitempty"        bson:"books,omitempty"`
    Games       []primitive.ObjectID    `json:"games,omitempty"        bson:"games,omitempty"`
    // related media
    Image       []string                `json:"image,omitempty"        bson:"image,omitempty"`
    // ids
    ExternalIds map[string]string       `json:"externalIds,omitempty"  bson:"externalIds,omitempty"`
}

var groupCollection *mongo.Collection = configs.GetCollection(configs.DB, "group")

func (o *Group) Collection() *mongo.Collection {return groupCollection}

func (o *Group) Write() mongo.WriteModel {
    updateModel := mongo.NewUpdateOneModel()
    updateModel.SetFilter(bson.M{"title": o.Title}) 
    updateModel.SetUpdate(bson.D{{"$set", *o}})
    updateModel.SetUpsert(true)

    return updateModel
}

func FindGroup(filter bson.D) ([]Group, error) {
    ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
    var objects []Group
    defer cancel()

    results, err := groupCollection.Find(ctx, filter)
    if err != nil {
        return objects, err
    }

    //reading from the db in an optimal way
    defer results.Close(ctx)
    for results.Next(ctx) {
        var singleGroup Group
        if err = results.Decode(&singleGroup); err != nil {
            return objects, err
        }
        objects = append(objects, singleGroup)
    }

    return objects, nil
}
