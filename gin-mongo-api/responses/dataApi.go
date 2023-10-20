package responses

import (
	"gin-mongo-api/models"
)

type BulkRequest struct {
	Movies       []models.Movie         `json:"movies"`
	Shows        []models.Show          `json:"shows"`
	ShowSeasons  [][]models.ShowSeason  `json:"showSeasons"`
	ShowEpisodes [][]models.ShowEpisode `json:"showEpisodes"`
	Books        []models.Book          `json:"books"`
	Games        []models.Game          `json:"games"`
	People       []models.Person        `json:"people"`
	Companies    []models.Company       `json:"companies"`
	Groups       []models.Group         `json:"groups"`
	Linked       bool                   `json:"linked"`
}
