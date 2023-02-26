package db

import (
	"errors"
)

var (
	ErrAlreadyExist = errors.New("already exist")
	ErrCantBeEmpty  = errors.New("cannot be empty")
	ErrScoreInvalid = errors.New("score invalid")
	ErrNotFound     = errors.New("not found")
	ErrUnauthorized = errors.New("unauthorized")
	ScoreUpdated    = errors.New("updated")
	ScoreNotUpdated = errors.New("keep highest score")
)

// CreateUser is a function to create a user with hashed and salted password
func CreateUser(user User) error {
	if user.Name == "" || user.Username == "" || user.Password == "" || user.ClassID == 0 || user.RoleID == 0 {
		return ErrCantBeEmpty
	}

	// check if a username record exists in the table
	if err := DB.Where("username = ?", user.Username).First(&user).Error; err != nil {
		hashedPassword, err := hashAndSaltPassword(user.Password)
		if err != nil {
			return err
		}
		user.Password = hashedPassword

		res := DB.Create(&user)
		if res.Error != nil {
			return res.Error
		}
	} else {
		return ErrAlreadyExist
	}
	return nil
}

// GetUsersByClassId is a function to get all users based on classId
// hide password, role_id, class_id
func GetUsersByClassId(classId int) (users []Student, err error) {
	res := DB.Table("users").Where("class_id = ?", classId).Select("id, username, name").Find(&users)
	if res.Error != nil {
		return nil, res.Error
	}

	return users, nil
}

// UpdateUsersByUsersId is a function to update user based on userId
func UpdateUsersByUsersId(userId int, user User) error {
	if user.Name == "" || user.Username == "" || user.ClassID == 0 || user.RoleID == 0 {
		return ErrCantBeEmpty
	}

	var oldUser User

	if err := DB.Where("id = ?", userId).First(&oldUser).Error; err != nil {
		return ErrNotFound
	}

	res := DB.Table("users").Where("id = ?", userId).Updates(user)
	if res.Error != nil {
		return res.Error
	}

	return nil
}

// PatchUsersByUsersId is a function to patch user based on userId
func PatchUsersByUsersId(userId int, user User) error {
	var oldUser User

	if err := DB.Where("id = ?", userId).First(&oldUser).Error; err != nil {
		return ErrNotFound
	}

	res := DB.Table("users").Where("id = ?", userId).Updates(user)
	if res.Error != nil {
		return res.Error
	}

	return nil
}

// DeleteUsersByUsersId is a function to delete user based on userId
func DeleteUsersByUsersId(userId int) error {
	var user User
	if err := DB.Where("id = ?", userId).First(&user).Error; err != nil {
		return ErrNotFound
	}

	res := DB.Delete(&user)
	if res.Error != nil {
		return res.Error
	}

	return nil
}

// ========================================================

// CreateClass is a function to create a class
func CreateClass(class Class) (result Class, err error) {
	if class.Name == "" {
		return result, ErrCantBeEmpty
	}

	// check if a class name record exists in the table
	if err := DB.Where("name = ?", class.Name).First(&class).Error; err != nil {
		res := DB.Create(&class)
		if res.Error != nil {
			return result, res.Error
		}
		return class, nil
	} else {
		return result, ErrAlreadyExist
	}
}

// GetClasses is a function to get all classes
func GetClasses() ([]Class, error) {
	var classes []Class
	res := DB.Find(&classes)

	if res.Error != nil {
		return nil, res.Error
	}

	return classes, nil
}

// GetClassNameById is a function to get class name based on classId
func GetClassNameById(classId int) (string, error) {
	var class Class
	res := DB.Where("id = ?", classId).First(&class)
	if res.Error != nil {
		return "", res.Error
	}

	return class.Name, nil
}

// UpdateClassByClassId is a function to update class based on classId
func UpdateClassByClassId(classId int, class Class) error {
	if class.Name == "" {
		return ErrCantBeEmpty
	}

	var oldClass Class

	if err := DB.Where("id = ?", classId).First(&oldClass).Error; err != nil {
		return ErrNotFound
	}

	res := DB.Table("classes").Where("id = ?", classId).Updates(class)
	if res.Error != nil {
		return res.Error
	}

	return nil
}

// DeleteClassByClassId is a function to delete class based on classId
func DeleteClassByClassId(classId int) error {
	var class Class
	if err := DB.Where("id = ?", classId).First(&class).Error; err != nil {
		return ErrNotFound
	}

	res := DB.Delete(&class)
	if res.Error != nil {
		return res.Error
	}

	return nil
}

// ========================================================

// CreateCourse is a function to create a course
func CreateCourse(course Course) (result Course, err error) {
	if course.Name == "" {
		return result, ErrCantBeEmpty
	}

	// check if a course name record exists in the table
	if err := DB.Where("name = ?", course.Name).First(&course).Error; err != nil {
		res := DB.Create(&course)
		if res.Error != nil {
			return result, res.Error
		}
		return course, nil
	} else {
		return result, ErrAlreadyExist
	}
}

// GetCourses is a function to get all courses
func GetCourses() ([]Course, error) {
	var courses []Course
	res := DB.Find(&courses)
	if res.Error != nil {
		return nil, res.Error
	}

	return courses, nil
}

func GetCourseNameById(courseId int) (courseName string, err error) {
	var course Course
	res := DB.Where("id = ?", courseId).First(&course)
	if res.Error != nil {
		return "", res.Error
	}

	return course.Name, nil
}

