package domain

import (
	"time"
)

// struct token sesuai database MySQL
type Token struct {
	TokenID      string    `json:"token_id" gorm:"type:char(36);not null;primaryKey"`
	UserID       string    `json:"user_id" gorm:"type:char(36);not null"`
	AccessToken  string    `json:"access_token" gorm:"type:text;not null"`
	RefreshToken string    `json:"refresh_token" gorm:"type:text;not null"`
	CreatedDate  time.Time `json:"created_date" gorm:"type:timestamp;default:CURRENT_TIMESTAMP"`
}
