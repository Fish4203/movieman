package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Company struct {
	Id          primitive.ObjectID   `json:"id,omitempty"           bson:"_id,omitempty"`
	Name        string               `json:"name"   tmdb:"name,omitempty"                validate:"required"`
	Description string               `json:"description"           tmdb:"description,omitempty"  `
	Date        string               `json:"date,omitempty"                 `
	Image       []string             `json:"image,omitempty"     tmdb:"logo_path,omitempty"   `
	ExternalIds map[string]string    `json:"externalIds,omitempty"  `
	Movies      []primitive.ObjectID `json:"movies,omitempty"       `
	Shows       []primitive.ObjectID `json:"shows,omitempty"        `
	Books       []primitive.ObjectID `json:"books,omitempty"        `
	Games       []primitive.ObjectID `json:"games,omitempty"        `
}
