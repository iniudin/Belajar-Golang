package main

import (
	"bookstore/controllers"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	user := controllers.NewUserController()

	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "Server is running!"})
	})

	apiRoutes := r.Group("/api")
	// userRoutes := apiRoutes.Group("/user")
	userProtectedRoutes := apiRoutes.Group("/users")
	{
		userProtectedRoutes.GET("/", user.GetAllUser)
	}

	r.Run()
}
