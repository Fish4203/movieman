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
    Id          primitive.ObjectID      `json:"id,omitempty"           bson:"_id,omitempty"         `
    // basic info
    Title       string                  `json:"title,omitempty"        bson:"title,omitempty"       tmdb:"title"        validate:"required"`
    Description string                  `json:"description,omitempty"  bson:"description,omitempty" tmdb:"overview,omitempty"     validate:"required"`
    Date        string                  `json:"date,omitempty"         bson:"date,omitempty"        tmdb:"release_date,omitempty" validate:"required"`
    Genre       []string                `json:"genre,omitempty"        bson:"genre,omitempty"       `
    Info        string                  `json:"info,omitempty"         bson:"info,omitempty"        tmdb:"homepage,omitempty"`
    Budget      int                     `json:"budget,omitempty"       bson:"budget,omitempty"      tmdb:"budget,omitempty"`
    Length      int                     `json:"length,omitempty"       bson:"length,omitempty"      tmdb:"runtime,omitempty"`
    Rating      string                  `json:"rating,omitempty"       bson:"rating,omitempty"      `
    // pupularity
    Popularity  float64                 `json:"popularity,omitempty"   bson:"popularity,omitempty"  tmdb:"popularity,omitempty"`
    VoteCount   int                     `json:"voteCount,omitempty"    bson:"voteCount, omitempty"  tmdb:"vote_count,omitempty"`
    VoteRating  float64                 `json:"voteRating,omitempty"   bson:"voteRating,omitempty"  tmdb:"vote_average,omitempty"`
    // related media
    Images      []string                `json:"images,omitempty"       bson:"images,omitempty"      tmdb:"poster_path,omitempty"`
    // ids
    ExternalIds map[string]string       `json:"externalIds,omitempty"  bson:"externalIds,omitempty" `
    Platforms   []string                `json:"platforms,omitempty"    bson:"platforms,omitempty"   `
}

var movieCollection *mongo.Collection = configs.GetCollection(configs.DB, "movie")

func WriteMovie(models []Movie) error {
    if len(models) == 0 {
        return nil
    }
    ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
    defer cancel()

    var writeObjs []mongo.WriteModel 
	for i := 0; i < len(models); i++ {
        updateModel := mongo.NewUpdateOneModel()
        updateModel.SetFilter(bson.M{"title": models[i].Title, "date": models[i].Date}) 
        updateModel.SetUpdate(bson.D{{"$set", models[i]}})
        updateModel.SetUpsert(true)

		writeObjs = append(writeObjs, updateModel)
	}
	
    _, err := movieCollection.BulkWrite(ctx, writeObjs)

    return err
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
