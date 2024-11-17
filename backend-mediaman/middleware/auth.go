package middleware

import (
	"fmt"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	jwt "github.com/golang-jwt/jwt/v5"
)

func GenerateUserToken(userID uint, userRole string) (string, error) {

	token_lifespan,err := strconv.Atoi(os.Getenv("TOKEN_HOUR_LIFESPAN"))

	if err != nil {
		return "",err
	}

	claims := jwt.MapClaims{}
	claims["authorized"] = true
	claims["type"] = "user"
  claims["user_id"] = userID
	claims["role"] = userRole
  claims["exp"] = time.Now().Add(time.Hour * time.Duration(token_lifespan)).Unix()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString([]byte(os.Getenv("API_SECRET")))
}

func GenerateDataProviderToken(providerID uint) (string, error) {
	claims := jwt.MapClaims{}
	claims["authorized"] = true
	claims["type"] = "dataProvider"
  claims["provider_id"] = providerID
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString([]byte(os.Getenv("API_SECRET")))
}

func ExtractToken(c *gin.Context) string {
	token := c.Query("token")
	if token != "" {
		return token
	}
	bearerToken := c.Request.Header.Get("Authorization")
	if len(strings.Split(bearerToken, " ")) == 2 {
		return strings.Split(bearerToken, " ")[1]
	}
	return ""
}

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set("userID", -1)
    c.Set("userRole", "")
    c.Set("providerID", -1)

		tokenString := ExtractToken(c)

		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
			}
			return []byte(os.Getenv("API_SECRET")), nil
		})

		if err == nil {
			claims, ok := token.Claims.(jwt.MapClaims)
			if ok && token.Valid {
        tokenType, exists := claims["type"].(string)

        if exists && tokenType == "user" {
          userRole, roleExists := claims["role"].(string)
          userID, userExists := claims["user_id"].(uint)
          
          if roleExists && userExists {
					  c.Set("userRole", userRole)
            c.Set("userID", uint(userID))
					  c.Next()
				  } else {
					  c.AbortWithStatusJSON(http.StatusUnauthorized, map[string]interface{}{"error": "Invalid jwt could not find user in jwt"})
					  return
				  }
        } else if exists && tokenType == "dataPrivider" {
          providerID, providerExists := claims["provider_id"].(uint)
          
          if providerExists {
					  c.Set("providerID", providerID)
					  c.Next()
				  } else {
					  c.AbortWithStatusJSON(http.StatusUnauthorized, map[string]interface{}{"error": "Invalid jwt could not find data provider in jwt"})
					  return
				  }
        } else {
					c.AbortWithStatusJSON(http.StatusUnauthorized, map[string]interface{}{"error": "Invalid jwt could not find type in jwt"})
					return
				}
			} else {
				c.AbortWithStatusJSON(http.StatusUnauthorized, map[string]interface{}{"error": "Invalid jwt"})
				return
			}
		} else {
			c.AbortWithStatusJSON(http.StatusUnauthorized, map[string]interface{}{"error": "Invalid jwt or no jwt sent"})
			return
		}
	}
}
