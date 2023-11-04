package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Person struct {
	Id          primitive.ObjectID   `json:"id,omitempty"           `
	Name        string               `json:"name"                           tmdb:"name"    validate:"required"`
	Description string               `json:"description"             tmdb:"biography"`
	Role        string               `json:"role"                           tmdb:"known_for_department"`
	Date        string               `json:"date,omitempty"                 tmdb:"birthday"`
	Images      []string             `json:"images,omitempty"       `
	ExternalIds map[string]string    `json:"externalIds,omitempty"  `
	Movies      []primitive.ObjectID `json:"movies,omitempty"       `
	Shows       []primitive.ObjectID `json:"shows,omitempty"        `
	Books       []primitive.ObjectID `json:"books,omitempty"        `
	Games       []primitive.ObjectID `json:"games,omitempty"        `
}
