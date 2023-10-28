package models

import (
	"context"
	"fmt"
	"gin-mongo-api/configs"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type Provider struct {
	Id     primitive.ObjectID `json:"id,omitempty"        bson:"_id,omitempty"`
	Name   string             `json:"name"                bson:"name,omitempty"       validate:"required"`
	Domain string             `json:"domain" bson:"domain" validate:"required"`
	Types  string             `json:"types" bson:"types" validate:"required"`
}

var provCollection *mongo.Collection = configs.GetCollection(configs.DB, "providers")

func WriteProv(model Provider) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var writeObjs []mongo.WriteModel

	updateModel := mongo.NewUpdateOneModel()
	updateModel.SetFilter(bson.M{"name": model.Name})
	updateModel.SetUpdate(bson.D{{"$set", model}})
	updateModel.SetUpsert(true)

	writeObjs = append(writeObjs, updateModel)

	_, err := provCollection.BulkWrite(ctx, writeObjs)

	return err
}

func DeleteProv(ids []string) error {
	if len(ids) == 0 {
		return nil
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	var idobjs bson.A

	for i := 0; i < len(ids); i++ {
		idobj, _ := primitive.ObjectIDFromHex(ids[i])
		idobjs = append(idobjs, bson.M{"_id": idobj})
	}

	res, err := provCollection.DeleteMany(ctx, bson.D{{"$or", idobjs}})
	if err != nil {
		return err
	}
	if res.DeletedCount == 0 {
		return fmt.Errorf("didnt find anything to delete")
	}
	return nil
}

func FindProvs() ([]Provider, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	var provs []Provider
	defer cancel()

	results, err := provCollection.Find(ctx, bson.M{})
	if err != nil {
		return provs, err
	}

	//reading from the db in an optimal way
	defer results.Close(ctx)
	for results.Next(ctx) {
		var singleProv Provider
		if err = results.Decode(&singleProv); err != nil {
			return provs, err
		}
		provs = append(provs, singleProv)
	}

	return provs, nil
}
