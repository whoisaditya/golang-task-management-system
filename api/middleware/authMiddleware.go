package middleware

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/whoisaditya/golang-task-management-system/api/initializers"
	"github.com/whoisaditya/golang-task-management-system/api/models"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)


func AuthMiddleware(c *gin.Context) {
	fmt.Println("In AuthMiddleware middleware")

	// Get cookie from request
	tokenString, err := c.Cookie("jwt")

	if err != nil {
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	// Validate cookie
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {

		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
	
		return []byte(os.Getenv("JWT_SECRET")), nil
	})

	if err!= nil {
        c.AbortWithStatus(http.StatusUnauthorized)
        return
    }
	
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		// Checking expiration
		if float64(time.Now().Unix()) > claims["exp"].(float64)  {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		// Find user with token subject
		var user models.User
		initializers.DB.First(&user, "id = ?", claims["sub"])

		if user.ID == 0 {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		// Attach user to req
		// c.Set("user", user)
		c.Set("user_id", user.ID)

		// Continue
		c.Next()
	} else {
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}
}


