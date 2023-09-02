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
    Title       string                  `json:"title,omitempty"        bson:"title,omitempty"       tmdb:"name"             validate:"required"`
    Description string                  `json:"description,omitempty"  bson:"description,omitempty" tmdb:"overview"         validate:"required"`
    Date        string                  `json:"date,omitempty"         bson:"date,omitempty"        tmdb:"first_air_date"   validate:"required"`
    Seasons     int                     `json:"seasons,omitempty"      bson:"seasons,omitempty"     tmdb:"number_of_seasons"`
    Genre       []string                `json:"genre,omitempty"        bson:"genre,omitempty"`
    Info        string                  `json:"info,omitempty"         bson:"info,omitempty"        tmdb:"homepage"`
    // reviews
    Popularity  float64                 `json:"popularity,omitempty"   bson:"popularity,omitempty"  tmdb:"popularity"`
    VoteCount   int                     `json:"voteCount,omitempty"    bson:"voteCount,omitempty"   tmdb:"vote_count"`
    VoteRating  float64                 `json:"voteRating,omitempty"   bson:"voteRating,omitempty"  tmdb:"vote_average"`
    Rating      string                  `json:"rating,omitempty"       bson:"rating,omitempty"`
    // other media
    Images      []string                `json:"images,omitempty"       bson:"images,omitempty"`
    // external ids
    Platforms   []string                `json:"platforms,omitempty"    bson:"platforms,omitempty"`
    ExternalIds map[string]string       `json:"externalIds,omitempty"  bson:"externalIds,omitempty"`
}


type ShowSeason struct {
    /// ids 
    ShowId      primitive.ObjectID      `json:"showId,omitempty"       bson:"showId,omitempty"                              validate:"required"`
    SeasonID    int                     `json:"seasonId,omitempty"     bson:"seasonId,omitempty"    tmdb:"season_number"    validate:"required"`
    // basic info 
    Episodes    int                     `json:"epesodes,omitempty"     bson:"epesodes,omitempty"    tmdb:"episodes"         validate:"required"`
    Description string                  `json:"description,omitempty"  bson:"description,omitempty" tmdb:"overview"         validate:"required"`
    Date        string                  `json:"date,omitempty"         bson:"date,omitempty"        tmdb:"air_date"         validate:"required"`
    // external media 
    Images      []string                `json:"images,omitempty"       bson:"images,omitempty"`
}


type ShowEpisode struct {
    // ids
    ShowId      primitive.ObjectID      `json:"showId,omitempty"       bson:"showId,omitempty"                              validate:"required"`
    SeasonID    int                     `json:"seasonId,omitempty"     bson:"seasonId,omitempty"                            validate:"required"`
    EpisodeID   int                     `json:"epesodeId,omitempty"    bson:"epesodeId,omitempty"   tmdb:"episode_number"   validate:"required"`
    // basic info 
    Title       string                  `json:"title,omitempty"        bson:"title,omitempty"       tmdb:"name"             validate:"required"`
    Description string                  `json:"description,omitempty"  bson:"description,omitempty" tmdb:"overview"         validate:"required"`
    Date        string                  `json:"date,omitempty"         bson:"date,omitempty"        tmdb:"air_date"`
    // ratings 
    VoteCount   int                     `json:"voteCount,omitempty"    bson:"voteCount,omitempty"   tmdb:"vote_count"`
    VoteRating  float64                 `json:"voteRating,omitempty"   bson:"voteRating,omitempty"  tmdb:"vote_average"`
    // external media
    Images      []string                `json:"images,omitempty"       bson:"images,omitempty"`
}

var showCollection *mongo.Collection = configs.GetCollection(configs.DB, "show")
var showSeasonCollection *mongo.Collection = configs.GetCollection(configs.DB, "showSeason")
var showEpisodeCollection *mongo.Collection = configs.GetCollection(configs.DB, "showEpisode")

func (s *Show) Collection() *mongo.Collection {return showCollection}
func (s *ShowSeason) Collection() *mongo.Collection {return showSeasonCollection}
func (s *ShowEpisode) Collection() *mongo.Collection {return showEpisodeCollection}


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
