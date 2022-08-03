package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/go-crud/pkg/forms"
	"github.com/go-crud/pkg/repositories"
)

type UserHandler struct{}

func (u UserHandler) Index(c *gin.Context) {
	users, err := repositories.UserRepository{}.Index()

	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})

		return
	}

	c.JSON(200, gin.H{"users": users})
}

func (u UserHandler) Show(c *gin.Context) {
	user, err := repositories.UserRepository{}.Show(c.Param("user"))

	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})

		return
	}

	c.JSON(200, gin.H{"user": user})
}

func (u UserHandler) Store(c *gin.Context) {
	var form forms.StoreUserForm

	if c.ShouldBind(&form) != nil {
		c.JSON(400, gin.H{"error": "Invalid form"})
		return
	}

	user, err := repositories.UserRepository{}.Store(form)

	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"user": user})
}

func (u UserHandler) Update(c *gin.Context) {
	var form forms.UpdateUserForm

	if c.ShouldBind(&form) != nil {
		c.JSON(400, gin.H{"error": "Invalid form"})
		return
	}

	user, err := repositories.UserRepository{}.Update(c.Param("user"), form)

	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"user": user})
}

func (u UserHandler) Destroy(c *gin.Context) {
	err := repositories.UserRepository{}.Destroy(c.Param("user"))

	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"message": "User deleted"})
}
