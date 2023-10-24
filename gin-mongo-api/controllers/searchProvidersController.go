package controllers

import (
	"gin-mongo-api/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetProviders() gin.HandlerFunc {
	return func(c *gin.Context) {

		provs, err := models.FindProvs()
		if err != nil {
			c.JSON(http.StatusInternalServerError, map[string]interface{}{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, provs)
	}
}

func CreateProvider() gin.HandlerFunc {
	return func(c *gin.Context) {

		var prov models.Provider

		if err := c.BindJSON(&prov); err != nil {
			c.JSON(http.StatusBadRequest, map[string]interface{}{"error": err.Error()})
			return
		}

		//use the validator library to validate required fields
		if err := validate.Struct(&prov); err != nil {
			c.JSON(http.StatusBadRequest, map[string]interface{}{"error": err.Error()})
			return
		}

		err := models.WriteProv(prov)
		if err != nil {
			c.JSON(http.StatusInternalServerError, map[string]interface{}{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, "success")
	}
}

func DeleteProvider() gin.HandlerFunc {
	return func(c *gin.Context) {
		query := c.Param("query")
		err := models.DeleteProv([]string{query})
		if err != nil {
			c.JSON(http.StatusInternalServerError, map[string]interface{}{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, "success")
	}
}
