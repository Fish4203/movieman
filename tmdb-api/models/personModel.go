package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Person struct {
	Id          primitive.ObjectID   `json:"id,omitempty"           bson:"_id,omitempty"`
	Name        string               `json:"name"                   bson:"name,omitempty"        tmdb:"name"    validate:"required"`
	Description string               `json:"description"            bson:"description,omitempty" tmdb:"biography"`
	Role        string               `json:"role"                   bson:"role,omitempty"        tmdb:"known_for_department"`
	Date        string               `json:"date,omitempty"         bson:"date,omitempty"        tmdb:"birthday"`
	Image       []string             `json:"image,omitempty"        bson:"image,omitempty"`
	ExternalIds map[string]string    `json:"externalIds,omitempty"  bson:"externalIds,omitempty"`
	Movies      []primitive.ObjectID `json:"movies,omitempty"       bson:"movies,omitempty"`
	Shows       []primitive.ObjectID `json:"shows,omitempty"        bson:"shows,omitempty"`
	Books       []primitive.ObjectID `json:"books,omitempty"        bson:"books,omitempty"`
	Games       []primitive.ObjectID `json:"games,omitempty"        bson:"games,omitempty"`
}