// UpdateCourseByCourseId is a function to update course based on courseId
func UpdateCourseByCourseId(courseId int, course Course) error {
	if course.Name == "" {
		return ErrCantBeEmpty
	}

	var oldCourse Course

	if err := DB.Where("id = ?", courseId).First(&oldCourse).Error; err != nil {
		return ErrNotFound
	}

	res := DB.Model(&course).Where("id = ?", courseId).Updates(course)
	if res.Error != nil {
		return res.Error
	}

	return nil
}

// DeleteCourseByCourseId is a function to delete course based on courseId
func DeleteCourseByCourseId(courseId int) error {
	var course Course
	if err := DB.Where("id = ?", courseId).First(&course).Error; err != nil {
		return ErrNotFound
	}

	res := DB.Delete(&course)
	if res.Error != nil {
		return res.Error
	}

	return nil
}

// ========================================================

// CreateLab is a function to create a lab
func CreateLab(lab Lab) (result Lab, err error) {
	if lab.Name == "" {
		return result, ErrCantBeEmpty
	}

	// check if a lab name record exists in the table
	if err := DB.Where("name = ?", lab.Name).First(&lab).Error; err != nil {
		res := DB.Create(&lab)
		if res.Error != nil {
			return result, res.Error
		}
		return lab, nil
	} else {
		return result, ErrAlreadyExist
	}
}

// GetLabs is a function to get all labs based on courseId
func GetLabs(courseId string) (labs []GeneralData, err error) {
	if courseId == "" {
		return labs, ErrCantBeEmpty
	}
	res := DB.Table("labs").Where("course_id = ?", courseId).Select("id, name").Find(&labs)
	if res.Error != nil {
		return nil, res.Error
	}

	return labs, nil
}

// UpdateLabByLabId is a function to update lab based on labId
func UpdateLabByLabId(labId int, lab Lab) error {
	if lab.Name == "" {
		return ErrCantBeEmpty
	}

	var oldLab Lab

	if err := DB.Where("id = ?", labId).First(&oldLab).Error; err != nil {
		return ErrNotFound
	}

	res := DB.Model(&lab).Where("id = ?", labId).Updates(lab)
	if res.Error != nil {
		return res.Error
	}

	return nil
}

// DeleteLabByLabId is a function to delete lab based on labId
func DeleteLabByLabId(labId int) error {
	var lab Lab
	if err := DB.Where("id = ?", labId).First(&lab).Error; err != nil {
		return ErrNotFound
	}

	res := DB.Delete(&lab)
	if res.Error != nil {
		return res.Error
	}

	return nil
}

// ========================================================

// PushScore is a function to push score to database
func PushScore(userId int, score ScorePush) error {
	if score.Username == "" || score.Lab == "" {
		return ErrCantBeEmpty
	}

	// check if score is higher than 100
	if score.Score > 100 || score.Score < 0 {
		return ErrScoreInvalid
	}

	// lookup lab id by lab name
	var lab Lab
	res := DB.Where("name = ?", score.Lab).First(&lab)
	if res.Error != nil {
		return res.Error
	}

	// lookup user id by username
	var user User
	res = DB.Where("username = ?", score.Username).First(&user)
	if res.Error != nil {
		return res.Error
	}

	// check if user_id and userId are the same
	if user.ID != userId {
		return ErrUnauthorized
	}

	// check if a score record exists in the table
	var scores Score
	if err := DB.Where("user_id = ? AND lab_id = ?", user.ID, lab.ID).First(&scores).Error; err != nil {
		// create new score record
		scores = Score{
			UserID: user.ID,
			LabID:  lab.ID,
			Score:  score.Score,
		}
		res := DB.Create(&scores)
		if res.Error != nil {
			return res.Error
		}
		return ScoreUpdated

	} else {
		// update score record if score is higher
		if scores.Score < score.Score {
			res := DB.Model(&scores).Update("score", score.Score)
			if res.Error != nil {
				return res.Error
			}
			return ScoreUpdated
		}
	}

	return ScoreNotUpdated
}

// GetScoreByLabName get score by lab name
func GetScoreByLabName(userId int, labName string) (labScore ScoreLab, err error) {
	if labName == "" {
		return labScore, ErrCantBeEmpty
	}

	// lookup lab id by lab name
	var lab Lab
	res := DB.Where("name = ?", labName).First(&lab)
	if res.Error != nil {
		return labScore, res.Error
	}

	// lookup score by user id and lab id
	var scores Score
	res = DB.Where("user_id = ? AND lab_id = ?", userId, lab.ID).First(&scores)
	if res.Error != nil {
		return labScore, ErrNotFound
	}

	return ScoreLab{
		LabName: labName,
		Score:   scores.Score,
	}, nil
}

// ExportScores get score by course id and class id
func ExportScores(userId int, courseId int, classId int) (scores []ScoreLab, err error) {
	// join the score table with the labs table using the course_id foreign key
	// then join the labs table with the class table using the class_id foreign key
	query := DB.Table("scores").Select("labs.name, scores.score")
	query = query.Joins("JOIN labs ON labs.id = scores.lab_id").Joins("JOIN users ON users.id = scores.user_id")
	query = query.Where("users.id = ? AND labs.course_id = ? AND users.class_id = ?", userId, courseId, classId)

	// find the results
	if err := query.Find(&scores).Error; err != nil {
		return nil, err
	}

	return scores, nil
}
