package models

import (
    "time"
    "context"
    "gin-mongo-api/configs"
    "go.mongodb.org/mongo-driver/mongo"
    "go.mongodb.org/mongo-driver/bson"
    "go.mongodb.org/mongo-driver/bson/primitive"
)

type Watch struct {
    Movie       primitive.ObjectID `json:"movie,omitempty"      bson:"movie,omitempty"`
    Show        primitive.ObjectID `json:"show,omitempty"       bson:"show,omitempty"`
    Book        primitive.ObjectID `json:"gook,omitempty"       bson:"gook,omitempty"`
    Game        primitive.ObjectID `json:"game,omitempty"       bson:"game,omitempty"`
    User        primitive.ObjectID `json:"user,omitempty"       bson:"user,omitempty"       validate:"required"`
    Watched     uint               `json:"watched,omitempty"    bson:"watched,omitempty"    validate:"required"`
    UserRating  float64            `json:"userRating,omitempty" bson:"userRating,omitempty"`
    Notes       string             `json:"notes,omitempty"      bson:"notes,omitempty"`
}

var watchCollection *mongo.Collection = configs.GetCollection(configs.DB, "watch")

func (o *Watch) Collection() *mongo.Collection {return watchCollection}

func (o *Watch) Write() mongo.WriteModel {
    updateModel := mongo.NewUpdateOneModel()
    updateModel.SetFilter(bson.M{"user": o.User, "$or": bson.A{
        bson.M{"movie": o.Movie},
        bson.M{"show": o.Game},
        bson.M{"book": o.Book},
        bson.M{"game": o.Game},
    }}) 
    updateModel.SetUpdate(bson.D{{"$set", *o}})
    updateModel.SetUpsert(true)

    return updateModel
}


func FindWatch(filter bson.D) ([]Watch, error) {
    ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
    var watchs []Watch
    defer cancel()

    results, err := watchCollection.Find(ctx, filter)
    if err != nil {
        return watchs, err
    }

    //reading from the db in an optimal way
    defer results.Close(ctx)
    for results.Next(ctx) {
        var singleWatch Watch
        if err = results.Decode(&singleWatch); err != nil {
            return watchs, err
        }
        watchs = append(watchs, singleWatch)
    }

    return watchs, nil
}
