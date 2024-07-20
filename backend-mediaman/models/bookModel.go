package models

import "github.com/gin-gonic/gin"

type Book struct {
  Media

  Length        uint              `json:"length"`
  
  ExternalInfo  []BookExternal   `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;ForeignKey:Title,Date;References:Title,Date"` 
  Review        []BookReview     `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;ForeignKey:Title,Date;References:Title,Date"` 
}

type BookExternal struct {
  MediaExternal
}

type BookReview struct {
  MediaReview
}

type BookUnion struct {
  Book
  BookExternal
}

func (book BookUnion) GetExternal() BookExternal {
  external := book.BookExternal
  external.Title = book.Book.Title
  external.Date  = book.Book.Date

  return external
}

func (book *Book) Save(c *gin.Context) error {
  return saveMedia(c, book)
}

func (book *Book) Get(c *gin.Context) error {
  return getMedia(c, book)
}

func (book *Book) Delete(c *gin.Context) error {
  return deleteMedia(c, book)
}

func (book *BookReview) Save(c *gin.Context) error {
  return saveReview(c, book)
}

func (book *BookReview) Get(c *gin.Context) error {
  return getReview(c, book)
}

func (book *BookReview) Delete(c *gin.Context) error {
  return deleteReview(c, book)
}
