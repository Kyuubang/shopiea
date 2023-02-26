package handlers

import (
	"github.com/Kyuubang/shopiea/db"
	"github.com/gin-gonic/gin"
	"net/http"
	"regexp"
	"strconv"
)

func isValidUsername(s string) bool {
	re := regexp.MustCompile(`^[a-zA-Z0-9]{1,16}$`)
	return re.MatchString(s)
}

func CreateUser(c *gin.Context) {
	// Bind the JSON payload to a User struct
	var user db.User
	if err := c.BindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Bad Request",
		})
		return
	}

	// check if username contains only alphanumeric
	if !isValidUsername(user.Username) {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "username must contain only alphanumeric",
		})
		return
	}

	// create user
	res, err := db.CreateUser(user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"username": user.Username,
		"name":     user.Name,
		"id":       res.ID,
		"message":  "successful create user",
	})
}

func GetUsers(c *gin.Context) {
	classId, err := strconv.Atoi(c.Query("class_id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "invalid class_id",
		})
		return
	}

	className, err := db.GetClassNameById(classId)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "Class Not Found",
		})
		return
	}

	users, err := db.GetUsersByClassId(classId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Internal Server Error",
		})
		return
	}

	if len(users) == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "Users Not Found",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"class":    className,
		"students": users,
	})
	return
}

// DeleteUser is a function to delete user
func DeleteUser(c *gin.Context) {
	userId := c.Query("id")
	if userId == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Bad Request",
		})
		return
	}

	// convert string to int
	userIdInt, err := strconv.Atoi(userId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Bad Request",
		})
		return
	}

	err = db.DeleteUsersByUsersId(userIdInt)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Internal Server Error",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Success delete user",
	})
	return
}

// UpdateUser is a function to update user
func UpdateUser(c *gin.Context) {
	// Bind the JSON payload to a User struct
	var user db.User
	if err := c.BindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Bad Request",
		})
		return
	}

	userId := c.Query("user_id")
	if userId == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Bad Request",
		})
		return
	}

	// convert string to int
	userIdInt, err := strconv.Atoi(userId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Bad Request",
		})
		return
	}

	// update user
	err = db.UpdateUsersByUsersId(userIdInt, user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Internal Server Error",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Success update user",
	})
	return
}
