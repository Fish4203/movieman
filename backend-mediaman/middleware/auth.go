package middleware

import (
	"errors"
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

	token_lifespan, err := strconv.Atoi(os.Getenv("TOKEN_HOUR_LIFESPAN"))

	if err != nil {
		token_lifespan = 4
	}

	claims := jwt.MapClaims{}
	claims["authorized"] = true
  claims["user_id"] = userID
	claims["user_role"] = userRole
  claims["exp"] = time.Now().Add(time.Hour * time.Duration(token_lifespan)).Unix()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString([]byte(os.Getenv("API_SECRET")))
}

func GenerateDataProviderToken(providerID uint) (string, error) {
	claims := jwt.MapClaims{}
	claims["authorized"] = true
  claims["provider_id"] = providerID
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString([]byte(os.Getenv("API_SECRET")))
}

func UserAuth() gin.HandlerFunc {
  return func(c *gin.Context) {
    claims, err := tokenExtractinator(c)
    if err != nil {
      c.AbortWithStatusJSON(http.StatusUnauthorized, map[string]interface{}{"error": err.Error()})
      return
    }
    
    userRole, roleExists := (*claims)["role"].(string)
    userID, userExists := (*claims)["user_id"].(float64)
    
    if !roleExists || !userExists {
      c.AbortWithStatusJSON(http.StatusUnauthorized, map[string]interface{}{"error": "can not get token claims", "userID": userExists, "userRole": roleExists})
      return
    }

    c.Set("userRole", userRole)
    c.Set("userID", uint(userID))
		c.Next()
  }
}

func DataProviderAuth() gin.HandlerFunc {
  return func(c *gin.Context) {
    claims, err := tokenExtractinator(c)
    if err != nil {
      c.AbortWithStatusJSON(http.StatusUnauthorized, map[string]interface{}{"error": err.Error()})
      return
    }
    
    providerID, providerExists := (*claims)["provider_id"].(float64)
    
    if !providerExists {
      c.AbortWithStatusJSON(http.StatusUnauthorized, map[string]interface{}{"error": "can not get token claims", "providerID": providerExists})
      return
    }

    c.Set("providerID", uint(providerID))
		c.Next()
  }
}

func tokenExtractinator(c *gin.Context) (*jwt.MapClaims, error) {
	bearerToken := c.Request.Header.Get("Authorization")
	if len(strings.Split(bearerToken, " ")) != 2 {
		return nil, errors.New("can not find auth token")
	}
	
  tokenString := strings.Split(bearerToken, " ")[1]
  token, parseError := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
	  if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(os.Getenv("API_SECRET")), nil
	})

  if parseError != nil {
    return nil, parseError
  }

  claims, ok := token.Claims.(jwt.MapClaims)
  if ok && token.Valid {
    return &claims, nil
  }

  return nil, errors.New("Unable to parse jwt token") 
}

