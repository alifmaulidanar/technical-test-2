package routes

import (
	"net/http"
	"technical-test-II/middleware"
	"technical-test-II/repository"

	"github.com/gin-gonic/gin"
)

// SetupTransactionRoutes sets up the routes for transactions
func TransactionRoutes(r *gin.Engine) {
	r.POST("/topup", middleware.AuthMiddleware(), topUp)
	r.GET("/transactions", middleware.AuthMiddleware(), getTransactionsReport)
	r.POST("/pay", middleware.AuthMiddleware(), makePayment)
	r.POST("/transfer", middleware.AuthMiddleware(), transferBalance)
}

func getTransactionsReport(c *gin.Context) {
	// Get user_id from JWT token (set by middleware)
	userID, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "Unauthenticated"})
		return
	}

	// Retrieve all transactions for the user from the repository
	transactions, err := repository.GetTransactionsByUserID(userID.(string))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Prepare the response data by adjusting transaction_reference
	var responseTransactions []gin.H
	for _, transaction := range transactions {
		// Determine the correct ID based on transaction_reference
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

		// Prepare the transaction data with the appropriate ID field
		transactionData := gin.H{
			referenceIDField:   transaction.TransactionID,
			"transaction_type": transaction.TransactionType,
			"amount":           transaction.Amount,
			"remarks":          transaction.Remarks,
			"balance_before":   transaction.BalanceBefore,
			"balance_after":    transaction.BalanceAfter,
			"created_date":     transaction.CreatedDate,
		}

		// Add the transaction to the response list
		responseTransactions = append(responseTransactions, transactionData)
	}

	// Return the transactions in the response
	c.JSON(http.StatusOK, gin.H{
		"status": "SUCCESS",
		"result": responseTransactions,
	})
}

func topUp(c *gin.Context) {
	// Get user_id from JWT token (set by middleware)
	userID, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "Unauthenticated"})
		return
	}

	// Bind JSON input to a struct
	var input struct {
		Amount float64 `json:"amount" binding:"required"`
	}
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Retrieve the user by user_id
	user, err := repository.FindUserByID(userID.(string))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	// Process the top-up in the repository
	transaction, err := repository.TopUp(user, input.Amount)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to process top-up"})
		return
	}

	// Return success response
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

// makePayment handles the user's payment transaction
func makePayment(c *gin.Context) {
	// Get user_id from JWT token (set by middleware)
	userID, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "Unauthenticated"})
		return
	}

	// Bind JSON input to a struct
	var input struct {
		Amount  float64 `json:"amount" binding:"required"`
		Remarks string  `json:"remarks"`
	}
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Retrieve the user by user_id
	user, err := repository.FindUserByID(userID.(string))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	// Check if the user has enough balance
	if user.Balance < input.Amount {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Balance is not enough"})
		return
	}

	// Process the payment in the repository
	transaction, err := repository.MakePayment(user, input.Amount, input.Remarks)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to process payment"})
		return
	}

	// Return success response
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
	// Get user_id from JWT token (set by middleware)
	senderID, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "Unauthenticated"})
		return
	}

	// Bind JSON input to a struct
	var input struct {
		TargetUserID string  `json:"target_user" binding:"required"`
		Amount       float64 `json:"amount" binding:"required"`
		Remarks      string  `json:"remarks"`
	}
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Retrieve the sender (user making the transfer)
	sender, err := repository.FindUserByID(senderID.(string))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Sender not found"})
		return
	}

	// Retrieve the target user
	targetUser, err := repository.FindUserByID(input.TargetUserID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Target user not found"})
		return
	}

	// Check if the sender has enough balance
	if sender.Balance < input.Amount {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Balance is not enough"})
		return
	}

	// Process the balance transfer in the repository
	transferTransaction, _, err := repository.TransferBalance(sender, targetUser, input.Amount, input.Remarks)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to process transfer"})
		return
	}

	// Return success response
	c.JSON(http.StatusOK, gin.H{
		"status": "SUCCESS",
		"result": gin.H{
			"transfer_id":    transferTransaction.TransactionID,
			"amount":         input.Amount,
			"remarks":        input.Remarks,
			"balance_before": transferTransaction.BalanceBefore,
			"balance_after":  transferTransaction.BalanceAfter,
			"created_date":   transferTransaction.CreatedDate,
		},
	})
}
