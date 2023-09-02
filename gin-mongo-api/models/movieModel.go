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


type Movie struct {
    Id          primitive.ObjectID      `json:"id,omitempty"           bson:"_id,omitempty"`
    // basic info
    Title       string                  `json:"title,omitempty"        bson:"title,omitempty"       tmdb:"title"        validate:"required"`
    Description string                  `json:"description,omitempty"  bson:"description,omitempty" tmdb:"overview"     validate:"required"`
    Date        string                  `json:"date,omitempty"         bson:"date,omitempty"        tmdb:"release_date" validate:"required"`
    Genre       []string                `json:"genre,omitempty"        bson:"genre,omitempty"`
    Info        string                  `json:"info,omitempty"         bson:"info,omitempty"        tmdb:"homepage"`
    Budget      int                     `json:"budget,omitempty"       bson:"budget,omitempty"      tmdb:"budget"`
    Length      int                     `json:"length,omitempty"       bson:"length,omitempty"      tmdb:"runtime"`
    Rating      string                  `json:"rating,omitempty"       bson:"rating,omitempty"`
    // pupularity
    Popularity  float64                 `json:"popularity,omitempty"   bson:"popularity,omitempty"  tmdb:"popularity"`
    VoteCount   int                     `json:"voteCount,omitempty"    bson:"voteCount, omitempty"  tmdb:"vote_count"`
    VoteRating  float64                 `json:"voteRating,omitempty"   bson:"voteRating,omitempty"  tmdb:"vote_average"`
    // related media
    Images      []string                `json:"images,omitempty"       bson:"images,omitempty"`
    // ids
    ExternalIds map[string]string       `json:"externalIds,omitempty"  bson:"externalIds,omitempty"`
    Platforms   []string                `json:"platforms,omitempty"    bson:"platforms,omitempty"`
}

var movieCollection *mongo.Collection = configs.GetCollection(configs.DB, "movie")

func (m *Movie) Collection() *mongo.Collection {return movieCollection}

func (m *Movie) Write() mongo.WriteModel {
    updateModel := mongo.NewUpdateOneModel()
    updateModel.SetFilter(bson.M{"title": m.Title, "date": m.Date}) 
    updateModel.SetUpdate(bson.D{{"$set", *m}})
    updateModel.SetUpsert(true)

    return updateModel
}


func FindMovie(filter bson.D) ([]Movie, error) {
    ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
    var movies []Movie
    defer cancel()

    results, err := movieCollection.Find(ctx, filter)
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
