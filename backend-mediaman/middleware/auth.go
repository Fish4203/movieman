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

func GenerateToken(user_id uint, user_role string) (string, error) {

	token_lifespan,err := strconv.Atoi(os.Getenv("TOKEN_HOUR_LIFESPAN"))

	if err != nil {
		return "",err
	}

	claims := jwt.MapClaims{}
	claims["authorized"] = true
	claims["user_id"] = user_id
	claims["role"] = user_role
  claims["exp"] = time.Now().Add(time.Hour * time.Duration(token_lifespan)).Unix()
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
		c.Set("userId", 0)
    c.Set("userRole", "")

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
				userRole, _ := claims["role"].(string)
        uid, exists := claims["user_id"].(float64)
				fmt.Println(claims)
        if exists {
					c.Set("userRole", userRole)
          c.Set("userId", uint(uid))
					c.Next()
				} else {
					c.AbortWithStatusJSON(http.StatusUnauthorized, map[string]interface{}{"error": "Invalid jwt could not find userid in jwt"})
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