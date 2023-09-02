package models

import (
    "time"
    "context"
    "gin-mongo-api/configs"
    "go.mongodb.org/mongo-driver/mongo"
    "go.mongodb.org/mongo-driver/bson"
    "go.mongodb.org/mongo-driver/bson/primitive"
)
type Company struct {
    Id          primitive.ObjectID      `json:"id,omitempty"           bson:"_id,omitempty"`
    //basic info
    Name        string                  `json:"name,omitempty"         bson:"name,omitempty"        validate:"required"`
    Description string                  `json:"description,omitempty"  bson:"description,omitempty" `
	Date        string                  `json:"date,omitempty"         bson:"date,omitempty"        `
    // extra media 
    Image       []string                `json:"image,omitempty"        bson:"image,omitempty"`
    ExternalIds map[string]string       `json:"externalIds,omitempty"  bson:"externalIds,omitempty"`
    // works 
    Movies      []primitive.ObjectID    `json:"movies,omitempty"       bson:"movies,omitempty"`
    Shows       []primitive.ObjectID    `json:"shows,omitempty"        bson:"shows,omitempty"`
    Books       []primitive.ObjectID    `json:"books,omitempty"        bson:"books,omitempty"`
    Games       []primitive.ObjectID    `json:"games,omitempty"        bson:"games,omitempty"`
}

var companyCollection *mongo.Collection = configs.GetCollection(configs.DB, "company")

func (o *Company) Collection() *mongo.Collection {return companyCollection}


func (p *Company) Write() mongo.WriteModel {
    updateModel := mongo.NewUpdateOneModel()
    updateModel.SetFilter(bson.M{"name": p.Name, "date": p.Date}) 
    updateModel.SetUpdate(bson.D{{"$set", *p}})
    updateModel.SetUpsert(true)

    return updateModel
}

func FindCompany(filter bson.D) ([]Company, error) {
    ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
    var companys []Company
    defer cancel()

    results, err := companyCollection.Find(ctx, filter)
    if err != nil {
        return companys, err
    }

    //reading from the db in an optimal way
    defer results.Close(ctx)
    for results.Next(ctx) {
        var singleCompany Company
        if err = results.Decode(&singleCompany); err != nil {
            return companys, err
        }
        companys = append(companys, singleCompany)
    }

    return companys, nil
}
