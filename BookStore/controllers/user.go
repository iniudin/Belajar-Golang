package controllers

import (
	"bookstore/repository"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserController interface {
	GetAllUser(ctx *gin.Context)
}

type userController struct {
	repo repository.UserRepository
}

func NewUserController() UserController {
	return &userController{
		repo: repository.NewUserRepository(),
	}
}

// Comment unused Function
// func hashPassword(password *string) {
// 	byte_pass := []byte(*password)
// 	hash_pass, _ := bcrypt.GenerateFromPassword(byte_pass, bcrypt.DefaultCost)
// 	*password = string(hash_pass)
// }

// func comparePassword(dbpass, pass string) bool {
// 	return bcrypt.CompareHashAndPassword([]byte(dbpass), []byte(pass)) == nil
// }

func (c *userController) GetAllUser(ctx *gin.Context) {
	users, err := c.repo.GetAllUser()
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, users)
}
