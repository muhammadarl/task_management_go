package controllers

import (
	"net/http"
	"task_management_backend/models"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type UserControllers struct {
	DB *gorm.DB
}

func (u *UserControllers) Login(ctx *gin.Context) {
	// Bind the JSON input to a struct that captures login credentials
	var loginData struct {
		Email    string `json:"uemail_address" binding:"required,email"`
		Password string `json:"password" binding:"required"`
	}

	// Bind the JSON input to loginData struct
	if err := ctx.ShouldBindJSON(&loginData); err != nil {
		var errMsg string
		for _, fieldErr := range err.(validator.ValidationErrors) {
			switch fieldErr.Field() {
			case "Email":
				if fieldErr.Tag() == "required" {
					errMsg = "Email is required."
				} else if fieldErr.Tag() == "email" {
					errMsg = "Email must be a valid email address."
				}
			case "Password":
				if fieldErr.Tag() == "required" {
					errMsg = "Password is required."
				}
			default:
				errMsg = "Invalid input data."
			}
			ctx.JSON(http.StatusBadRequest, gin.H{"error": errMsg})
		}
		return
	}

	// Find the user by email
	var userEmail models.UserEmail
	var user models.User

	// Check if the email exists in the database
	if err := u.DB.Where("uemail_address = ?", loginData.Email).First(&userEmail).Error; err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Email not found"})
		return
	}

	// Fetch the user details using the email ID
	if err := u.DB.Where("id = ?", userEmail.UserId).First(&user).Error; err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "User not found"})
		return
	}

	// Compare the hashed password with the password provided
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(loginData.Password)); err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid email or password"})
		return
	}

	// Successful login
	ctx.JSON(http.StatusOK, user)
}
func (u *UserControllers) Register(ctx *gin.Context) {
	var userData struct {
		FirstName     string `json:"first_name" binding:"required,max=255"`
		LastName      string `json:"last_name" binding:"required,max=255"`
		UemailAddress string `json:"uemail_address" binding:"required,email,max=255"`
		Password      string `json:"password" binding:"required"`
		UemailType    string `json:"uemail_type" binding:"required"`
		UemailPrimary bool   `json:"uemail_primary" binding:"required"`
	}

	// Bind the JSON input to userData struct
	if err := ctx.ShouldBindJSON(&userData); err != nil {
		var errMsg string
		for _, fieldErr := range err.(validator.ValidationErrors) {
			switch fieldErr.Field() {
			case "FirstName":
				if fieldErr.Tag() == "required" {
					errMsg = "First Name is required."
				} else if fieldErr.Tag() == "max" {
					errMsg = "First Name must be less than 255 characters."
				}
			case "LastName":
				if fieldErr.Tag() == "max" {
					errMsg = "Last Name must be less than 255 characters."
				}
			case "UemailAddress":
				if fieldErr.Tag() == "required" {
					errMsg = "Email is required."
				} else if fieldErr.Tag() == "email" {
					errMsg = "Email must be a valid email address."
				} else if fieldErr.Tag() == "max" {
					errMsg = "Email must be less than 255 characters."
				}
			case "Password":
				if fieldErr.Tag() == "required" {
					errMsg = "Password is required."
				}
			default:
				errMsg = "Invalid input data."
			}
			ctx.JSON(http.StatusBadRequest, gin.H{"error": errMsg})
		}
		return
	}
	var userEmail models.UserEmail
	if u.DB.Where("uemail_address = ?", userData.UemailAddress).First(&userEmail).RowsAffected != 0 {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Email already exists"})
		return
	}

	_, err := bcrypt.GenerateFromPassword([]byte(userData.Password), bcrypt.DefaultCost)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Error hashing the password"})
		return
	}

	// user := models.User{
	// 	FirstName: userData.FirstName,
	// 	LastName:  userData.LastName,
	// 	Password:  string(hashedPassword),
	// }
	// userEmail = models.UserEmail{
	// 	UemailAddress: userData.UemailAddress,
	// 	UemailType:    userData.UemailType,
	// 	IsPrimary:     userData.UemailPrimary,
	// }
}
