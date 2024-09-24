package domain

import (
	"time"
)

// struct transaction sesuai database MySQL
type Transaction struct {
	TransactionID        string    `json:"transaction_id" gorm:"type:char(36);not null;primaryKey"`
	UserID               string    `json:"user_id" gorm:"type:char(36);not null"`
	TransactionType      string    `json:"transaction_type" gorm:"type:enum('CREDIT', 'DEBIT');not null"`
	Amount               float64   `json:"amount" gorm:"type:decimal(15,2);not null"`
	Remarks              string    `json:"remarks" gorm:"type:text"`
	BalanceBefore        float64   `json:"balance_before" gorm:"type:decimal(15,2);not null"`
	BalanceAfter         float64   `json:"balance_after" gorm:"type:decimal(15,2);not null"`
	TransactionReference string    `json:"transaction_reference" gorm:"type:char(36);not null"`
	CreatedDate          time.Time `json:"created_date" gorm:"type:timestamp;default:CURRENT_TIMESTAMP"`
}
