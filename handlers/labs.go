package handlers

import (
	"errors"
	"github.com/Kyuubang/shopiea/db"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// CreateLabs is a function to create labs based on course name
func CreateLabs(c *gin.Context) {
	var labs db.Lab

	if err := c.BindJSON(&labs); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Bad Request",
		})
		return
	}

	err := db.CreateLab(labs)
	if err != nil {
		if errors.Is(err, db.ErrAlreadyExist) {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "Labs already exist",
			})
			return
		} else if errors.Is(err, db.ErrCantBeEmpty) {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "Lab name cant be empty",
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
		"labs":    labs.Name,
		"message": "Success create labs!",
	})
	return
}

// GetLabs is a function to get all labs based on course name
func GetLabs(c *gin.Context) {
	labs, err := db.GetLabs(c.Query("course_id"))
	if err != nil {
		if errors.Is(err, db.ErrCantBeEmpty) {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "course_id cant empety",
			})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Internal Server Error",
		})
		return
	}

	if len(labs) == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "Labs Not Found",
		})
		return
	}

	courseId := c.Query("course_id")

	courseIdInt, err := strconv.Atoi(courseId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "course_id must be integer",
		})
		return
	}

	courseName, err := db.GetCourseNameById(courseIdInt)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "Course Not Found",
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"course": courseName,
		"labs":   labs,
	})
	return
}

// TODO: create update handlers lab name by lab_id
