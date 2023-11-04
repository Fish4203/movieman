package models

import (
	"context"
	"gin-mongo-api/configs"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type Person struct {
	Id primitive.ObjectID `json:"id,omitempty"           bson:"_id,omitempty"`
	//basic info
	Name        string `json:"name"                   bson:"name,omitempty"        validate:"required"`
	Description string `json:"description"            bson:"description,omitempty" `
	Role        string `json:"role"                   bson:"role,omitempty"        `
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

var personCollection *mongo.Collection = configs.GetCollection(configs.DB, "person")

func WritePerson(models []Person) error {
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

	_, err := personCollection.BulkWrite(ctx, writeObjs)

	return err
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
