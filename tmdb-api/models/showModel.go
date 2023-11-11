package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Show struct {
	Id          primitive.ObjectID `json:"id,omitempty"           bson:"_id,omitempty"         `
	Title       string             `json:"title"                  bson:"title,omitempty"       tmdb:"name,omitempty"             validate:"required"`
	Description string             `json:"description"            bson:"description,omitempty" tmdb:"overview,omitempty"         validate:"required"`
	Date        string             `json:"date,omitempty"         bson:"date,omitempty"        tmdb:"first_air_date,omitempty"   validate:"required"`
	Seasons     int                `json:"seasons"                bson:"seasons,omitempty"     tmdb:"number_of_seasons,omitempty"`
	Genre       []string           `json:"genre,omitempty"        bson:"genre,omitempty"       `
	Info        string             `json:"info,omitempty"         bson:"info,omitempty"        tmdb:"homepage,omitempty"`
	Rating      string             `json:"rating,omitempty"       bson:"rating,omitempty"      `
	Images      []string           `json:"images,omitempty"       bson:"images,omitempty"      tmdb"poster_path,omitempty"`
	Platforms   []string           `json:"platforms,omitempty"    bson:"platforms,omitempty"   `
	ExternalIds map[string]string  `json:"externalIds,omitempty"  bson:"externalIds,omitempty" `
	Reviews     map[string]string  `json:"reviews,omitempty"      bson:"reviews,omitempty" `
}

type ShowSeason struct {
	ShowId      primitive.ObjectID `json:"showId"                 bson:"showId,omitempty"                              validate:"required"`
	SeasonID    int                `json:"seasonId"               bson:"seasonId,omitempty"    tmdb:"season_number"    validate:"required"`
	Episodes    int                `json:"episodes"               bson:"episodes,omitempty"    tmdb:"episode_count"         validate:"required"`
	Description string             `json:"description"            bson:"description,omitempty" tmdb:"overview"         validate:"required"`
	Date        string             `json:"date,omitempty"         bson:"date,omitempty"        tmdb:"air_date"         validate:"required"`
	Images      []string           `json:"images,omitempty"       bson:"images,omitempty"`
}

type ShowEpisode struct {
	ShowId      primitive.ObjectID `json:"showId"                 bson:"showId,omitempty"                              validate:"required"`
	SeasonID    int                `json:"seasonId"               bson:"seasonId,omitempty"                            validate:"required"`
	EpisodeID   int                `json:"episodeId"              bson:"episodeId,omitempty"   tmdb:"episode_number"   validate:"required"`
	Title       string             `json:"title"                  bson:"title,omitempty"       tmdb:"name"             validate:"required"`
	Description string             `json:"description"            bson:"description,omitempty" tmdb:"overview"         validate:"required"`
	Date        string             `json:"date,omitempty"         bson:"date,omitempty"        tmdb:"air_date"`
	Reviews     map[string]string  `json:"reviews,omitempty"      bson:"reviews,omitempty" `
	Images      []string           `json:"images,omitempty"       bson:"images,omitempty"`
}
