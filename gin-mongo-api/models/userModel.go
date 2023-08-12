package models

import (
    "gin-mongo-api/configs"
    "go.mongodb.org/mongo-driver/mongo"
    "go.mongodb.org/mongo-driver/bson/primitive"
)
type User struct {
    Id       primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
    Name     string             `json:"name,omitempty" validate:"required"`
    Password string             `json:"password,omitempty" validate:"required"`
    Role     string             `json:"role,omitempty"`
}

var UserCollection *mongo.Collection = configs.GetCollection(configs.DB, "users")
