package routes

import (
	"fmt"
	"net/http"
	"technical-test-II/middleware"
	"technical-test-II/repository"

	"github.com/gin-gonic/gin"
)

func TransactionRoutes(r *gin.Engine) {
	r.POST("/topup", middleware.AuthMiddleware(), topUp)
	r.GET("/transactions", middleware.AuthMiddleware(), getTransactionsReport)
	r.POST("/pay", middleware.AuthMiddleware(), makePayment)
	r.POST("/transfer", middleware.AuthMiddleware(), transferBalance)
}

func getTransactionsReport(c *gin.Context) {
	userID, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "Unauthenticated"})
		return
	}

	// menampilkan seluruh transaksi milik user tersebut
	transactions, err := repository.GetTransactionsByUserID(userID.(string))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	var responseTransactions []gin.H
	for _, transaction := range transactions {
		var referenceIDField string
		switch transaction.TransactionReference {
		case "top_up_id":
			referenceIDField = "top_up_id"
		case "payment_id":
			referenceIDField = "payment_id"
		case "transfer_id":
			referenceIDField = "transfer_id"
		default:
			referenceIDField = "transaction_id"
		}

		transactionData := gin.H{
			referenceIDField:   transaction.TransactionID,
			"transaction_type": transaction.TransactionType,
			"amount":           transaction.Amount,
			"remarks":          transaction.Remarks,
			"balance_before":   transaction.BalanceBefore,
			"balance_after":    transaction.BalanceAfter,
			"created_date":     transaction.CreatedDate,
		}

		responseTransactions = append(responseTransactions, transactionData)
	}

	// kembalikan response sukses
	c.JSON(http.StatusOK, gin.H{
		"status": "SUCCESS",
		"result": responseTransactions,
	})
}

func topUp(c *gin.Context) {
	userID, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "Unauthenticated"})
		return
	}

	var input struct {
		Amount float64 `json:"amount" binding:"required"`
	}
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := repository.FindUserByID(userID.(string))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	// proses topup di repository
	transaction, err := repository.TopUp(user, input.Amount)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to process top-up"})
		return
	}

	// kembalikan response sukses
	c.JSON(http.StatusOK, gin.H{
		"status": "SUCCESS",
		"result": gin.H{
			"top_up_id":      transaction.TransactionID,
			"amount_top_up":  transaction.Amount,
			"balance_before": transaction.BalanceBefore,
			"balance_after":  transaction.BalanceAfter,
			"created_date":   transaction.CreatedDate,
		},
	})
}

// melakukan pembayaran
func makePayment(c *gin.Context) {
	userID, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "Unauthenticated"})
		return
	}

	var input struct {
		Amount  float64 `json:"amount" binding:"required"`
		Remarks string  `json:"remarks"`
	}
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := repository.FindUserByID(userID.(string))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	if user.Balance < input.Amount {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Balance is not enough"})
		return
	}

	// proses pembayaran di repository
	transaction, err := repository.MakePayment(user, input.Amount, input.Remarks)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to process payment"})
		return
	}

	// kembaikan response sukses
	c.JSON(http.StatusOK, gin.H{
		"status": "SUCCESS",
		"result": gin.H{
			"payment_id":     transaction.TransactionID,
			"amount":         transaction.Amount,
			"remarks":        transaction.Remarks,
			"balance_before": transaction.BalanceBefore,
			"balance_after":  transaction.BalanceAfter,
			"created_date":   transaction.CreatedDate,
		},
	})
}

func transferBalance(c *gin.Context) {
	senderID, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "Unauthenticated"})
		return
	}

	var input struct {
		TargetUserID string  `json:"target_user" binding:"required"`
		Amount       float64 `json:"amount" binding:"required"`
		Remarks      string  `json:"remarks"`
	}
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// langsung kembalikan response
	c.JSON(http.StatusAccepted, gin.H{
		"status": "Processing transfer in background",
	})

	// transfer diproses pada background menggunakan Goroutine
	go func() {
		sender, err := repository.FindUserByID(senderID.(string))
		if err != nil {
			fmt.Println("Sender not found: ", err)
			return
		}

		targetUser, err := repository.FindUserByID(input.TargetUserID)
		if err != nil {
			fmt.Println("Target user not found: ", err)
			return
		}

		// cek pengirim memiliki saldo yang cukup
		if sender.Balance < input.Amount {
			fmt.Println("Sender does not have enough balance")
			return
		}

		// proses transfer di repository
		_, _, err = repository.TransferBalance(sender, targetUser, input.Amount, input.Remarks)
		if err != nil {
			fmt.Println("Failed to process transfer: ", err)
			return
		}

		fmt.Println("Transfer successfully processed in background")
	}()
}
