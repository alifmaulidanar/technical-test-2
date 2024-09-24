package domain

import (
	"time"
)

// struct user sesuai database MySQL
type User struct {
	UserID      string    `json:"user_id" gorm:"type:char(36);not null;primaryKey"`
	FirstName   string    `json:"first_name" gorm:"type:varchar(50);not null"`
	LastName    string    `json:"last_name" gorm:"type:varchar(50);not null"`
	PhoneNumber string    `json:"phone_number" gorm:"type:varchar(15);not null;unique"`
	Address     string    `json:"address" gorm:"type:text;not null"`
	Pin         string    `json:"pin" gorm:"type:varchar(255);not null"`
	Balance     float64   `json:"balance" gorm:"type:decimal(15,2);default:0"`
	CreatedDate time.Time `json:"created_date" gorm:"type:timestamp;default:CURRENT_TIMESTAMP"`
	UpdatedDate time.Time `json:"updated_date" gorm:"type:timestamp;default:CURRENT_TIMESTAMP"`
}
