package main

import (
	"technical-test-II/database"
	"technical-test-II/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	database.StartDB()
	r := gin.Default()

	// routes
	routes.UserRoutes(r)
	routes.TransactionRoutes(r)

	r.Run(":8080")
}
