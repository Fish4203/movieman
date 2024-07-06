package models

type Movie struct {
  Media

  Budget        uint              `json:"budget"`
  Length        uint              `json:"length"`
  Rating        string            `json:"rating"`
  ExternalInfo  []MovieExternal   
  Review        []MovieReview
}

type MovieExternal struct {
  MediaExternal
  Movie   Movie  `json:"movie"  gorm:"not null"`
}

type MovieReview struct {
  MediaReview
  Movie   Movie  `json:"movie"    gorm:"not null;primaryKey"`
}
