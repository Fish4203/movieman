package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Movie struct {
    Id          primitive.ObjectID `json:"id,omitempty"`
    Title       string             `json:"title,omitempty" validate:"required"`
    Description string             `json:"description"`
    Date        string             `json:"date"`
    Budget      uint               `json:"budget"`
    Popularity  float64            `json:"popularity"`
    VoteCount   uint               `json:"voteCount"`
    VoteRating  float64            `json:"voteRating"`
    Rating      string             `json:"rating"`
    Length      uint               `json:"length"`
    Image       string             `json:"role"`
    TMDB        uint               `json:"role"`
    IMDB        string             `json:"role"`
}
