package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Company struct {
	Id primitive.ObjectID `json:"id,omitempty"           bson:"_id,omitempty"`
	//basic info
	Name        string `json:"name"                   validate:"required"`
	Description string `json:"description"             `
	Date        string `json:"date,omitempty"                 `
	// extra media
	Image       []string          `json:"image,omitempty"        `
	ExternalIds map[string]string `json:"externalIds,omitempty"  `
	// works
	Movies []primitive.ObjectID `json:"movies,omitempty"       `
	Shows  []primitive.ObjectID `json:"shows,omitempty"        `
	Books  []primitive.ObjectID `json:"books,omitempty"        `
	Games  []primitive.ObjectID `json:"games,omitempty"        `
}
