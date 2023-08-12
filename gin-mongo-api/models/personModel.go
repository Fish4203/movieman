package models

import (
    "gin-mongo-api/configs"
    "go.mongodb.org/mongo-driver/mongo"
    "go.mongodb.org/mongo-driver/bson/primitive"
)
type Person struct {
    Id          primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
    Name        string             `json:"name,omitempty" validate:"required"`
    Description string             `json:"description,omitempty"`
    Date        string             `json:"date,omitempty"`
    Image       string             `json:"image,omitempty"`
    TMDB        uint               `json:"TMDB,omitempty"`
    IMDB        string             `json:"IMDB,omitempty"`
}

var PersonCollection *mongo.Collection = configs.GetCollection(configs.DB, "person")
