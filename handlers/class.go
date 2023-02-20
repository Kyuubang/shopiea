package handlers

import (
	"errors"
	"github.com/Kyuubang/shopiea/db"
	"github.com/gin-gonic/gin"
	"net/http"
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

	err := db.CreateClass(class)
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

// TODO: create update class name by class id
