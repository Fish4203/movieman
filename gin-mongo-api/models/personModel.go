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
    Id          primitive.ObjectID `json:"id,omitempty"           bson:"_id,omitempty"`
    Name        string             `json:"name,omitempty"         validate:"required"`
    Description string             `json:"description,omitempty"  bson:"description,omitempty"`
    Role        string             `json:"role,omitempty"         bson:"role,omitempty"`
    Date        string             `json:"date,omitempty"         bson:"date,omitempty"`
    Image       []string           `json:"image,omitempty"        bson:"image,omitempty"`
    Popularity  float64            `json:"popularity,omitempty"   bson:"popularity,omitempty"`
    TMDB        int                `json:"TMDB,omitempty"         bson:"TMDB,omitempty"`
    IMDB        string             `json:"IMDB,omitempty"         bson:"IMDB,omitempty"`
}

var PersonCollection *mongo.Collection = configs.GetCollection(configs.DB, "person")



func (p *Person) Write() mongo.WriteModel {
    updateModel := mongo.NewUpdateOneModel()
    updateModel.SetFilter(bson.M{"name": p.Name, "date": p.Date}) 
    updateModel.SetUpdate(bson.D{{"$set", *p}})
    updateModel.SetUpsert(true)

    return updateModel
}

func WritePerson(models []mongo.WriteModel) error {
    if len(models) == 0 {
        return nil
    }
    ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
    defer cancel()

    _, err := PersonCollection.BulkWrite(ctx, models)

    return err
}

func FindPerson(filter bson.D) ([]Person, error) {
    ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
    var persons []Person
    defer cancel()

    results, err := PersonCollection.Find(ctx, filter)
    if err != nil {
        return persons, err
    }

    //reading from the db in an optimal way
    defer results.Close(ctx)
    for results.Next(ctx) {
        var singlePerson Person
        if err = results.Decode(&singlePerson); err != nil {
            return persons, err
        }
        persons = append(persons, singlePerson)
    }

    return persons, nil
}
