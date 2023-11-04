package models

import (
    "context"
    "gin-mongo-api/configs"
    "time"

    "go.mongodb.org/mongo-driver/bson"
    "go.mongodb.org/mongo-driver/bson/primitive"
    "go.mongodb.org/mongo-driver/mongo"
)

type Company struct {
    Id primitive.ObjectID `json:"id,omitempty"           bson:"_id,omitempty"`
    //basic info
    Name        string `json:"name"                   bson:"name,omitempty"        validate:"required"`
    Description string `json:"description"            bson:"description,omitempty" `
    Date        string `json:"date,omitempty"         bson:"date,omitempty"        `
    // extra media
    Images      []string          `json:"images,omitempty"        bson:"images,omitempty"`
    ExternalIds map[string]string `json:"externalIds,omitempty"  bson:"externalIds,omitempty"`
    // works
    Movies []primitive.ObjectID `json:"movies,omitempty"       bson:"movies,omitempty"`
    Shows  []primitive.ObjectID `json:"shows,omitempty"        bson:"shows,omitempty"`
    Books  []primitive.ObjectID `json:"books,omitempty"        bson:"books,omitempty"`
    Games  []primitive.ObjectID `json:"games,omitempty"        bson:"games,omitempty"`
}

var companyCollection *mongo.Collection = configs.GetCollection(configs.DB, "company")

func WriteCompany(models []Company) error {
    if len(models) == 0 {
        return nil
    }
    ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
    defer cancel()

    var writeObjs []mongo.WriteModel
    for i := 0; i < len(models); i++ {
        updateModel := mongo.NewUpdateOneModel()
        updateModel.SetFilter(bson.M{"title": models[i].Name, "date": models[i].Date})
        updateModel.SetUpdate(bson.D{{"$set", models[i]}})
        updateModel.SetUpsert(true)

        writeObjs = append(writeObjs, updateModel)
    }

    _, err := companyCollection.BulkWrite(ctx, writeObjs)

    return err
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
