package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/zanuardinovanda/go-testing/models"
	"gorm.io/gorm"
)

type UserController struct {
	gorm.Model
	Username string
	Email    string
	Password string
}

//Get
func GetUser(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	var user []models.User
	db.Find(&user)
	c.JSON(http.StatusOK, gin.H{"data": user})
}

//Edit User
func EditUser(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	//check in database
	var user models.User
	if err := db.Where("id = ?", c.Param("id")).First(&user).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"data": "Data not found"})
	}

	//validation
	var input UserController
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	//edit data
	db.Model(&user).Where(user).Updates(input)
	c.JSON(http.StatusOK, gin.H{"data": user})
}

//DELETE DATA
func DeleteUser(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	//check in database
	var user models.User
	if err := db.Where("id = ?", c.Param("id")).First(&user).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"data": "Data not found"})
	}

	//delete data
	db.Delete(&user)
	c.JSON(http.StatusOK, gin.H{"data": "User deleted"})
}

//Register & Login User
func RegisterUser(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	//validation
	var input UserController
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	//input
	user := models.User{
		Username: input.Username,
		Email:    input.Email,
		Password: input.Password,
	}

	db.Create(&user)
	c.JSON(http.StatusOK, gin.H{"data:": user})
}
