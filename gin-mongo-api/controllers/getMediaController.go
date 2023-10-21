package controllers

import (
	"gin-mongo-api/models"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func GetMediaSearch() gin.HandlerFunc {
	return func(c *gin.Context) {
		query := c.Query("q")
		types := c.Query("types")
		var out = make(map[string]interface{})

		if strings.ContainsAny(types, "m") {
			movies, err := models.FindMovie(bson.D{{"title", bson.D{{"$regex", query}, {"$options", "i"}}}})
			if err != nil {
				c.JSON(http.StatusInternalServerError, map[string]interface{}{"error": err.Error()})
				return
			}
			out["movies"] = movies
		}

		if strings.ContainsAny(types, "s") {
			shows, err := models.FindShow(bson.D{{"title", bson.D{{"$regex", query}, {"$options", "i"}}}})
			if err != nil {
				c.JSON(http.StatusInternalServerError, map[string]interface{}{"error": err.Error()})
				return
			}
			out["shows"] = shows
		}

		if strings.ContainsAny(types, "e") {
			episodes, err := models.FindShowEpisode(bson.D{{"title", bson.D{{"$regex", query}, {"$options", "i"}}}})
			if err != nil {
				c.JSON(http.StatusInternalServerError, map[string]interface{}{"error": err.Error()})
				return
			}
			out["episodes"] = episodes
		}

		if strings.ContainsAny(types, "p") {
			persons, err := models.FindPerson(bson.D{{"name", bson.D{{"$regex", query}, {"$options", "i"}}}})
			if err != nil {
				c.JSON(http.StatusInternalServerError, map[string]interface{}{"error": err.Error()})
				return
			}
			out["people"] = persons
		}

		if strings.ContainsAny(types, "b") {
			books, err := models.FindBook(bson.D{{Key: "title", Value: bson.D{{Key: "$regex", Value: query}, {Key: "$options", Value: "i"}}}})
			if err != nil {
				c.JSON(http.StatusInternalServerError, map[string]interface{}{"error": err.Error()})
				return
			}
			out["books"] = books
		}

		if strings.ContainsAny(types, "v") {
			games, err := models.FindGame(bson.D{{"title", bson.D{{"$regex", query}, {"$options", "i"}}}})
			if err != nil {
				c.JSON(http.StatusInternalServerError, map[string]interface{}{"error": err.Error()})
				return
			}
			out["games"] = games
		}

		if strings.ContainsAny(types, "g") {
			groups, err := models.FindGroup(bson.D{{"title", bson.D{{"$regex", query}, {"$options", "i"}}}})
			if err != nil {
				c.JSON(http.StatusInternalServerError, map[string]interface{}{"error": err.Error()})
				return
			}
			out["groups"] = groups
		}

		if strings.ContainsAny(types, "m") {
			companies, err := models.FindCompany(bson.D{{"name", bson.D{{"$regex", query}, {"$options", "i"}}}})
			if err != nil {
				c.JSON(http.StatusInternalServerError, map[string]interface{}{"error": err.Error()})
				return
			}
			out["companies"] = companies
		}

		c.JSON(http.StatusOK, out)
	}
}

func MovieDetails() gin.HandlerFunc {
	return func(c *gin.Context) {
		query := c.Param("movieId")
		objId, _ := primitive.ObjectIDFromHex(query)

		movies, err := models.FindMovie(bson.D{{"_id", objId}})
		if err != nil {
			c.JSON(http.StatusInternalServerError, map[string]interface{}{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, map[string]interface{}{"movies": movies})
	}
}

func ShowDetails() gin.HandlerFunc {
	return func(c *gin.Context) {
		query := c.Param("showId")
		objId, _ := primitive.ObjectIDFromHex(query)

		show, err := models.FindShow(bson.D{{"_id", objId}})
		if err != nil {
			c.JSON(http.StatusInternalServerError, map[string]interface{}{"error": err.Error()})
			return
		}

		seasons, err := models.FindShowSeason(bson.D{{"showId", objId}})
		if err != nil {
			c.JSON(http.StatusInternalServerError, map[string]interface{}{"error": err.Error()})
			return
		}

		episodes, err := models.FindShowEpisode(bson.D{{"showId", objId}})
		if err != nil {
			c.JSON(http.StatusInternalServerError, map[string]interface{}{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, map[string]interface{}{"shows": show, "seasons": seasons, "episodes": episodes})
	}
}

func BookDetails() gin.HandlerFunc {
	return func(c *gin.Context) {
		query := c.Param("Id")
		objId, _ := primitive.ObjectIDFromHex(query)

		obj, err := models.FindBook(bson.D{{Key: "_id", Value: objId}})
		if err != nil {
			c.JSON(http.StatusInternalServerError, map[string]interface{}{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, map[string]interface{}{"books": obj})
	}
}

func GameDetails() gin.HandlerFunc {
	return func(c *gin.Context) {
		query := c.Param("Id")
		objId, _ := primitive.ObjectIDFromHex(query)

		obj, err := models.FindGame(bson.D{{Key: "_id", Value: objId}})
		if err != nil {
			c.JSON(http.StatusInternalServerError, map[string]interface{}{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, map[string]interface{}{"games": obj})
	}
}

func PersonDetails() gin.HandlerFunc {
	return func(c *gin.Context) {
		query := c.Param("personId")
		objId, _ := primitive.ObjectIDFromHex(query)

		results, err := models.FindPerson(bson.D{{"_id", objId}})
		if err != nil {
			c.JSON(http.StatusInternalServerError, map[string]interface{}{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, map[string]interface{}{"people": results})
	}
}

func GroupDetails() gin.HandlerFunc {
	return func(c *gin.Context) {
		query := c.Param("Id")
		objId, _ := primitive.ObjectIDFromHex(query)

		obj, err := models.FindGroup(bson.D{{"_id", objId}})
		if err != nil {
			c.JSON(http.StatusInternalServerError, map[string]interface{}{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, map[string]interface{}{"groups": obj})
	}
}

func CompanyDetails() gin.HandlerFunc {
	return func(c *gin.Context) {
		query := c.Param("Id")
		objId, _ := primitive.ObjectIDFromHex(query)

		obj, err := models.FindCompany(bson.D{{"_id", objId}})
		if err != nil {
			c.JSON(http.StatusInternalServerError, map[string]interface{}{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, map[string]interface{}{"companies": obj})
	}
}
