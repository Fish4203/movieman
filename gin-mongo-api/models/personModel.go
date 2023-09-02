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
    Id          primitive.ObjectID      `json:"id,omitempty"           bson:"_id,omitempty"`
    //basic info
    Name        string                  `json:"name,omitempty"         bson:"name,omitempty"        tmdb:"name"    validate:"required"`
    Description string                  `json:"description,omitempty"  bson:"description,omitempty" tmdb:"biography"`
    Role        string                  `json:"role,omitempty"         bson:"role,omitempty"        tmdb:"known_for_department"`
    Date        string                  `json:"date,omitempty"         bson:"date,omitempty"        tmdb:"birthday"`
    // extra media 
    Image       []string                `json:"image,omitempty"        bson:"image,omitempty"`
    ExternalIds map[string]string       `json:"externalIds,omitempty"  bson:"externalIds,omitempty"`
    // works 
    Movies      []primitive.ObjectID    `json:"movies,omitempty"       bson:"movies,omitempty"`
    Shows       []primitive.ObjectID    `json:"shows,omitempty"        bson:"shows,omitempty"`
    Books       []primitive.ObjectID    `json:"books,omitempty"        bson:"books,omitempty"`
    Games       []primitive.ObjectID    `json:"games,omitempty"        bson:"games,omitempty"`
}

var personCollection *mongo.Collection = configs.GetCollection(configs.DB, "person")

func (o *Person) Collection() *mongo.Collection {return personCollection}


func (p *Person) Write() mongo.WriteModel {
    updateModel := mongo.NewUpdateOneModel()
    updateModel.SetFilter(bson.M{"name": p.Name, "date": p.Date}) 
    updateModel.SetUpdate(bson.D{{"$set", *p}})
    updateModel.SetUpsert(true)

    return updateModel
}

func FindPerson(filter bson.D) ([]Person, error) {
    ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
    var persons []Person
    defer cancel()

    results, err := personCollection.Find(ctx, filter)
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
