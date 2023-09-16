package models

import (
    "fmt"
    "time"
    "context"
    "gin-mongo-api/configs"
    // "github.com/go-playground/validator/v10"
    "go.mongodb.org/mongo-driver/mongo"
    // "go.mongodb.org/mongo-driver/mongo/options"
    "go.mongodb.org/mongo-driver/bson"
    "go.mongodb.org/mongo-driver/bson/primitive"
)


type User struct {
    Id       primitive.ObjectID `json:"id,omitempty"        bson:"_id,omitempty"`
    Name     string             `json:"name,omitempty"      bson:"name,omitempty"       validate:"required"`
    Password string             `json:"password,omitempty"  bson:"password,omitempty"   validate:"required"`
    Role     string             `json:"role,omitempty"      bson:"role,omitempty"`
}

var userCollection *mongo.Collection = configs.GetCollection(configs.DB, "users")

func WriteUser(models []User) error {
    if len(models) == 0 {
        return nil
    }
    ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
    defer cancel()

    var writeObjs []mongo.WriteModel 
	for i := 0; i < len(models); i++ {
        updateModel := mongo.NewUpdateOneModel()
        updateModel.SetFilter(bson.M{"name": models[i].Name}) 
        updateModel.SetUpdate(bson.D{{"$set", models[i]}})
        updateModel.SetUpsert(true)

		writeObjs = append(writeObjs, updateModel)
	}
	
    _, err := userCollection.BulkWrite(ctx, writeObjs)

    return err
}

func DeleteUser(ids []string) error {
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

    res, err := userCollection.DeleteMany(ctx, bson.D{{"$or", idobjs}})
    if (err != nil) {
        return err
    }
    if (res.DeletedCount == 0) {
        return fmt.Errorf("Didnt find anything to delete")
    }
    return nil
}


func FindUser(filter bson.D) ([]User, error) {
    ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
    var users []User
    defer cancel()

    results, err := userCollection.Find(ctx, filter)
    if err != nil {
        return users, err
    }

    //reading from the db in an optimal way
    defer results.Close(ctx)
    for results.Next(ctx) {
        var singleUser User
        if err = results.Decode(&singleUser); err != nil {
            return users, err
        }
        users = append(users, singleUser)
    }

    return users, nil
}
