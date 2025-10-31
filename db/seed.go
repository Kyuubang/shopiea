package db

import (
	"golang.org/x/crypto/bcrypt"
	"log"
)

// Seed populates the database with initial data
func Seed() error {
	// Seed Roles
	if err := seedRoles(); err != nil {
		return err
	}

	// Seed Classes
	if err := seedClasses(); err != nil {
		return err
	}

	// Seed Admin Users
	if err := seedAdminUsers(); err != nil {
		return err
	}

	// Seed Courses
	if err := seedCourses(); err != nil {
		return err
	}

	// Seed Labs
	if err := seedLabs(); err != nil {
		return err
	}

	log.Println("Database seeding completed successfully")
	return nil
}

// seedRoles creates the default roles (admin and student)
func seedRoles() error {
	roles := []Role{
		{ID: 1, Name: "admin"},
		{ID: 2, Name: "student"},
	}

	for _, role := range roles {
		var existingRole Role
		result := DB.Where("name = ?", role.Name).First(&existingRole)
		if result.Error != nil {
			// Role doesn't exist, create it
			if err := DB.Create(&role).Error; err != nil {
				log.Printf("Error creating role %s: %v", role.Name, err)
				return err
			}
			log.Printf("Created role: %s", role.Name)
		} else {
			log.Printf("Role already exists: %s", role.Name)
		}
	}

	return nil
}

// seedClasses creates the default class
func seedClasses() error {
	classes := []Class{
		{ID: 1, Name: "Shopiea"},
	}

	for _, class := range classes {
		var existingClass Class
		result := DB.Where("name = ?", class.Name).First(&existingClass)
		if result.Error != nil {
			// Class doesn't exist, create it
			if err := DB.Create(&class).Error; err != nil {
				log.Printf("Error creating class %s: %v", class.Name, err)
				return err
			}
			log.Printf("Created class: %s", class.Name)
		} else {
			log.Printf("Class already exists: %s", class.Name)
		}
	}

	return nil
}

// seedAdminUsers creates the default admin user
func seedAdminUsers() error {
	// Hash the default password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte("admin123"), bcrypt.DefaultCost)
	if err != nil {
		log.Printf("Error hashing password: %v", err)
		return err
	}

	users := []User{
		{
			ID:       1,
			Username: "admin",
			Password: string(hashedPassword),
			Name:     "Administrator",
			RoleID:   1, // admin role
			ClassID:  1, // Shopiea class
		},
	}

	for _, user := range users {
		var existingUser User
		result := DB.Where("username = ?", user.Username).First(&existingUser)
		if result.Error != nil {
			// User doesn't exist, create it
			if err := DB.Create(&user).Error; err != nil {
				log.Printf("Error creating user %s: %v", user.Username, err)
				return err
			}
			log.Printf("Created admin user: %s (default password set - change immediately!)", user.Username)
		} else {
			log.Printf("User already exists: %s", user.Username)
		}
	}

	return nil
}

// seedCourses creates the default course
func seedCourses() error {
	courses := []Course{
		{ID: 1, Name: "golang"},
	}

	for _, course := range courses {
		var existingCourse Course
		result := DB.Where("name = ?", course.Name).First(&existingCourse)
		if result.Error != nil {
			// Course doesn't exist, create it
			if err := DB.Create(&course).Error; err != nil {
				log.Printf("Error creating course %s: %v", course.Name, err)
				return err
			}
			log.Printf("Created course: %s", course.Name)
		} else {
			log.Printf("Course already exists: %s", course.Name)
		}
	}

	return nil
}

// seedLabs creates the default lab
func seedLabs() error {
	labs := []Lab{
		{
			ID:       1,
			Name:     "golang-001",
			CourseID: 1, // golang course
		},
	}

	for _, lab := range labs {
		var existingLab Lab
		result := DB.Where("name = ?", lab.Name).First(&existingLab)
		if result.Error != nil {
			// Lab doesn't exist, create it
			if err := DB.Create(&lab).Error; err != nil {
				log.Printf("Error creating lab %s: %v", lab.Name, err)
				return err
			}
			log.Printf("Created lab: %s for course ID %d", lab.Name, lab.CourseID)
		} else {
			log.Printf("Lab already exists: %s", lab.Name)
		}
	}

	return nil
}
