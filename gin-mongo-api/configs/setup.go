package configs

import (
    "context"
    "fmt"
    "log"
    "time"
    "os"
    "github.com/joho/godotenv"
    // "go.mongodb.org/mongo-driver/bson"
    "go.mongodb.org/mongo-driver/mongo"
    "go.mongodb.org/mongo-driver/mongo/options"
)

func InitClient() *mongo.Client  {
    err := godotenv.Load()
    if err != nil {
        log.Fatal("Error loading .env file")
    }

    client, err := mongo.NewClient(options.Client().ApplyURI(os.Getenv("MONGOURI")))
    if err != nil {
        log.Fatal(err)
    }

    ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
    err = client.Connect(ctx)
    if err != nil {
        log.Fatal(err)
    }

    //ping the database
    err = client.Ping(ctx, nil)
    if err != nil {
        log.Fatal(err)
    }
    fmt.Println("Connected to MongoDB")
    return client
}
//getting database collections
func GetCollection(client *mongo.Client, collectionName string) *mongo.Collection {
    fmt.Println(collectionName)
    collection := client.Database("movieman").Collection(collectionName)
    return collection
}

//Client instance
var DB *mongo.Client = InitClient()
