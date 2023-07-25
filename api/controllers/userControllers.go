package controllers

import (
	"net/http"
	"os"
	"time"

	"github.com/whoisaditya/golang-task-management-system/api/initializers"
	"github.com/whoisaditya/golang-task-management-system/api/models"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

func UserRegistration(c *gin.Context) {
	// get email/password from request body
	var body struct {
		Username string `json:"username" binding:"required"`
		Email    string `json:"email" binding:"required"`
		Password string `json:"password" binding:"required"`
	}

	if c.Bind(&body) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Fields are empty",
		})
		return
	}

	// Check if User is already exists
	var user models.User
	initializers.DB.First(&user, "email = ?", body.Email)

	if user.ID != 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "User already exixts",
		})
		return
	}

	// hash the password
	hash, err := bcrypt.GenerateFromPassword([]byte(body.Password), 10)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Error hashing password",
		})
		return
	}

	// create user
	newUser := models.User{Username: body.Username, Email: body.Email, Password: string(hash)}
	result := initializers.DB.Create(&newUser)

	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Error creating user",
		})
		return
	}

	// respond
	c.JSON(http.StatusOK, gin.H{
		"message": "User created successfully",
	})
}

func UserLogin(c *gin.Context) {
	// get email/password from request body
	var body struct {
		Email    string
		Password string
	}

	if c.Bind(&body) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Fields are empty",
		})
		return
	}

	// find user by email
	var user models.User
	initializers.DB.First(&user, "email = ?", body.Email)

	if user.ID == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "User not found",
		})
		return
	}

	// compare password
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(body.Password))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid password",
		})
		return
	}

	// generate jwt
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": user.ID,
		"exp": time.Now().Add(time.Hour * 24).Unix(),
	})

	// Sign and get the complete encoded token as a string using the secret
	tokenString, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Error generating jwt",
		})
		return
	}

	// respond with jwt
	c.SetSameSite(http.SameSiteNoneMode)
	c.SetCookie("jwt", tokenString, 3600*24*30, "", "", false, true)

	c.JSON(http.StatusOK, gin.H{
		"message": "Logged in successfully",
	})
}

func UserLogout(c *gin.Context) {
	// Clear the JWT cookie on the client-side
	c.SetSameSite(http.SameSiteNoneMode)
	c.SetCookie("jwt", "", -1, "", "", false, true)

	c.JSON(http.StatusOK, gin.H{
		"message": "Logged out successfully",
	})
}

func UserDeletion(c *gin.Context) {
	user_id := c.GetUint("user_id") // Assuming you have a middleware that sets the user ID in the context

	// Clear the JWT cookie on the client-side
	c.SetSameSite(http.SameSiteNoneMode)
	c.SetCookie("jwt", "", -1, "", "", false, true)

	// Delete the user
	err := initializers.DB.Delete(&models.User{}, user_id).Error
    if err!= nil {
        c.JSON(http.StatusInternalServerError, gin.H{
            "error": "Failed to delete user",
        })
        return
    }

	c.JSON(http.StatusOK, gin.H{
		"message": "User deleted successfully",
	})
}
