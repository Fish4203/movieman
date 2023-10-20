package models

import (

	// "github.com/go-playground/validator/v10"

	// "go.mongodb.org/mongo-driver/mongo/options"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Group struct {
	Id primitive.ObjectID `json:"id,omitempty"          `
	// basic info
	Title       string   `json:"title"                         validate:"required"`
	Description string   `json:"description"             validate:"required"`
	Genre       []string `json:"genre,omitempty"        `
	Info        string   `json:"info,omitempty"         `
	// media
	Movies []primitive.ObjectID `json:"movies,omitempty"       `
	Shows  []primitive.ObjectID `json:"shows,omitempty"        `
	Books  []primitive.ObjectID `json:"books,omitempty"        `
	Games  []primitive.ObjectID `json:"games,omitempty"        `
	// related media
	Image []string `json:"image,omitempty"        `
	// ids
	ExternalIds map[string]string `json:"externalIds,omitempty"  `
}
