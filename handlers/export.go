package handlers

import (
	"github.com/Kyuubang/shopiea/db"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"time"
)

// ExportScore create function for export with userId from Auth middleware, course_id, labs_id from query
func ExportScore(c *gin.Context) {
	classId := c.Query("class_id")
	courseId := c.Query("course_id")

	courseIdInt, err := strconv.Atoi(courseId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Internal Server Error",
		})
		return
	}

	classIdInt, err := strconv.Atoi(classId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Internal Server Error",
		})
		return
	}
	courseName, err := db.GetCourseNameById(courseIdInt)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "Course Not Found",
		})
		return
	}

	className, err := db.GetClassNameById(classIdInt)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "Class Not Found",
		})
		return
	}

	report, err := structureReport(courseIdInt, classIdInt)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Internal Server Error",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"class":   className,
		"course":  courseName,
		"date":    time.Now().Format("2006-01-02"),
		"time":    time.Now().Format("15:04:05"),
		"reports": report,
	})
	return
}

func calculateAverageScore(scores []db.ScoreLabs) (avarage float64, total int) {
	var sum int
	if len(scores) == 0 {
		return 0, 0
	}
	for _, s := range scores {
		sum += s.Score
	}
	return float64(sum) / float64(len(scores)), sum
}

func structureReport(courseId int, classId int) (report []db.Report, err error) {
	students, err := db.GetUsersByClassId(classId)
	if err != nil {
		return nil, err
	}

	// convert int to string
	courseIdStr := strconv.Itoa(courseId)

	for _, student := range students {

		labs, err := db.GetLabs(courseIdStr)
		if err != nil {
			return nil, err
		}

		scoreLabs, err := db.ExportScores(student.ID, courseId, classId)
		if err != nil {
			return nil, err
		}
		var scoreLabsStruct []db.ScoreLabs
		for _, lab := range labs {
			if len(scoreLabs) == 0 {
				scoreLabsStruct = append(scoreLabsStruct, db.ScoreLabs{
					LabName: lab.Name,
					Score:   0,
					ID:      lab.ID,
				})
				continue
			}
			var isExist bool
			for _, scoreLab := range scoreLabs {
				if scoreLab.LabName == lab.Name {
					scoreLabsStruct = append(scoreLabsStruct, db.ScoreLabs{
						LabName: lab.Name,
						Score:   scoreLab.Score,
						ID:      lab.ID,
					})
					isExist = true
				}
			}
			if !isExist {
				scoreLabsStruct = append(scoreLabsStruct, db.ScoreLabs{
					LabName: lab.Name,
					Score:   0,
					ID:      lab.ID,
				})
			}
		}

		averageScore, totalScore := calculateAverageScore(scoreLabsStruct)
		report = append(report, db.Report{
			Name:     student.Name,
			Username: student.Username,
			Average:  averageScore,
			Scores:   scoreLabsStruct,
			Total:    totalScore,
		})
	}
	return report, nil
}
