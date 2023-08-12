package models

import (
    "gin-mongo-api/configs"
    "go.mongodb.org/mongo-driver/mongo"
    "go.mongodb.org/mongo-driver/bson/primitive"
)

type Watch struct {
    Object      primitive.ObjectID `json:"object,omitempty" validate:"required"`
    User        primitive.ObjectID `json:"user,omitempty" validate:"required"`
    Type        string             `json:"type,omitempty" validate:"required"`
    Watched     uint               `json:"watched,omitempty" validate:"required"`
    UserRating  float64            `json:"userRating,omitempty"`
    Notes       string             `json:"notes,omitempty"`
}

var WatchCollection *mongo.Collection = configs.GetCollection(configs.DB, "watch")
