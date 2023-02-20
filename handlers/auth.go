package handlers

import (
	"github.com/Kyuubang/shopiea/db"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
	"time"
)

// Define the key to sign the JWT
var jwtKey = []byte("my_secret_key")

func Login(c *gin.Context) {
	// Bind the JSON payload to a Login struct
	var login db.Login
	if err := c.BindJSON(&login); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Bad Request",
		})
		return
	}

	// login user
	if userId, success := db.ValidationUserLogin(login); success {
		// Generate a JWT token for the user ID
		token, err := generateJWT(userId)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": "Internal Server Error",
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"message": "Success login!",
			"token":   token,
		})
		return
	} else {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "Invalid username or password",
		})
		return
	}

}

// AuthMiddleware function to authenticate JWT token
func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Get the JWT token from the Authorization header
		tokenString := c.GetHeader("Authorization")
		if tokenString == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Authorization header missing"})
			return
		}

		tokens := strings.Split(tokenString, " ")

		// Verify the JWT token and get the user ID from it
		userId, err := verifyJWT(tokens[1])
		// check if token is valid
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid token from middleware"})
			return
		}
		c.Set("userId", userId)

		c.Next()
	}
}

func CheckToken(c *gin.Context) {
	userId := c.MustGet("userId").(string)

	c.JSON(http.StatusOK, gin.H{
		"message": "Success check!",
		"userId":  userId,
		"admin":   db.IsAdmin(userId),
	})
	return
}

func AdminOnly() gin.HandlerFunc {
	return func(c *gin.Context) {
		userId, ok := c.Get("userId")
		if !ok {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"message": "Unauthorized!",
			})
			return
		}

		if !db.IsAdmin(userId.(string)) {
			c.AbortWithStatusJSON(http.StatusForbidden, gin.H{
				"message": "Forbidden!",
			})
			return
		}

		c.Next()
	}
}

// Generate a JWT token for a user ID
func generateJWT(userID string) (string, error) {
	// Set the expiration time for the token to 1 day from now
	expirationTime := time.Now().Add(24 * time.Hour)

	// Create the claims for the token
	claims := &jwt.StandardClaims{
		ExpiresAt: expirationTime.Unix(),
		Subject:   userID,
	}

	// Create the JWT token with the claims and sign it using the key
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	jwtString, err := token.SignedString(jwtKey)
	if err != nil {
		return "", err
	}

	return jwtString, nil
}

// Verify a JWT token and return the user ID if the token is valid
func verifyJWT(jwtString string) (string, error) {
	// Parse the JWT token with the key
	token, err := jwt.ParseWithClaims(jwtString, &jwt.StandardClaims{}, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})

	// Check if the token is valid and has not expired
	if claims, ok := token.Claims.(*jwt.StandardClaims); ok && token.Valid {
		userID := claims.Subject
		return userID, nil
	} else {
		return "", err
	}
}
