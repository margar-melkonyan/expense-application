package middleware

import (
	"encoding/json"
	"expense-application/internal/db"
	"expense-application/internal/model"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"log/slog"
	"net/http"
	"os"
	"strings"
	"time"
)

func RequireAuth(c *gin.Context) {
	authorizationHeader := c.GetHeader("Authorization")
	if authorizationHeader == "" {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{})
		return
	}

	tokenParts := strings.Split(authorizationHeader, "Bearer ")
	if len(tokenParts) != 2 {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{})
		return
	}

	token := tokenParts[1]

	parsedToken, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		return []byte(os.Getenv("JWT_ACCESS_TOKEN_SECRET")), nil
	})

	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{})
		return
	}

	if claims, ok := parsedToken.Claims.(jwt.Claims); ok && parsedToken.Valid {
		exp, err := claims.GetExpirationTime()
		if err != nil {
			slog.Error(err.Error())
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{})
			return
		}

		if time.Now().Unix() > exp.Time.Unix() {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{})
			return
		}

		sub, err := claims.GetSubject()
		if err != nil {
			slog.Error(err.Error())
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{})
			return
		}

		var user model.User
		err = json.Unmarshal([]byte(sub), &user)
		if err != nil {
			slog.Error(err.Error())
		}

		postgresDB, err := db.NewPostgresDB()
		if err != nil {
			return
		}

		if postgresDB.Preload("Roles").First(&user).Error != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"message": "User doesn't have access to this resource",
			})
		}

		if len(user.Roles) != 0 {
			_ = json.Unmarshal(user.Roles[0].Permissions, &user.Roles[0].PermissionsUnmarshalled)
		} else {
			c.AbortWithStatusJSON(http.StatusForbidden, gin.H{})
			return
		}

		c.Set("user", user)
	}

	c.Next()
}
