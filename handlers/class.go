package handlers

import (
	"errors"
	"fmt"
	"github.com/Kyuubang/shopiea/db"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// CreateClass is a function to create class
func CreateClass(c *gin.Context) {
	var class db.Class

	if err := c.BindJSON(&class); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Bad Request",
		})
		return
	}

	class, err := db.CreateClass(class)
	if err != nil {
		if errors.Is(err, db.ErrAlreadyExist) {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "Labs already exist",
			})
			return
		} else if errors.Is(err, db.ErrCantBeEmpty) {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "Class name cannot be empty",
			})
			return
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": "Internal Server Error",
			})
			return
		}
	}

	c.JSON(http.StatusCreated, gin.H{
		"class":   class.Name,
		"id":      class.ID,
		"message": "Success create class!",
	})
	return
}

// GetClasses is a function to get all class
func GetClasses(c *gin.Context) {
	classes, err := db.GetClasses()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Internal Server Error",
		})
		return
	}

	if len(classes) == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "Classes Not Found",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"classes": classes,
	})
	return
}

// DeleteClass is a function to delete class by class id
func DeleteClass(c *gin.Context) {
	classID := c.Query("class_id")

	// convert string to int
	classIDInt, err := strconv.Atoi(classID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Bad Request",
		})
		return
	}

	err = db.DeleteClassByClassId(classIDInt)
	if err != nil {
		if errors.Is(err, db.ErrNotFound) {
			c.JSON(http.StatusNotFound, gin.H{
				"message": "Class Not Found",
			})
			return
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": "Internal Server Error",
			})
			return
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Success delete class!",
	})
	return
}

// UpdateClass is a function to update class name by class id
func UpdateClass(c *gin.Context) {
	classID := c.Query("class_id")

	// convert string to int
	classIDInt, err := strconv.Atoi(classID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Bad Request",
		})
		return
	}

	var class db.Class

	if err := c.BindJSON(&class); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Bad Request",
		})
		return
	}

	err = db.UpdateClassByClassId(classIDInt, class)
	if err != nil {
		fmt.Println(err)
		switch err {
		case db.ErrNotFound:
			c.JSON(http.StatusNotFound, gin.H{
				"message": "Class Not Found",
			})
			return
		case db.ErrCantBeEmpty:
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "Class name cannot be empty",
			})
			return
		default:
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": "Internal Server Error",
			})
			return
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"class":   class.Name,
		"message": "Success update class!",
	})
	return
}
