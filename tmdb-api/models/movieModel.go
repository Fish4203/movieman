package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Movie struct {
	Id primitive.ObjectID `json:"id,omitempty"          bson:"_id,omitempty"         `
	// basic info
	Title       string   `json:"title"                  bson:"title,omitempty"       tmdb:"title"        			validate:"required"`
	Description string   `json:"description"            bson:"description,omitempty" tmdb:"overview,omitempty"     	validate:"required"`
	Date        string   `json:"date,omitempty"         bson:"date,omitempty"        tmdb:"release_date,omitempty" 	validate:"required"`
	Genre       []string `json:"genre,omitempty"        bson:"genre,omitempty"       `
	Info        string   `json:"info,omitempty"         bson:"info,omitempty"        tmdb:"homepage,omitempty"`
	Budget      int      `json:"budget,omitempty"       bson:"budget,omitempty"      tmdb:"budget,omitempty"`
	Length      int      `json:"length,omitempty"       bson:"length,omitempty"      tmdb:"runtime,omitempty"`
	Rating      string   `json:"rating,omitempty"       bson:"rating,omitempty"      `
	// pupularity
	Reviews map[string]string `json:"reviews,omitempty"      bson:"reviews,omitempty" `
	// related media
	Images []string `json:"images,omitempty"       bson:"images,omitempty"      tmdb:"poster_path,omitempty"`
	// ids
	ExternalIds map[string]string `json:"externalIds,omitempty"  bson:"externalIds,omitempty" `
	Platforms   []string          `json:"platforms,omitempty"    bson:"platforms,omitempty"   `
}
