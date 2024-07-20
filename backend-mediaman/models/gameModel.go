package models

import "github.com/gin-gonic/gin"

type Game struct {
  Media

  Budget        uint              `json:"budget"`
  Rating        string            `json:"rating"`
  NumA          uint              `json:"aaaness"`

  Windows       bool                            
  Mac           bool
  Linux         bool
  Xbox          bool
  PlayStation   bool              `json:"playStation"`
  Nintendo      bool


  ExternalInfo  []GameExternal   `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;ForeignKey:Title,Date;References:Title,Date"` 
  Review        []GameReview     `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;ForeignKey:Title,Date;References:Title,Date"` 
}

type GameExternal struct {
  MediaExternal
}

type GameReview struct {
  MediaReview
  
  PlayTime    uint            `json:"playTime"`
}

type GameUnion struct {
  Game
  GameExternal
}

func (game *Game) Save(c *gin.Context) error {
  return saveMedia(c, game)
}

func (game *Game) Get(c *gin.Context) error {
  return getMedia(c, game)
}

func (game *Game) Delete(c *gin.Context) error {
  return deleteMedia(c, game)
}

func (game *GameReview) Save(c *gin.Context) error {
  return saveReview(c, game)
}

func (game *GameReview) Get(c *gin.Context) error {
  return getReview(c, game)
}

func (game *GameReview) Delete(c *gin.Context) error {
  return deleteReview(c, game)
}
