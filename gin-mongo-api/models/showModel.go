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
    Id          primitive.ObjectID `json:"id,omitempty"           bson:"_id,omitempty"`
    Title       string             `json:"title,omitempty"        validate:"required"`
    Description string             `json:"description,omitempty"  validate:"required"`
    Date        string             `json:"date,omitempty"         validate:"required"`
    Seasons     int                `json:"seasons,omitempty"      bson:"seasons,omitempty"`
    Genre       []string           `json:"genre,omitempty"        bson:"genre,omitempty"`
    Info        string             `json:"info,omitempty"         bson:"info,omitempty"`
    Popularity  float64            `json:"popularity,omitempty"   bson:"popularity,omitempty"`
    VoteCount   int                `json:"voteCount,omitempty"    bson:"voteCount,omitempty"`
    VoteRating  float64            `json:"voteRating,omitempty"   bson:"voteRating,omitempty"`
    Rating      string             `json:"rating,omitempty"       bson:"rating,omitempty"`
    Image       string             `json:"image,omitempty"        bson:"image,omitempty"`
    TMDB        int                `json:"TMDB,omitempty"         bson:"TMDB,omitempty"`
    IMDB        string             `json:"IMDB,omitempty"         bson:"IMDB,omitempty"`
}


type ShowSeason struct {
    ShowId      primitive.ObjectID `json:"showId,omitempty"       validate:"required"`
    SeasonID    uint               `json:"seasonId,omitempty"     validate:"required"`
    Epesodes    uint               `json:"epesodes,omitempty"     validate:"required"`
    Date        string             `json:"date,omitempty"         bson:"date,omitempty"`
    Popularity  float64            `json:"popularity,omitempty"   bson:"popularity,omitempty"`
    VoteCount   int                `json:"voteCount,omitempty"    bson:"voteCount,omitempty"`
    VoteRating  float64            `json:"voteRating,omitempty"   bson:"voteRating,omitempty"`
    Rating      string             `json:"rating,omitempty"       bson:"rating,omitempty"`
    Image       string             `json:"image,omitempty"        bson:"image,omitempty"`
}


type ShowEpisode struct {
    ShowId      primitive.ObjectID `json:"showId,omitempty"       validate:"required"`
    SeasonID    uint               `json:"seasonId,omitempty"     validate:"required"`
    EpesodeID   uint               `json:"epesodeId,omitempty"    validate:"required"`
    Title       string             `json:"title,omitempty"        validate:"required"`
    Date        string             `json:"date,omitempty"         bson:"date,omitempty"`
    Popularity  float64            `json:"popularity,omitempty"   bson:"popularity,omitempty"`
    VoteCount   int                `json:"voteCount,omitempty"    bson:"voteCount,omitempty"`
    VoteRating  float64            `json:"voteRating,omitempty"   bson:"voteRating,omitempty"`
    Image       string             `json:"image,omitempty"        bson:"image,omitempty"`
}

var ShowCollection *mongo.Collection = configs.GetCollection(configs.DB, "show")
var ShowSeasonCollection *mongo.Collection = configs.GetCollection(configs.DB, "showSeason")
var ShowEpisodeCollection *mongo.Collection = configs.GetCollection(configs.DB, "showEpisode")


func (s *Show) Save() error {
    ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
    defer cancel()

    if err := validate.Struct(s); err != nil {
        return err
    }

    update := bson.D{{"$set", *s}}
    filter := bson.M{"title": s.Title, "date": s.Date}
    _, err := ShowCollection.UpdateOne(ctx, filter, update, updateOpts)

    return err
}

func (s *ShowSeason) Save() error {
    ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
    defer cancel()

    if err := validate.Struct(s); err != nil {
        return err
    }

    update := bson.D{{"$set", *s}}
    filter := bson.M{"seasonId": s.SeasonID, "showId": s.ShowId}
    _, err := ShowSeasonCollection.UpdateOne(ctx, filter, update, updateOpts)

    return err
}

func (s *ShowEpisode) Save() error {
    ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
    defer cancel()

    if err := validate.Struct(s); err != nil {
        return err
    }

    update := bson.D{{"$set", *s}}
    filter := bson.M{"seasonId": s.SeasonID, "showId": s.ShowId, "epesodeId": s.EpesodeID}
    _, err := ShowEpisodeCollection.UpdateOne(ctx, filter, update, updateOpts)

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
