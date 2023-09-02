package models

import (
    "time"
    "context"
    // "gin-mongo-api/configs"
    // "github.com/go-playground/validator/v10"
    "go.mongodb.org/mongo-driver/mongo"
    // "go.mongodb.org/mongo-driver/mongo/options"
    // "go.mongodb.org/mongo-driver/bson"
    // "go.mongodb.org/mongo-driver/bson/primitive"
)

type ModelInterface interface {
	Write() mongo.WriteModel
	Collection() *mongo.Collection 
}


func BulkWrite(models []ModelInterface) error {
    if len(models) == 0 {
        return nil
    }
    ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
    defer cancel()

    var writeObjs []mongo.WriteModel 
	for i := 0; i < len(models); i++ {
		writeObjs = append(writeObjs, models[i].Write())
	}
	
    _, err := models[1].Collection().BulkWrite(ctx, writeObjs)

    return err
}
