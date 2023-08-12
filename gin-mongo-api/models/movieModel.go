package models

import (
    "context"
    "gin-mongo-api/configs"
    "go.mongodb.org/mongo-driver/mongo"
    "github.com/go-playground/validator/v10"
    "time"
    "go.mongodb.org/mongo-driver/bson/primitive"
)

var validate = validator.New()


type Movie struct {
    Id          primitive.ObjectID `json:"id,omitempty"           bson:"_id,omitempty"`
    Title       string             `json:"title,omitempty"        validate:"required"`
    Description string             `json:"description,omitempty"  validate:"required"`
    Date        string             `json:"date,omitempty"         validate:"required"`
    Genre       []string           `json:"genre,omitempty"        bson:"genre,omitempty"`
    Info        string             `json:"info,omitempty"         bson:"info,omitempty"`
    Budget      int                `json:"budget,omitempty"       bson:"budget,omitempty"`
    Popularity  float64            `json:"popularity,omitempty"   bson:"popularity,omitempty"`
    VoteCount   int                `json:"voteCount,omitempty"    bson:"voteCount, omitempty"`
    VoteRating  float64            `json:"voteRating,omitempty"   bson:"voteRating,omitempty"`
    Rating      string             `json:"rating,omitempty"       bson:"rating,omitempty"`
    Length      uint               `json:"length,omitempty"       bson:"length,omitempty"`
    Image       string             `json:"image,omitempty"        bson:"image,omitempty"`
    TMDB        int                `json:"TMDB,omitempty"         bson:"TMDB,omitempty"`
    IMDB        string             `json:"IMDB,omitempty"         bson:"IMDB,omitempty"`
}

var MovieCollection *mongo.Collection = configs.GetCollection(configs.DB, "movie")


func (m *Movie) Save() error {
    ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
    defer cancel()

    if err := validate.Struct(m); err != nil {
        return err
    }

    _, err := MovieCollection.InsertOne(ctx, *m)

    return err
}
