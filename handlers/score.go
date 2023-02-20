package handlers

import (
	"github.com/Kyuubang/shopiea/db"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// TODO: create function to do handle score push by user_id, labs_id

// PushScore endpoint to push score
func PushScore(c *gin.Context) {
	// Bind the JSON payload to a User struct
	var score db.ScorePush
	if err := c.BindJSON(&score); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Bad Request",
		})
		return
	}

	userId := c.MustGet("userId").(string)

	// convert userId to int
	userIdInt, err := strconv.Atoi(userId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Internal Server Error",
		})
		return
	}

	// push score
	err = db.PushScore(userIdInt, score)
	if err != nil {
		switch err {
		case db.ErrCantBeEmpty:
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "user_id, labs_id, score cant be empty",
			})
			return
		case db.ScoreUpdated:
			c.JSON(http.StatusCreated, gin.H{
				"message": "score updated",
			})
			return
		case db.ScoreNotUpdated:
			c.JSON(http.StatusAccepted, gin.H{
				"message": "keep high score",
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
		"message": "successful push score",
	})
	return

}

// GetScore endpoint to get score
func GetScore(c *gin.Context) {
	userId := c.MustGet("userId").(string)

	// convert userId to int
	userIdInt, err := strconv.Atoi(userId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Internal Server Error",
		})
		return
	}

	//var score db.Score
	score, err := db.GetScoreByLabName(userIdInt, c.Query("lab_name"))
	if err != nil {
		switch err {
		case db.ErrCantBeEmpty:
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "lab_name cant be empty",
			})
			return
		case db.ErrNotFound:
			c.JSON(http.StatusNotFound, gin.H{
				"message": "score not found",
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
		"lab_name": score.LabName,
		"score":    score.Score,
	})
	return
}
