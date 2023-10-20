package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Show struct {
	Id primitive.ObjectID `json:"id,omitempty"           bson:"_id,omitempty"         `
	// basic info
	Title       string   `json:"title"                  bson:"title,omitempty"       tmdb:"name,omitempty"             validate:"required"`
	Description string   `json:"description"            bson:"description,omitempty" tmdb:"overview,omitempty"         validate:"required"`
	Date        string   `json:"date,omitempty"         bson:"date,omitempty"        tmdb:"first_air_date,omitempty"   validate:"required"`
	Seasons     int      `json:"seasons"                bson:"seasons,omitempty"     tmdb:"number_of_seasons,omitempty"`
	Genre       []string `json:"genre,omitempty"        bson:"genre,omitempty"       `
	Info        string   `json:"info,omitempty"         bson:"info,omitempty"        tmdb:"homepage,omitempty"`
	Rating      string   `json:"rating,omitempty"       bson:"rating,omitempty"      `
	// other media
	Images []string `json:"images,omitempty"       bson:"images,omitempty"      tmdb"poster_path,omitempty"`
	// external ids
	Platforms   []string          `json:"platforms,omitempty"    bson:"platforms,omitempty"   `
	ExternalIds map[string]string `json:"externalIds,omitempty"  bson:"externalIds,omitempty" `
	Reviews     map[string]string `json:"reviews,omitempty"      bson:"reviews,omitempty" `
}

type ShowSeason struct {
	/// ids
	ShowId   primitive.ObjectID `json:"showId"                 bson:"showId,omitempty"                              validate:"required"`
	SeasonID int                `json:"seasonId"               bson:"seasonId,omitempty"    tmdb:"season_number"    validate:"required"`
	// basic info
	Episodes    int    `json:"episodes"               bson:"episodes,omitempty"    tmdb:"episodes"         validate:"required"`
	Description string `json:"description"            bson:"description,omitempty" tmdb:"overview"         validate:"required"`
	Date        string `json:"date,omitempty"         bson:"date,omitempty"        tmdb:"air_date"         validate:"required"`
	// external media
	Images []string `json:"images,omitempty"       bson:"images,omitempty"`
}

type ShowEpisode struct {
	// ids
	ShowId    primitive.ObjectID `json:"showId"                 bson:"showId,omitempty"                              validate:"required"`
	SeasonID  int                `json:"seasonId"               bson:"seasonId,omitempty"                            validate:"required"`
	EpisodeID int                `json:"episodeId"              bson:"episodeId,omitempty"   tmdb:"episode_number"   validate:"required"`
	// basic info
	Title       string `json:"title"                  bson:"title,omitempty"       tmdb:"name"             validate:"required"`
	Description string `json:"description"            bson:"description,omitempty" tmdb:"overview"         validate:"required"`
	Date        string `json:"date,omitempty"         bson:"date,omitempty"        tmdb:"air_date"`
	// ratings
	Reviews map[string]string `json:"reviews,omitempty"      bson:"reviews,omitempty" `
	// external media
	Images []string `json:"images,omitempty"       bson:"images,omitempty"`
}
