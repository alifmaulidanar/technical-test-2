package repository

import (
	"technical-test-II/database"
	"technical-test-II/domain"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// mengambil seluruh transaksi untuk suatu user
func GetTransactionsByUserID(userID string) ([]domain.Transaction, error) {
	db := database.GetDB()

	var transactions []domain.Transaction
	err := db.Where("user_id = ?", userID).Find(&transactions).Error

	return transactions, err
}

// proses top-up saldo user
func TopUp(user *domain.User, amount float64) (*domain.Transaction, error) {
	db := database.GetDB()
	balanceBefore := user.Balance
	balanceAfter := user.Balance + amount
	topUpID := uuid.New().String()

	// transaksi top-up
	topUpTransaction := domain.Transaction{
		TransactionID:        topUpID,
		UserID:               user.UserID,
		TransactionType:      "CREDIT",
		Amount:               amount,
		BalanceBefore:        balanceBefore,
		BalanceAfter:         balanceAfter,
		TransactionReference: "top_up_id",
		CreatedDate:          time.Now(),
	}

	user.Balance = balanceAfter
	user.UpdatedDate = time.Now()

	// simpan transaksi dan update user
	err := db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(&topUpTransaction).Error; err != nil {
			return err
		}
		if err := tx.Save(user).Error; err != nil {
			return err
		}
		return nil
	})

	return &topUpTransaction, err
}

// proses pembayaran oleh user
func MakePayment(user *domain.User, amount float64, remarks string) (*domain.Transaction, error) {
	db := database.GetDB()
	balanceBefore := user.Balance
	balanceAfter := user.Balance - amount
	paymentID := uuid.New().String()

	// transaksi pembayaran
	paymentTransaction := domain.Transaction{
		TransactionID:        paymentID,
		UserID:               user.UserID,
		TransactionType:      "DEBIT",
		Amount:               amount,
		BalanceBefore:        balanceBefore,
		BalanceAfter:         balanceAfter,
		Remarks:              remarks,
		TransactionReference: "payment_id",
		CreatedDate:          time.Now(),
	}

	user.Balance = balanceAfter
	user.UpdatedDate = time.Now()

	// simpan transaksi dan update user
	err := db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(&paymentTransaction).Error; err != nil {
			return err
		}
		if err := tx.Save(user).Error; err != nil {
			return err
		}
		return nil
	})

	return &paymentTransaction, err
}

// proses transfer saldo antar user
func TransferBalance(sender *domain.User, targetUser *domain.User, amount float64, remarks string) (*domain.Transaction, *domain.Transaction, error) {
	db := database.GetDB()
	senderBalanceBefore := sender.Balance
	senderBalanceAfter := sender.Balance - amount
	targetBalanceBefore := targetUser.Balance
	targetBalanceAfter := targetUser.Balance + amount
	transferID := uuid.New().String()

	// transaksi transfer debit, untuk user pengirim
	transferTransaction := domain.Transaction{
		TransactionID:        transferID,
		UserID:               sender.UserID,
		TransactionType:      "DEBIT",
		Amount:               amount,
		BalanceBefore:        senderBalanceBefore,
		BalanceAfter:         senderBalanceAfter,
		Remarks:              remarks,
		TransactionReference: "transfer_id",
		CreatedDate:          time.Now(),
	}

	// transaksi transfer kredit, untuk user penerima
	creditTransaction := domain.Transaction{
		TransactionID:        uuid.New().String(),
		UserID:               targetUser.UserID,
		TransactionType:      "CREDIT",
		Amount:               amount,
		BalanceBefore:        targetBalanceBefore,
		BalanceAfter:         targetBalanceAfter,
		Remarks:              remarks,
		TransactionReference: "transfer_id",
		CreatedDate:          time.Now(),
	}

	// simpan transaksi kedua dan update kedua user
	err := db.Transaction(func(tx *gorm.DB) error {
		sender.Balance = senderBalanceAfter
		sender.UpdatedDate = time.Now()
		if err := tx.Save(sender).Error; err != nil {
			return err
		}

		targetUser.Balance = targetBalanceAfter
		targetUser.UpdatedDate = time.Now()
		if err := tx.Save(targetUser).Error; err != nil {
			return err
		}

		if err := tx.Create(&transferTransaction).Error; err != nil {
			return err
		}

		if err := tx.Create(&creditTransaction).Error; err != nil {
			return err
		}

		return nil
	})

	return &transferTransaction, &creditTransaction, err
}
