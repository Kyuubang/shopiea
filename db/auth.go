package db

import (
	"golang.org/x/crypto/bcrypt"
	"strconv"
)

const (
	salt = "shopiea"
)

// IsAdmin Check User if admin or not
func IsAdmin(userId string) bool {
	var user User

	// convert string to int
	userID, _ := strconv.Atoi(userId)

	DB.Preload("Role").First(&user, userID)
	return user.Role.Name == "admin"
}

// ValidationUserLogin for user login
func ValidationUserLogin(login Login) (string, bool) {
	var user User
	if err := DB.Where("username = ?", login.Username).First(&user).Error; err != nil {
		return "", false
	}
	userId := strconv.FormatUint(uint64(user.ID), 10)

	return userId, verifyPassword(login.Password, user.Password)
}

// TODO: change static salt for password to dynamic salt

// VerifyPassword takes a plain text password and a hashed and salted password, and returns whether they match
func verifyPassword(plainPassword, hashedPassword string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(plainPassword+salt))
	return err == nil
}

func hashAndSaltPassword(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password+salt), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hash), nil
}
