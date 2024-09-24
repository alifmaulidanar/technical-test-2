package routes

import (
	"net/http"
	"technical-test-II/domain"
	"technical-test-II/middleware"
	"technical-test-II/repository"
	"technical-test-II/service"

	"github.com/gin-gonic/gin"
)

func UserRoutes(r *gin.Engine) {
	r.GET("/users", getUsers)
	r.POST("/register", registerUser)
	r.POST("/login", login)
	r.PUT("/profile", middleware.AuthMiddleware(), updateProfile)
}

// menampilkan seluruh user yang terdaftar
func getUsers(c *gin.Context) {
	users, err := repository.GetUsers()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": "SUCCESS",
		"result": users,
	})
}

// register user
func registerUser(c *gin.Context) {
	var user domain.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if user.FirstName == "" || user.LastName == "" || user.PhoneNumber == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "First name, last name, and phone number are required"})
		return
	}

	existingUser, err := repository.FindUserByPhoneNumber(user.PhoneNumber)
	if err == nil && existingUser != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Phone Number already registered"})
		return
	}

	if err := repository.RegisterUser(&user); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// setelah semua validasi berhasil, kembalikan data
	c.JSON(http.StatusOK, gin.H{
		"status": "SUCCESS",
		"result": user,
	})
}

// login user
func login(c *gin.Context) {
	var loginRequest struct {
		PhoneNumber string `json:"phone_number"`
		Pin         string `json:"pin"`
	}
	if err := c.ShouldBindJSON(&loginRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := repository.FindUserByPhoneNumber(loginRequest.PhoneNumber)
	if err != nil || user == nil {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "Phone Number and PIN doesn't match."})
		return
	}

	if !service.CheckPinHash(loginRequest.Pin, user.Pin) {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "Phone Number and PIN doesn't match."})
		return
	}

	accessToken, err := service.GenerateJWT(*user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	refreshToken, err := service.GenerateJWT(*user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// membuat dan menyimpan token JWT user yg login
	token := domain.Token{
		UserID:       user.UserID,
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}

	if err := repository.StoreToken(&token); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// kembalikan response login berhasil
	c.JSON(http.StatusOK, gin.H{
		"status": "SUCCESS",
		"result": gin.H{
			"access_token":  accessToken,
			"refresh_token": refreshToken,
		},
	})
}

// memperbarui profile user
func updateProfile(c *gin.Context) {
	userID, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User ID not found in token"})
		return
	}

	var input domain.User
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := repository.FindUserByID(userID.(string))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	// perbarui field data user
	user.FirstName = input.FirstName
	user.LastName = input.LastName
	user.Address = input.Address

	// simpan pembaruan
	if err := repository.UpdateUserProfile(user); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update user"})
		return
	}

	// kembalikan response update profil berhasil
	c.JSON(http.StatusOK, gin.H{
		"status": "SUCCESS",
		"result": gin.H{
			"user_id":      user.UserID,
			"first_name":   user.FirstName,
			"last_name":    user.LastName,
			"address":      user.Address,
			"updated_date": user.UpdatedDate,
		},
	})
}
