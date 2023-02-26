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

	res, err := db.CreateLab(labs)
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
		"lab":       res.Name,
		"id":        res.ID,
		"course_id": res.CourseID,
		"message":   "Success create labs!",
	})
	return
}

// TODO: Attach course id to response

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

// UpdateLabs is a function to update labs name based on lab_id from query
func UpdateLabs(c *gin.Context) {
	var labs db.Lab

	if err := c.BindJSON(&labs); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Bad Request",
		})
		return
	}

	labsId := c.Query("id")

	labsIdInt, err := strconv.Atoi(labsId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "lab_id must be integer",
		})
		return
	}

	err = db.UpdateLabByLabId(labsIdInt, labs)
	if err != nil {
		switch err {
		case db.ErrCantBeEmpty:
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "Lab name cant be empty",
			})
			return
		case db.ErrNotFound:
			c.JSON(http.StatusNotFound, gin.H{
				"message": "Labs Not Found",
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
		"lab":     labs.Name,
		"message": "Success update labs!",
	})
	return
}

// DeleteLabs is a function to delete labs based on lab_id from query
func DeleteLabs(c *gin.Context) {
	labsId := c.Query("id")

	labsIdInt, err := strconv.Atoi(labsId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "lab_id must be integer",
		})
		return
	}

	err = db.DeleteLabByLabId(labsIdInt)
	if err != nil {
		if errors.Is(err, db.ErrNotFound) {
			c.JSON(http.StatusNotFound, gin.H{
				"message": "Labs Not Found",
			})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Internal Server Error",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Success delete labs!",
	})
	return
}
