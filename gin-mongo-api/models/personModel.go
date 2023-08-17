package models

import (
    "time"
    "context"
    "gin-mongo-api/configs"
    "go.mongodb.org/mongo-driver/mongo"
    "go.mongodb.org/mongo-driver/bson"
    "go.mongodb.org/mongo-driver/bson/primitive"
)
type Person struct {
    Id          primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
    Name        string             `json:"name,omitempty" validate:"required"`
    Description string             `json:"description,omitempty"`
    Date        string             `json:"date,omitempty"`
    Image       string             `json:"image,omitempty"`
    TMDB        uint               `json:"TMDB,omitempty"`
    IMDB        string             `json:"IMDB,omitempty"`
}

var PersonCollection *mongo.Collection = configs.GetCollection(configs.DB, "person")


func (p *Person) Save() error {
    ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
    defer cancel()

    if err := validate.Struct(p); err != nil {
        return err
    }

    _, err := PersonCollection.UpdateOne(ctx, bson.M{"title": p.Name, "date": p.Date}, *p)

    return err
}
