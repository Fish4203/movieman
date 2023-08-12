package controllers

import (
    "context"
    // "gin-mongo-api/configs"
    "gin-mongo-api/models"
    // "gin-mongo-api/responses"
    "gin-mongo-api/middleware"
    "net/http"
    "time"

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
        ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
        var user models.User
        defer cancel()

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

        hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password),bcrypt.DefaultCost)
        if err != nil {
            c.JSON(http.StatusBadRequest, map[string]interface{}{"error": err.Error()})
            return
        }

        newUser := models.User{
            Id:       primitive.NewObjectID(),
            Name:     user.Name,
            Password: string(hashedPassword),
            Role:     user.Role,
        }

        result, err := models.UserCollection.InsertOne(ctx, newUser)
        if err != nil {
            c.JSON(http.StatusInternalServerError, map[string]interface{}{"error": err.Error()})
            return
        }

        user.Password = ""
        c.JSON(http.StatusCreated, map[string]interface{}{"data": result})
    }
}



func GetAUser() gin.HandlerFunc {
    return func(c *gin.Context) {
        ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
        userId := c.Param("userId")
        var user models.User
        defer cancel()

        objId, _ := primitive.ObjectIDFromHex(userId)
        err := models.UserCollection.FindOne(ctx, bson.M{"_id": objId}).Decode(&user)
        if err != nil {
            c.JSON(http.StatusNotFound, map[string]interface{}{"error": err.Error()})
            return
        }

        user.Password = ""
        c.JSON(http.StatusOK, map[string]interface{}{"user": user})
    }
}


func EditAUser() gin.HandlerFunc {
    return func(c *gin.Context) {
        ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
        var user models.User
        defer cancel()

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

        hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password),bcrypt.DefaultCost)
        if err != nil {
            c.JSON(http.StatusBadRequest, map[string]interface{}{"error": err.Error()})
            return
        }

        update := bson.M{"name": user.Name, "password": string(hashedPassword), "role": user.Role}
        result, err := models.UserCollection.UpdateOne(ctx, bson.M{"_id": objId}, bson.M{"$set": update})
        if err != nil {
            c.JSON(http.StatusInternalServerError, map[string]interface{}{"error": err.Error()})
            return
        }

        //get updated user details
        var updatedUser models.User
        if result.MatchedCount != 1 {
            c.JSON(http.StatusNotFound, map[string]interface{}{"error": "User not found"})
            return
        }

        err = models.UserCollection.FindOne(ctx, bson.M{"_id": objId}).Decode(&updatedUser)
        if err != nil {
            c.JSON(http.StatusInternalServerError, map[string]interface{}{"error": err.Error()})
            return
        }

        updatedUser.Password = ""
        c.JSON(http.StatusOK, map[string]interface{}{"user": updatedUser})
    }
}


func DeleteAUser() gin.HandlerFunc {
    return func(c *gin.Context) {
        ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
        defer cancel()

        userId := c.MustGet("userId").(string)
        if userId == "" {
            c.JSON(http.StatusUnauthorized, map[string]interface{}{"error": "Invalid jwt or no jwt sent"})
            return
        }

        objId, _ := primitive.ObjectIDFromHex(userId)

        result, err := models.UserCollection.DeleteOne(ctx, bson.M{"_id": objId})
        if err != nil {
            c.JSON(http.StatusNotFound, map[string]interface{}{"error": err.Error()})
            return
        }

        if result.DeletedCount < 1 {
            c.JSON(http.StatusNotFound, map[string]interface{}{"error": "User not found"})
            return
        }

        c.JSON(http.StatusOK, map[string]interface{}{"result": "User successfully deleted"})
    }
}

func GetAllUsers() gin.HandlerFunc {
    return func(c *gin.Context) {
        ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
        var users []models.User
        defer cancel()

        results, err := models.UserCollection.Find(ctx, bson.M{})

        if err != nil {
            c.JSON(http.StatusInternalServerError, map[string]interface{}{"error": err.Error()})
            return
        }

        //reading from the db in an optimal way
        defer results.Close(ctx)
        for results.Next(ctx) {
            var singleUser models.User
            if err = results.Decode(&singleUser); err != nil {
                c.JSON(http.StatusInternalServerError, map[string]interface{}{"error": err.Error()})
            }

            singleUser.Password = ""
            users = append(users, singleUser)
        }

        c.JSON(http.StatusOK, map[string]interface{}{"users": users})
    }
}


func Login() gin.HandlerFunc {
    return func(c *gin.Context) {
        ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
        var user models.User
        defer cancel()

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

        password := user.Password

        err := models.UserCollection.FindOne(ctx, bson.M{"name": user.Name}).Decode(&user)
        if err != nil {
            c.JSON(http.StatusNotFound, map[string]interface{}{"error": err.Error()})
            return
        }


        if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
            c.JSON(http.StatusInternalServerError, map[string]interface{}{"error": err.Error()})
            return
        }

        token, err := middleware.GenerateToken(user.Id.Hex())
        if  err != nil {
            c.JSON(http.StatusInternalServerError, map[string]interface{}{"error": err.Error()})
            return
        }

        c.JSON(http.StatusOK, map[string]interface{}{"token": token})
    }
}
