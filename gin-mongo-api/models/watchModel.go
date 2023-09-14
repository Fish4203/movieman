package models

import (
    "time"
    "context"
    "gin-mongo-api/configs"
    "go.mongodb.org/mongo-driver/mongo"
    "go.mongodb.org/mongo-driver/bson"
    "go.mongodb.org/mongo-driver/bson/primitive"
)

type Watch struct {
    Movie       primitive.ObjectID `json:"movie,omitempty"      bson:"movie,omitempty"`
    Show        primitive.ObjectID `json:"show,omitempty"       bson:"show,omitempty"`
    Book        primitive.ObjectID `json:"gook,omitempty"       bson:"gook,omitempty"`
    Game        primitive.ObjectID `json:"game,omitempty"       bson:"game,omitempty"`
    User        primitive.ObjectID `json:"user,omitempty"       bson:"user,omitempty"       validate:"required"`
    Title       string             `json:"title,omitempty"      bson:"title,omitempty"      validate:"required"`
    Watched     int                `json:"watched,omitempty"    bson:"watched,omitempty"    validate:"required"`
    UserRating  float64            `json:"userRating,omitempty" bson:"userRating,omitempty"`
    Notes       string             `json:"notes,omitempty"      bson:"notes,omitempty"`
}

var watchCollection *mongo.Collection = configs.GetCollection(configs.DB, "watch")

func WriteWatch(models []Watch) error {
    if len(models) == 0 {
        return nil
    }
    ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
    defer cancel()

    var writeObjs []mongo.WriteModel 
	for i := 0; i < len(models); i++ {
        updateModel := mongo.NewUpdateOneModel()
        updateModel.SetFilter(bson.M{"user": models[i].User, "$or": bson.A{
            bson.M{"movie": models[i].Movie},
            bson.M{"show": models[i].Game},
            bson.M{"book": models[i].Book},
            bson.M{"game": models[i].Game},
        }}) 
        updateModel.SetUpdate(bson.D{{"$set", models[i]}})
        updateModel.SetUpsert(true)

		writeObjs = append(writeObjs, updateModel)
	}
	
    _, err := watchCollection.BulkWrite(ctx, writeObjs)

    return err
}

func DeleteWatch(ids []string) error {
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

    _, err := watchCollection.DeleteMany(ctx, bson.D{{"$or", idobjs}})
    return err
}


func FindWatch(filter bson.D) ([]Watch, error) {
    ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
    var watchs []Watch
    defer cancel()

    results, err := watchCollection.Find(ctx, filter)
    if err != nil {
        return watchs, err
    }

    //reading from the db in an optimal way
    defer results.Close(ctx)
    for results.Next(ctx) {
        var singleWatch Watch
        if err = results.Decode(&singleWatch); err != nil {
            return watchs, err
        }
        watchs = append(watchs, singleWatch)
    }

    return watchs, nil
}
