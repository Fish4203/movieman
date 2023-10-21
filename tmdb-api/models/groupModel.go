package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Group struct {
	Id          primitive.ObjectID   `json:"id,omitempty"      `
	Title       string               `json:"title"  			tmdb:"name,omitempty"       validate:"required"`
	Description string               `json:"description"        tmdb:"overview,omitempty" 	validate:"required"`
	Genre       []string             `json:"genre,omitempty"    `
	Info        string               `json:"info,omitempty"     `
	Movies      []primitive.ObjectID `json:"movies,omitempty"       `
	Shows       []primitive.ObjectID `json:"shows,omitempty"        `
	Books       []primitive.ObjectID `json:"books,omitempty"        `
	Games       []primitive.ObjectID `json:"games,omitempty"        `
	Image       []string             `json:"image,omitempty"  tmdb:"poster_path,omitempty"       `
	ExternalIds map[string]string    `json:"externalIds,omitempty"  `
}
