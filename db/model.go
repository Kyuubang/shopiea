package db

import "time"

// User represents a user of the system
type User struct {
	ID       int    `gorm:"primaryKey" json:"id"`
	Username string `gorm:"uniqueIndex;not null" json:"username"`
	Password string `gorm:"not null" json:"password"`
	Name     string `gorm:"not null" json:"name"`
	RoleID   int    `gorm:"not null" json:"role_id"`
	Role     Role   `gorm:"foreignKey:RoleID"`
	ClassID  int    `gorm:"not null" json:"class_id"`
	Class    Class  `gorm:"foreignKey:ClassID"`
}

// Role represents a role of the system
type Role struct {
	ID   int    `gorm:"primaryKey" json:"id"`
	Name string `gorm:"uniqueIndex;not null" json:"name"`
}

// Class represents a class of the student
type Class struct {
	ID   int    `gorm:"primaryKey" json:"id"`
	Name string `gorm:"uniqueIndex;not null" json:"name"`
}

// Course represents a course
type Course struct {
	ID   int    `gorm:"primaryKey" json:"id"`
	Name string `gorm:"uniqueIndex;not null" json:"name"`
}

// Lab represents a lab of a course
type Lab struct {
	ID       int    `gorm:"primaryKey" json:"id"`
	Name     string `gorm:"uniqueIndex;not null" json:"name"`
	CourseID int    `gorm:"not null" json:"course_id"`
	Course   Course `gorm:"foreignKey:CourseID"`
}

// Score represents a student score of a lab
type Score struct {
	ID        int       `gorm:"primaryKey" json:"id"`
	UserID    int       `gorm:"not null" json:"user_id"`
	User      User      `gorm:"foreignKey:UserID"`
	LabID     int       `gorm:"not null" json:"lab_id"`
	Lab       Lab       `gorm:"foreignKey:LabID"`
	Score     int       `gorm:"not null" json:"score"`
	CreatedAt time.Time `gorm:"not null" json:"created_at"`
	UpdatedAt time.Time `gorm:"not null" json:"updated_at"`
}

// Student response struct
type Student struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Name     string `json:"name"`
}

// GeneralData struct for data
type GeneralData struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

// Report struct
type Report struct {
	Name     string      `json:"name"`
	Username string      `json:"username"`
	Scores   []ScoreLabs `json:"scores,omitempty"`
	Average  float64     `json:"average"`
	Total    int         `json:"total"`
}

// ScoreLab single report score based on lab name
type ScoreLab struct {
	LabName string `json:"lab" gorm:"column:name"`
	Score   int    `json:"score"`
}

type ScoreLabs struct {
	LabName string `json:"lab_name"`
	Score   int    `json:"score"`
	ID      int    `json:"id"`
}

// Login Model
type Login struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// ScorePush struct
type ScorePush struct {
	Username string `json:"username"`
	Lab      string `json:"lab"`
	Score    int    `json:"score"`
}
