package models

import (
    "time"
    "context"
    "gin-mongo-api/configs"
    "go.mongodb.org/mongo-driver/mongo"
    "go.mongodb.org/mongo-driver/bson"
    "go.mongodb.org/mongo-driver/bson/primitive"
)
type Show struct {
    Id          primitive.ObjectID      `json:"id,omitempty"           bson:"_id,omitempty"         `
    // basic info
    Title       string                  `json:"title"                  bson:"title,omitempty"       validate:"required"`
    Description string                  `json:"description"            bson:"description,omitempty" validate:"required"`
    Date        string                  `json:"date,omitempty"         bson:"date,omitempty"        validate:"required"`
    Seasons     int                     `json:"seasons"                bson:"seasons,omitempty"     `
    Genre       []string                `json:"genre,omitempty"        bson:"genre,omitempty"       `      
    Info        string                  `json:"info,omitempty"         bson:"info,omitempty"        `
    Rating      string                  `json:"rating,omitempty"       bson:"rating,omitempty"      `
    // other media
    Images      []string                `json:"images,omitempty"       bson:"images,omitempty"      `
    // external ids
    Platforms   []string                `json:"platforms,omitempty"    bson:"platforms,omitempty"   `
    ExternalIds map[string]string       `json:"externalIds,omitempty"  bson:"externalIds,omitempty" `
    Reviews     map[string]string       `json:"reviews,omitempty"      bson:"reviews,omitempty" `
}


type ShowSeason struct {
    /// ids 
    ShowId      primitive.ObjectID      `json:"showId"                 bson:"showId,omitempty"      validate:"required"`
    SeasonID    int                     `json:"seasonId"               bson:"seasonId,omitempty"    validate:"required"`
    // basic info 
    Episodes    int                     `json:"episodes"               bson:"episodes,omitempty"    validate:"required"`
    Description string                  `json:"description"            bson:"description,omitempty" validate:"required"`
    Date        string                  `json:"date,omitempty"         bson:"date,omitempty"        validate:"required"`
    // external media 
    Images      []string                `json:"images,omitempty"       bson:"images,omitempty"`
}


type ShowEpisode struct {
    // ids
    ShowId      primitive.ObjectID      `json:"showId"                 bson:"showId,omitempty"      validate:"required"`
    SeasonID    int                     `json:"seasonId"               bson:"seasonId,omitempty"    validate:"required"`
    EpisodeID   int                     `json:"episodeId"              bson:"episodeId,omitempty"   validate:"required"`
    // basic info 
    Title       string                  `json:"title"                  bson:"title,omitempty"       validate:"required"`
    Description string                  `json:"description"            bson:"description,omitempty" validate:"required"`
    Date        string                  `json:"date,omitempty"         bson:"date,omitempty"        `
    // ratings 
    Reviews     map[string]string       `json:"reviews,omitempty"      bson:"reviews,omitempty" `
    // external media
    Images      []string                `json:"images,omitempty"       bson:"images,omitempty"`
}

var showCollection *mongo.Collection = configs.GetCollection(configs.DB, "show")
var showSeasonCollection *mongo.Collection = configs.GetCollection(configs.DB, "showSeason")
var showEpisodeCollection *mongo.Collection = configs.GetCollection(configs.DB, "showEpisode")

func WriteShow(models []Show) error {
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
	
    _, err := showCollection.BulkWrite(ctx, writeObjs)

    return err
}

func WriteShowSeason(models []ShowSeason) error {
    if len(models) == 0 {
        return nil
    }
    ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
    defer cancel()

    var writeObjs []mongo.WriteModel 
	for i := 0; i < len(models); i++ {
        updateModel := mongo.NewUpdateOneModel()
        updateModel.SetFilter(bson.M{"seasonId": models[i].SeasonID, "showId": models[i].ShowId}) 
        updateModel.SetUpdate(bson.D{{"$set", models[i]}})
        updateModel.SetUpsert(true)

		writeObjs = append(writeObjs, updateModel)
	}
	
    _, err := showSeasonCollection.BulkWrite(ctx, writeObjs)

    return err
}

func WriteShowEpisode(models []ShowEpisode) error {
    if len(models) == 0 {
        return nil
    }
    ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
    defer cancel()

    var writeObjs []mongo.WriteModel 
	for i := 0; i < len(models); i++ {
        updateModel := mongo.NewUpdateOneModel()
        updateModel.SetFilter(bson.M{"seasonId": models[i].SeasonID, "showId": models[i].ShowId, "episodeId": models[i].EpisodeID}) 
        updateModel.SetUpdate(bson.D{{"$set", models[i]}})
        updateModel.SetUpsert(true)

		writeObjs = append(writeObjs, updateModel)
	}
	
    _, err := showEpisodeCollection.BulkWrite(ctx, writeObjs)

    return err
}


func FindShow(filter bson.D) ([]Show, error) {
    ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
    var shows []Show
    defer cancel()

    results, err := showCollection.Find(ctx, filter)
    if err != nil {
        return shows, err
    }

    //reading from the db in an optimal way
    defer results.Close(ctx)
    for results.Next(ctx) {
        var singleShow Show
        if err = results.Decode(&singleShow); err != nil {
            return shows, err
        }
        shows = append(shows, singleShow)
    }

    return shows, nil
}

func FindShowSeason(filter bson.D) ([]ShowSeason, error) {
    ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
    var shows []ShowSeason
    defer cancel()

    results, err := showSeasonCollection.Find(ctx, filter)
    if err != nil {
        return shows, err
    }

    //reading from the db in an optimal way
    defer results.Close(ctx)
    for results.Next(ctx) {
        var singleShow ShowSeason
        if err = results.Decode(&singleShow); err != nil {
            return shows, err
        }
        shows = append(shows, singleShow)
    }

    return shows, nil
}

func FindShowEpisode(filter bson.D) ([]ShowEpisode, error) {
    ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
    var shows []ShowEpisode
    defer cancel()

    results, err := showEpisodeCollection.Find(ctx, filter)
    if err != nil {
        return shows, err
    }

    //reading from the db in an optimal way
    defer results.Close(ctx)
    for results.Next(ctx) {
        var singleShow ShowEpisode
        if err = results.Decode(&singleShow); err != nil {
            return shows, err
        }
        shows = append(shows, singleShow)
    }

    return shows, nil
}
