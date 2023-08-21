package models

import (
    "time"
    "context"
    "gin-mongo-api/configs"
    "github.com/go-playground/validator/v10"
    "go.mongodb.org/mongo-driver/mongo"
    "go.mongodb.org/mongo-driver/mongo/options"
    "go.mongodb.org/mongo-driver/bson"
    "go.mongodb.org/mongo-driver/bson/primitive"
)

var validate = validator.New()
var updateOpts = options.Update().SetUpsert(true)

type Movie struct {
    Id          primitive.ObjectID      `json:"id,omitempty"           bson:"_id,omitempty"`
    // basic info
    Title       string                  `json:"title,omitempty"        validate:"required"`
    Description string                  `json:"description,omitempty"  validate:"required"`
    Date        string                  `json:"date,omitempty"         validate:"required"`
    Genre       []string                `json:"genre,omitempty"        bson:"genre,omitempty"`
    Info        string                  `json:"info,omitempty"         bson:"info,omitempty"`
    Budget      int                     `json:"budget,omitempty"       bson:"budget,omitempty"`
    Length      int                     `json:"length,omitempty"       bson:"length,omitempty"`
    Rating      string                  `json:"rating,omitempty"       bson:"rating,omitempty"`
    // pupularity
    Popularity  float64                 `json:"popularity,omitempty"   bson:"popularity,omitempty"`
    VoteCount   int                     `json:"voteCount,omitempty"    bson:"voteCount, omitempty"`
    VoteRating  float64                 `json:"voteRating,omitempty"   bson:"voteRating,omitempty"`
    // related media
    Image       string                  `json:"image,omitempty"        bson:"image,omitempty"`
    // ids
    TMDB        int                     `json:"TMDB,omitempty"         bson:"TMDB,omitempty"`
    IMDB        string                  `json:"IMDB,omitempty"         bson:"IMDB,omitempty"`
    // adjacent media
    AdjShows    []primitive.ObjectID    `json:"adjShows,omitempty"     bson:"adjShows,omitempty"`
    AdjMovies   []primitive.ObjectID    `json:"adjMovies,omitempty"    bson:"adjMovies,omitempty"`
}

var MovieCollection *mongo.Collection = configs.GetCollection(configs.DB, "movie")


func (m *Movie) Save() error {
    ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
    defer cancel()

    if err := validate.Struct(m); err != nil {
        return err
    }

    filter := bson.M{"title": m.Title, "date": m.Date}
    _, err := MovieCollection.UpdateOne(ctx, filter, bson.D{{"$set", *m}}, updateOpts)

    err = MovieCollection.FindOne(ctx, filter).Decode(m)

    return err
}

func FindMovie(filter bson.D) ([]Movie, error) {
    ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
    var movies []Movie
    defer cancel()

    results, err := MovieCollection.Find(ctx, filter)
    if err != nil {
        return movies, err
    }

    //reading from the db in an optimal way
    defer results.Close(ctx)
    for results.Next(ctx) {
        var singleMovie Movie
        if err = results.Decode(&singleMovie); err != nil {
            return movies, err
        }
        movies = append(movies, singleMovie)
    }

    return movies, nil
}
