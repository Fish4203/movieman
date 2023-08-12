package models

import (
    "gin-mongo-api/configs"
    "go.mongodb.org/mongo-driver/mongo"
    "go.mongodb.org/mongo-driver/bson/primitive"
)
type Show struct {
    Id          primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
    Title       string             `json:"title,omitempty" validate:"required"`
    Description string             `json:"description,omitempty" validate:"required"`
    Date        string             `json:"date,omitempty" validate:"required"`
    Seasons     uint               `json:"seasons,omitempty" validate:"required"`
    Genre       []string           `json:"genre,omitempty" validate:"required"`
    Info        string             `json:"info,omitempty"`
    Popularity  float64            `json:"popularity,omitempty"`
    VoteCount   uint               `json:"voteCount,omitempty"`
    VoteRating  float64            `json:"voteRating,omitempty"`
    Rating      string             `json:"rating,omitempty"`
    Image       string             `json:"image,omitempty"`
    TMDB        uint               `json:"TMDB,omitempty"`
    IMDB        string             `json:"IMDB,omitempty"`
}


type ShowSeason struct {
    ShowId      primitive.ObjectID `json:"showId,omitempty" validate:"required"`
    SeasonID    uint               `json:"seasonId,omitempty" validate:"required"`
    Epesodes    uint               `json:"epesodes,omitempty" validate:"required"`
    Date        string             `json:"date,omitempty"`
    Popularity  float64            `json:"popularity,omitempty"`
    VoteCount   uint               `json:"voteCount,omitempty"`
    VoteRating  float64            `json:"voteRating,omitempty"`
    Rating      string             `json:"rating,omitempty"`
    Image       string             `json:"image,omitempty"`
}


type ShowEpisode struct {
    ShowId      primitive.ObjectID `json:"showId,omitempty" validate:"required"`
    SeasonID    uint               `json:"seasonId,omitempty" validate:"required"`
    EpesodeID   uint               `json:"epesodeId,omitempty" validate:"required"`
    Title       string             `json:"title,omitempty" validate:"required"`
    Date        string             `json:"date,omitempty"`
    Popularity  float64            `json:"popularity,omitempty"`
    VoteCount   uint               `json:"voteCount,omitempty"`
    VoteRating  float64            `json:"voteRating,omitempty"`
    Image       string             `json:"image,omitempty"`
}

var ShowCollection *mongo.Collection = configs.GetCollection(configs.DB, "show")
var ShowSeasonCollection *mongo.Collection = configs.GetCollection(configs.DB, "showSeason")
var ShowEpisodeCollection *mongo.Collection = configs.GetCollection(configs.DB, "showEpisode")
