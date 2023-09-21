package middleware

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
	"net/http"


	"github.com/gin-gonic/gin"
	jwt "github.com/golang-jwt/jwt/v5"
)

func GenerateToken(user_id string) (string, error) {

	token_lifespan,err := strconv.Atoi(os.Getenv("TOKEN_HOUR_LIFESPAN"))

	if err != nil {
		return "",err
	}

	claims := jwt.MapClaims{}
	claims["authorized"] = true
	claims["user_id"] = user_id
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
		c.Set("userId", "")

		tokenString := ExtractToken(c)

		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
			}
			return []byte(os.Getenv("API_SECRET")), nil
		})
		if err == nil {
			claims, ok := token.Claims.(jwt.MapClaims)
			// fmt.Println(claims)
			if ok && token.Valid {
				uid, exists := claims["user_id"].(string)
				if exists {
					c.Set("userId", uid)
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
