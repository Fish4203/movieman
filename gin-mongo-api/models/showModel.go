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
    Id          primitive.ObjectID      `json:"id,omitempty"           bson:"_id,omitempty"`
    // basic info
    Title       string                  `json:"title,omitempty"        validate:"required"`
    Description string                  `json:"description,omitempty"  validate:"required"`
    Date        string                  `json:"date,omitempty"         validate:"required"`
    Seasons     int                     `json:"seasons,omitempty"      bson:"seasons,omitempty"`
    Genre       []string                `json:"genre,omitempty"        bson:"genre,omitempty"`
    Info        string                  `json:"info,omitempty"         bson:"info,omitempty"`
    // reviews
    Popularity  float64                 `json:"popularity,omitempty"   bson:"popularity,omitempty"`
    VoteCount   int                     `json:"voteCount,omitempty"    bson:"voteCount,omitempty"`
    VoteRating  float64                 `json:"voteRating,omitempty"   bson:"voteRating,omitempty"`
    Rating      string                  `json:"rating,omitempty"       bson:"rating,omitempty"`
    // other media
    Image       []string                `json:"image,omitempty"        bson:"image,omitempty"`
    // external ids
    TMDB        int                     `json:"TMDB,omitempty"         bson:"TMDB,omitempty"`
    IMDB        string                  `json:"IMDB,omitempty"         bson:"IMDB,omitempty"`
    // adjcent media
    AdjShows    []primitive.ObjectID    `json:"adjShows,omitempty"     bson:"adjShows,omitempty"`
    AdjMovies   []primitive.ObjectID    `json:"adjMovies,omitempty"    bson:"adjMovies,omitempty"`
}


type ShowSeason struct {
    ShowId      primitive.ObjectID      `json:"showId,omitempty"       validate:"required"`
    SeasonID    int                     `json:"seasonId,omitempty"     validate:"required"`
    Episodes    int                     `json:"epesodes,omitempty"     validate:"required"`
    Description string                  `json:"description,omitempty"  validate:"required"`
    Date        string                  `json:"date,omitempty"         validate:"required"`
    Image       string                  `json:"image,omitempty"        bson:"image,omitempty"`
}


type ShowEpisode struct {
    ShowId      primitive.ObjectID      `json:"showId,omitempty"       validate:"required"`
    SeasonID    int                     `json:"seasonId,omitempty"     validate:"required"`
    EpisodeID   int                     `json:"epesodeId,omitempty"    validate:"required"`
    Title       string                  `json:"title,omitempty"        validate:"required"`
    Description string                  `json:"description,omitempty"  validate:"required"`
    Date        string                  `json:"date,omitempty"         bson:"date,omitempty"`
    VoteCount   int                     `json:"voteCount,omitempty"    bson:"voteCount,omitempty"`
    VoteRating  float64                 `json:"voteRating,omitempty"   bson:"voteRating,omitempty"`
    Image       string                  `json:"image,omitempty"        bson:"image,omitempty"`
}

var ShowCollection *mongo.Collection = configs.GetCollection(configs.DB, "show")
var ShowSeasonCollection *mongo.Collection = configs.GetCollection(configs.DB, "showSeason")
var ShowEpisodeCollection *mongo.Collection = configs.GetCollection(configs.DB, "showEpisode")

func (s *Show) Write() mongo.WriteModel {
    updateModel := mongo.NewUpdateOneModel()
    updateModel.SetFilter(bson.M{"title": s.Title, "date": s.Date}) 
    updateModel.SetUpdate(bson.D{{"$set", *s}})
    updateModel.SetUpsert(true)

    return updateModel
}

func (s *ShowSeason) Write() mongo.WriteModel {
    updateModel := mongo.NewUpdateOneModel()
    updateModel.SetFilter(bson.M{"seasonId": s.SeasonID, "showId": s.ShowId}) 
    updateModel.SetUpdate(bson.D{{"$set", *s}})
    updateModel.SetUpsert(true)
    
    return updateModel
}

func (s *ShowEpisode) Write() mongo.WriteModel {
    updateModel := mongo.NewUpdateOneModel()
    updateModel.SetFilter(bson.M{"seasonId": s.SeasonID, "showId": s.ShowId, "epesodeId": s.EpisodeID}) 
    updateModel.SetUpdate(bson.D{{"$set", *s}})
    updateModel.SetUpsert(true)

    return updateModel
}

func WriteShow(models []mongo.WriteModel) error {
    if len(models) == 0 {
        return nil
    }
    ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
    defer cancel()

    _, err := ShowCollection.BulkWrite(ctx, models)

    return err
}

func WriteShowSeason(models []mongo.WriteModel) error {
    if len(models) == 0 {
        return nil
    }
    ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
    defer cancel()

    _, err := ShowSeasonCollection.BulkWrite(ctx, models)

    return err
}

func WriteShowEpisode(models []mongo.WriteModel) error {
    if len(models) == 0 {
        return nil
    }
    ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
    defer cancel()

    _, err := ShowEpisodeCollection.BulkWrite(ctx, models)

    return err
}


func FindShow(filter bson.D) ([]Show, error) {
    ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
    var shows []Show
    defer cancel()

    results, err := ShowCollection.Find(ctx, filter)
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

    results, err := ShowSeasonCollection.Find(ctx, filter)
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

    results, err := ShowEpisodeCollection.Find(ctx, filter)
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
