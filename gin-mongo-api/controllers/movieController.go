package controllers

import (
	"fmt"
	"gin-mongo-api/configs"
	"gin-mongo-api/middleware"
	"gin-mongo-api/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func CreateUser() gin.HandlerFunc {
  return func(c *gin.Context) {
    var user models.User

    //validate the request body
    if err := c.BindJSON(&user); err != nil {
      c.JSON(http.StatusBadRequest, map[string]interface{}{"error": err.Error()})
      return
    }

    password, err := bcrypt.GenerateFromPassword([]byte(user.Password),bcrypt.DefaultCost)
    if err != nil {
      c.JSON(http.StatusBadRequest, map[string]interface{}{"error": err.Error()})
      return
    }
    user.Password = string(password)

    if result := configs.DB.Create(&user); result.Error != nil {
      c.JSON(http.StatusInternalServerError, map[string]interface{}{"error": result.Error})
      return
    }
        
    user.Password = ""
    c.JSON(http.StatusCreated, map[string]interface{}{"user": user})
  }
}



func GetAUser() gin.HandlerFunc {
  return func(c *gin.Context) {
    userID := c.Param("userId")
    var user models.User

    if result := configs.DB.First(&user, userID); result.Error != nil || result.RowsAffected != 1 {
      c.JSON(http.StatusNotFound, map[string]interface{}{"error": result.Error})
       return
    }
        
    user.Password = ""
    c.JSON(http.StatusOK, map[string]interface{}{"user": user})
  }
}

func GetUser() gin.HandlerFunc {
  return func(c *gin.Context) {
    userID := c.MustGet("userId").(string)
    var user models.User

    if result := configs.DB.First(&user, userID); result.Error != nil || result.RowsAffected != 1 {
      c.JSON(http.StatusNotFound, map[string]interface{}{"error": result.Error})
       return
    }
        
    user.Password = ""
    c.JSON(http.StatusOK, map[string]interface{}{"user": user})
  }
}


func EditAUser() gin.HandlerFunc {
  return func(c *gin.Context) {
    var oldUser models.User
    var newUser models.User
        
    //validate the request body
    if err := c.BindJSON(&newUser); err != nil {
      c.JSON(http.StatusBadRequest, map[string]interface{}{"error": err.Error()})
      return
    }

    userID := c.MustGet("userId").(string)
    if userID == "" {
      c.JSON(http.StatusUnauthorized, map[string]interface{}{"error": "Invalid jwt or no jwt sent"})
      return
    }

    if result := configs.DB.First(&oldUser, userID); result.Error != nil || result.RowsAffected != 1 {
      c.JSON(http.StatusNotFound, map[string]interface{}{"error": result.Error})
      return
    }
        
    passwordHash, err := bcrypt.GenerateFromPassword([]byte(newUser.Password),bcrypt.DefaultCost)
    if err != nil {
      c.JSON(http.StatusBadRequest, map[string]interface{}{"error": err.Error()})
      return
    }

    oldUser.Password = string(passwordHash)
    oldUser.Name = newUser.Name
    oldUser.Email = newUser.Email
    oldUser.Role = newUser.Role
    
    if result := configs.DB.Save(&oldUser); result.Error != nil {
      c.JSON(http.StatusInternalServerError, map[string]interface{}{"error": result.Error})
      return
    }

    oldUser.Password = ""
    c.JSON(http.StatusOK, map[string]interface{}{"user": oldUser})
  }
}


func DeleteAUser() gin.HandlerFunc {
  return func(c *gin.Context) {
    userID := c.MustGet("userId").(string)
    if userID == "" {
      c.JSON(http.StatusUnauthorized, map[string]interface{}{"error": "Invalid jwt or no jwt sent"})
      return
    }

    if result := configs.DB.Delete(&models.User{}, userID); result.Error != nil {
      c.JSON(http.StatusInternalServerError, map[string]interface{}{"error": result.Error})
      return
    }
    
    c.JSON(http.StatusOK, map[string]interface{}{"result": "User successfully deleted"})
  }
}

func GetAllUsers() gin.HandlerFunc {
  return func(c *gin.Context) {
    var users []models.User

    if result := configs.DB.Find(&users); result.Error != nil {
      c.JSON(http.StatusInternalServerError, map[string]interface{}{"error": result.Error})
      return
    }
    
    for i := 0; i < len(users); i++ {
      users[i].Password = ""
    }

    c.JSON(http.StatusOK, map[string]interface{}{"users": users})
  }
}


func Login() gin.HandlerFunc {
  return func(c *gin.Context) {
    var user models.User
    var dbUser models.User

    //validate the request body
    if err := c.BindJSON(&user); err != nil {
      c.JSON(http.StatusBadRequest, map[string]interface{}{"error": err.Error()})
      return
    }

    if result := configs.DB.First(&dbUser, "name = ?", user.Name); result.Error != nil || result.RowsAffected != 1 {
      c.JSON(http.StatusNotFound, map[string]interface{}{"error": result.Error})
      return
    }

    if err := bcrypt.CompareHashAndPassword([]byte(dbUser.Password), []byte(user.Password)); err != nil {
      c.JSON(http.StatusInternalServerError, map[string]interface{}{"error": err.Error()})
      return
    }

    token, err := middleware.GenerateToken(fmt.Sprint(dbUser.ID))
    if  err != nil {
      c.JSON(http.StatusInternalServerError, map[string]interface{}{"error": err.Error()})
      return
    }

    c.JSON(http.StatusOK, map[string]interface{}{"token": token})
  }
}
