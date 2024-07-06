package models

type Movie struct {
  Media

  Budget        uint              `json:"budget"`
  Length        uint              `json:"length"`
  Rating        string            `json:"rating"`
  ExternalInfo  []MovieExternal   `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;ForeignKey:Title,Date;References:Title,Date"` 
  Review        []MovieReview     `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;ForeignKey:Title,Date;References:Title,Date"` 
}

type MovieExternal struct {
  MediaExternal
}

type MovieReview struct {
  MediaReview
}
