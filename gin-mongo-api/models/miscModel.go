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
    Object      primitive.ObjectID `json:"object,omitempty"     validate:"required"`
    User        primitive.ObjectID `json:"user,omitempty"       validate:"required"`
    Type        string             `json:"type,omitempty"       validate:"required"`
    Watched     uint               `json:"watched,omitempty"    validate:"required"`
    UserRating  float64            `json:"userRating,omitempty" bson:"userRating,omitempty"`
    Notes       string             `json:"notes,omitempty"      bson:"notes,omitempty"`
}

var WatchCollection *mongo.Collection = configs.GetCollection(configs.DB, "watch")


func (w *Watch) Save() error {
    ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
    defer cancel()

    if err := validate.Struct(w); err != nil {
        return err
    }

    _, err := WatchCollection.UpdateOne(ctx, bson.M{"user": w.User, "object": w.Object, "type": w.Type}, *w)

    return err
}
