package controllers

import (
	"backend-mediaman/configs"
	"backend-mediaman/middleware"
	"backend-mediaman/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm/clause"
)

func GetProvidor() gin.HandlerFunc {
  return func(c *gin.Context) {
    var providor models.DataProvider

    providorID, err := strconv.Atoi(c.Param("providorID"))
    if err != nil {
      c.JSON(http.StatusBadRequest, map[string]interface{}{"error": "Invalid providorID"})
      return
    }

    if result := configs.DB.Preload(clause.Associations).First(&providor, providorID); result.Error != nil || result.RowsAffected != 1 {
      c.JSON(http.StatusNotFound, map[string]interface{}{"error": result.Error})
       return
    }
        
    providor.BaseUrl = ""
    c.JSON(http.StatusOK, map[string]interface{}{"dataProvidor": providor})
  }
}

func GetProvidors() gin.HandlerFunc {
  return func(c *gin.Context) {
    var providors []models.DataProvider

    if result := configs.DB.Find(&providors); result.Error != nil {
      c.JSON(http.StatusNotFound, map[string]interface{}{"error": result.Error})
       return
    }
       
    for i := 0; i < len(providors); i++ {
      providors[i].BaseUrl = ""
    }

    c.JSON(http.StatusOK, map[string]interface{}{"dataProvidors": providors})
  }
}

func GenerateProviderAuthToken() gin.HandlerFunc {
  return func(c *gin.Context) {
    userID := c.GetUint("userID")
    var user models.User
    var provider models.DataProvider

    if result := configs.DB.First(&user, userID); result.Error != nil || result.RowsAffected != 1 {
      c.JSON(http.StatusNotFound, map[string]interface{}{"error": result.Error})
      return
    }

    if user.Role == "admin" {
      c.JSON(http.StatusUnauthorized, map[string]interface{}{"error": "Only admins can get data provider auth keys"})
      return
    }

    providerID, err := strconv.Atoi(c.Param("providerID"))
    if err != nil {
      c.JSON(http.StatusUnauthorized, map[string]interface{}{"error": "Invalid data provider ID"})
      return
    } 
 
    if result := configs.DB.First(&provider, providerID); result.Error != nil || result.RowsAffected != 1 {
      c.JSON(http.StatusNotFound, map[string]interface{}{"error": result.Error})
      return
    }

    token, err := middleware.GenerateDataProviderToken(provider.ID)
    if  err != nil {
      c.JSON(http.StatusInternalServerError, map[string]interface{}{"stage": "generate token", "error": err.Error()})
      return
    }

    c.JSON(http.StatusOK, map[string]interface{}{"token": token})
  }
}
