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
    Title       string                  `json:"title"                  bson:"title,omitempty"       validate:"required"`
    Description string                  `json:"description"            bson:"description,omitempty" validate:"required"`
    Date        string                  `json:"date,omitempty"         bson:"date,omitempty"        validate:"required"`
    Genre       []string                `json:"genre,omitempty"        bson:"genre,omitempty"       `
    Info        string                  `json:"info,omitempty"         bson:"info,omitempty"        `
    Budget      int                     `json:"budget,omitempty"       bson:"budget,omitempty"      `
    Length      int                     `json:"length,omitempty"       bson:"length,omitempty"      `
    Rating      string                  `json:"rating,omitempty"       bson:"rating,omitempty"      `
    // pupularity
    Reviews     map[string]string       `json:"reviews,omitempty"      bson:"reviews,omitempty" `
    // related media
    Images      []string                `json:"images,omitempty"       bson:"images,omitempty"      `
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
