package models

type Movie struct {
  Media

  Budget        uint              `json:"budget"        binding:"required"`
  Length        uint              `json:"length"        binding:"required"`
  ExternalInfo  []MovieExternal   `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;ForeignKey:MediaID;References:ID"` 
  Review        []MovieReview     `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;ForeignKey:MediaID;References:ID"` 
}

type MovieExternal struct {
  MediaExternal

  AgeRating     string            `json:"ageRating"     binding:"required"`
}

type MovieReview struct {
  MediaReview
}

type MovieUnion struct {
  Movie
  MovieExternal
}


