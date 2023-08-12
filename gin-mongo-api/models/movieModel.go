package models

import (
    "gin-mongo-api/configs"
    "go.mongodb.org/mongo-driver/mongo"
    "go.mongodb.org/mongo-driver/bson/primitive"
)
type Movie struct {
    Id          primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
    Title       string             `json:"title,omitempty" validate:"required"`
    Description string             `json:"description,omitempty" validate:"required"`
    Date        string             `json:"date,omitempty" validate:"required"`
    Genre       []string           `json:"genre,omitempty" validate:"required"`
    Info        string             `json:"info,omitempty"`
    Budget      uint               `json:"budget,omitempty"`
    Popularity  float64            `json:"popularity,omitempty"`
    VoteCount   uint               `json:"voteCount,omitempty"`
    VoteRating  float64            `json:"voteRating,omitempty"`
    Rating      string             `json:"rating,omitempty"`
    Length      uint               `json:"length,omitempty"`
    Image       string             `json:"image,omitempty"`
    TMDB        uint               `json:"TMDB,omitempty"`
    IMDB        string             `json:"IMDB,omitempty"`
}

var MovieCollection *mongo.Collection = configs.GetCollection(configs.DB, "movie")
