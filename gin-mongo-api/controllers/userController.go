package controllers

import (
    // "context"
    // "gin-mongo-api/configs"
    "gin-mongo-api/models"
    // "gin-mongo-api/responses"
    "gin-mongo-api/middleware"
    "net/http"
    // "time"
    // "fmt"

    "github.com/gin-gonic/gin"
    "github.com/go-playground/validator/v10"
    "go.mongodb.org/mongo-driver/bson"
    "go.mongodb.org/mongo-driver/bson/primitive"
    // "go.mongodb.org/mongo-driver/mongo"
    "golang.org/x/crypto/bcrypt"
)


var validate = validator.New()

func CreateUser() gin.HandlerFunc {
    return func(c *gin.Context) {
        var user models.User
        var err error

        //validate the request body
        if err := c.BindJSON(&user); err != nil {
            c.JSON(http.StatusBadRequest, map[string]interface{}{"error": err.Error()})
            return
        }

        //use the validator library to validate required fields
        if err := validate.Struct(&user); err != nil {
            c.JSON(http.StatusBadRequest, map[string]interface{}{"error": err.Error()})
            return
        }

        res, _ := models.FindUser(bson.D{{"name", user.Name}})
        if len(res) != 0 {
            c.JSON(http.StatusBadRequest, map[string]interface{}{"error": "User already exists"})
            return
        }

        password, err := bcrypt.GenerateFromPassword([]byte(user.Password),bcrypt.DefaultCost)
        if err != nil {
            c.JSON(http.StatusBadRequest, map[string]interface{}{"error": err.Error()})
            return
        }
        user.Password = string(password)

        err = models.WriteUser([]models.User{user})
        if err != nil {
            c.JSON(http.StatusInternalServerError, map[string]interface{}{"error": err.Error()})
            return
        }

        res, _ = models.FindUser(bson.D{{"name", user.Name}})
        user = res[0]
        
        user.Password = ""
        c.JSON(http.StatusCreated, map[string]interface{}{"user": user})
    }
}



func GetAUser() gin.HandlerFunc {
    return func(c *gin.Context) {
        userName := c.Param("userId")

        res, err := models.FindUser(bson.D{{"name", userName}})
        if err != nil {
            c.JSON(http.StatusNotFound, map[string]interface{}{"error": err.Error()})
            return
        }
        if len(res) == 0 {
            c.JSON(http.StatusNotFound, map[string]interface{}{"error": "cant find user"})
            return
        }
        
        user := res[0]
        user.Password = ""
        c.JSON(http.StatusOK, map[string]interface{}{"user": user})
    }
}


func EditAUser() gin.HandlerFunc {
    return func(c *gin.Context) {
        var user models.User

        //validate the request body
        if err := c.BindJSON(&user); err != nil {
            c.JSON(http.StatusBadRequest, map[string]interface{}{"error": err.Error()})
            return
        }
        //use the validator library to validate required fields
        if err := validate.Struct(&user); err != nil {
            c.JSON(http.StatusBadRequest, map[string]interface{}{"error": err.Error()})
            return
        }

        userId := c.MustGet("userId").(string)
        if userId == "" {
            c.JSON(http.StatusUnauthorized, map[string]interface{}{"error": "Invalid jwt or no jwt sent"})
            return
        }

        objId, _ := primitive.ObjectIDFromHex(userId)
        
        passwordHash, err := bcrypt.GenerateFromPassword([]byte(user.Password),bcrypt.DefaultCost)
        if err != nil {
            c.JSON(http.StatusBadRequest, map[string]interface{}{"error": err.Error()})
            return
        }

        user.Password = string(passwordHash)
        
        res, _ := models.FindUser(bson.D{{"_id", objId}})
        if len(res) == 0 || res[0].Name != user.Name {
            c.JSON(http.StatusUnauthorized, map[string]interface{}{"error": "Username Mismatch"})
            return
        }
        
        err = models.WriteUser([]models.User{user})
        if err != nil {
            c.JSON(http.StatusInternalServerError, map[string]interface{}{"error": err.Error()})
            return
        }

        user.Id = res[0].Id
        user.Password = ""
        c.JSON(http.StatusOK, map[string]interface{}{"user": user})
    }
}


func DeleteAUser() gin.HandlerFunc {
    return func(c *gin.Context) {
        userId := c.MustGet("userId").(string)
        if userId == "" {
            c.JSON(http.StatusUnauthorized, map[string]interface{}{"error": "Invalid jwt or no jwt sent"})
            return
        }

        err := models.DeleteUser([]string{userId})
        if err != nil {
            c.JSON(http.StatusNotFound, map[string]interface{}{"error": err.Error()})
            return
        }

        c.JSON(http.StatusOK, map[string]interface{}{"result": "User successfully deleted"})
    }
}

func GetAllUsers() gin.HandlerFunc {
    return func(c *gin.Context) {
        users, err := models.FindUser(bson.D{})
        if err != nil {
            c.JSON(http.StatusInternalServerError, map[string]interface{}{"error": err.Error()})
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

        //validate the request body
        if err := c.BindJSON(&user); err != nil {
            c.JSON(http.StatusBadRequest, map[string]interface{}{"error": err.Error()})
            return
        }

        //use the validator library to validate required fields
        if err := validate.Struct(&user); err != nil {
            c.JSON(http.StatusBadRequest, map[string]interface{}{"error": err.Error()})
            return
        }

        res, err := models.FindUser(bson.D{{"name", user.Name}})
        if err != nil {
            c.JSON(http.StatusNotFound, map[string]interface{}{"error": err.Error()})
            return
        }


        if err := bcrypt.CompareHashAndPassword([]byte(res[0].Password), []byte(user.Password)); err != nil {
            c.JSON(http.StatusInternalServerError, map[string]interface{}{"error": err.Error()})
            return
        }

        token, err := middleware.GenerateToken(res[0].Id.Hex())
        if  err != nil {
            c.JSON(http.StatusInternalServerError, map[string]interface{}{"error": err.Error()})
            return
        }

        c.JSON(http.StatusOK, map[string]interface{}{"token": token})
    }
}
