package repository

import (
	"technical-test-II/database"
	"technical-test-II/domain"
	"time"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

// get all users
func GetUsers() ([]domain.User, error) {
	db := database.GetDB()
	var users []domain.User
	err := db.Find(&users).Error
	return users, err
}

// register user
func RegisterUser(user *domain.User) error {
	db := database.GetDB()
	user.UserID = uuid.New().String()

	// hash pin
	hashedPin, err := bcrypt.GenerateFromPassword([]byte(user.Pin), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	user.Pin = string(hashedPin)

	// simpan user
	err = db.Create(user).Error
	return err
}

// cari user by phone_number
func FindUserByPhoneNumber(phoneNumber string) (*domain.User, error) {
	db := database.GetDB()
	var user domain.User
	err := db.Where("phone_number = ?", phoneNumber).First(&user).Error
	return &user, err
}

// cari user by user_id
func FindUserByID(userID string) (*domain.User, error) {
	db := database.GetDB()
	var user domain.User
	err := db.Where("user_id = ?", userID).First(&user).Error
	return &user, err
}

// update user's profile
func UpdateUserProfile(user *domain.User) error {
	db := database.GetDB()
	user.UpdatedDate = time.Now()
	return db.Save(user).Error
}

// menyimpan token JWT ke database
func StoreToken(token *domain.Token) error {
	db := database.GetDB()
	token.TokenID = uuid.New().String()
	return db.Create(token).Error
}
