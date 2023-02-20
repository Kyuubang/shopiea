package handlers

import (
	"github.com/Kyuubang/shopiea/db"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func CreateUser(c *gin.Context) {
	// Bind the JSON payload to a User struct
	var user db.User
	if err := c.BindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Bad Request",
		})
		return
	}

	// create user
	err := db.CreateUser(user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"user":    user.Username,
		"message": "successful create user",
	})
}

func GetUsers(c *gin.Context) {
	classId, err := strconv.Atoi(c.Query("class_id"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Internal Server Error",
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
		"class":   className,
		"student": users,
	})
	return
}
