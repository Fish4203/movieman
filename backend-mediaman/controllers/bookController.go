package controllers

import (
	"backend-mediaman/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateBook() gin.HandlerFunc {
  return func(c *gin.Context) {
    var book models.Book

    if err := book.Save(c); err != nil {
      c.JSON(http.StatusBadRequest, map[string]interface{}{"error": err.Error()})
    } else {
      c.JSON(http.StatusCreated, map[string]interface{}{"book": book})
    }
  }
}

func GetBooks() gin.HandlerFunc {
  return func(c *gin.Context) {
    var books []*models.Book
    var book models.Book

    if err := models.SearchMedia(c, &book, &books); err != nil {
      c.JSON(http.StatusNotFound, map[string]interface{}{"error": err})
    } else {
      c.JSON(http.StatusOK, map[string]interface{}{"books": books})
    }
  }
}

func DeleteBook() gin.HandlerFunc {
  return func(c *gin.Context) {
    var book models.Book

    if err := book.Delete(c); err != nil {
      c.JSON(http.StatusBadRequest, map[string]interface{}{"error": err.Error()})
    } else {
      c.JSON(http.StatusOK, map[string]interface{}{"book": book})
    }
  }
}

func CreateBookReview() gin.HandlerFunc {
  return func(c *gin.Context) {
    var bookReview models.BookReview

    if err := bookReview.Save(c); err != nil {
      c.JSON(http.StatusBadRequest, map[string]interface{}{"error": err.Error()})
    } else {
      c.JSON(http.StatusCreated, map[string]interface{}{"review": bookReview})
    } 
  }
}

func GetBookReview() gin.HandlerFunc {
  return func(c *gin.Context) {
    var bookReview models.BookReview

    if err := bookReview.Get(c); err != nil {
      c.JSON(http.StatusBadRequest, map[string]interface{}{"error": err.Error()})
    } else {
      c.JSON(http.StatusOK, map[string]interface{}{"review": bookReview})
    } 
  }
}

func DeleteBookReview() gin.HandlerFunc {
  return func(c *gin.Context) {
    var bookReview models.BookReview

    if err := bookReview.Delete(c); err != nil {
      c.JSON(http.StatusBadRequest, map[string]interface{}{"error": err.Error()})
    } else {
      c.JSON(http.StatusOK, map[string]interface{}{"status": "deleted"})
    } 
  }
}
