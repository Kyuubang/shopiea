package handlers

import (
	"errors"
	"github.com/Kyuubang/shopiea/db"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
	"strconv"
)

// CreateCourse is a function to create course
func CreateCourse(c *gin.Context) {
	var course db.Course

	if err := c.BindJSON(&course); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Bad Request",
		})
		return
	}

	res, err := db.CreateCourse(course)
	if err != nil {
		if errors.Is(err, db.ErrAlreadyExist) {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "Labs already exist",
			})
			return
		} else if errors.Is(err, db.ErrCantBeEmpty) {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "Course name cant be empty",
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
		"course":  res.Name,
		"id":      res.ID,
		"message": "Success create course!",
	})
	return
}

// GetCourses is a function to get all courses
func GetCourses(c *gin.Context) {
	courses, err := db.GetCourses()
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, gin.H{
				"message": "Courses Not Found",
			})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Internal Server Error",
		})
		return
	}

	if len(courses) == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "Course Not Found",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"courses": courses,
	})
	return
}

// DeleteCourse is a function to delete course by course_id
func DeleteCourse(c *gin.Context) {
	courseID := c.Query("id")

	// convert string to int
	courseIdInt, err := strconv.Atoi(courseID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Bad Request",
		})
		return
	}

	err = db.DeleteCourseByCourseId(courseIdInt)
	if err != nil {
		switch err {
		case db.ErrNotFound:
			c.JSON(http.StatusNotFound, gin.H{
				"message": "Course Not Found",
			})
			return
		case db.ErrCantBeEmpty:
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "Course name cant be empty",
			})
			return
		default:
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": "Internal Server Error",
			})
		}
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "Success delete course!",
	})
	return
}

// UpdateCourse is a function to update course by course_id
func UpdateCourse(c *gin.Context) {
	// get course_id from query
	courseID := c.Query("id")

	// convert string to int
	courseIdInt, err := strconv.Atoi(courseID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Bad Request",
		})
		return
	}

	var course db.Course

	if err := c.BindJSON(&course); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Bad Request",
		})
		return
	}

	err = db.UpdateCourseByCourseId(courseIdInt, course)
	if err != nil {
		switch err {
		case db.ErrNotFound:
			c.JSON(http.StatusNotFound, gin.H{
				"message": "Course Not Found",
			})
			return
		case db.ErrCantBeEmpty:
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "Course name cant be empty",
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
		"course":  course.Name,
		"message": "Success update course!",
	})
	return
}
