package config

import (
	"fmt"
	"task_management_backend/models"
	"time"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

const (
	DBDRIVER = "mysql"
	host     = "45.143.81.248"
	user     = "n1572675_arl"
	password = "Junijuli1!"
	port     = "3306"
	dbname   = "n1572675_task_mate"
)

func DatabaseConnection() *gorm.DB {
	dsn := user + ":" + password + "@tcp(" + host + ":" + port + ")/" + dbname + "?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	return db
}

func generateSecureUserHistoryId() (string, error) {
	uuid, err := uuid.NewRandom()
	if err != nil {
		return "", err
	}
	id := fmt.Sprintf("UH%s", uuid.String())
	return id, nil
}

func CreateOwnerAccount(db *gorm.DB) {
	// Check if email already exists before creating user and related records
	var existingEmail models.UserEmail
	emailCheck := db.Where("uemail_address = ?", "syiarul45@gmail.com").First(&existingEmail)
	if emailCheck.RowsAffected > 0 {
		fmt.Println("User already exists")
		return
	}

	// Generate hashed password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte("taskmate123"), bcrypt.DefaultCost)
	if err != nil {
		panic(err)
	}

	// Generate secure user ID
	ownerId, err := models.User.GenerateSecureUserId()
	if err != nil {
		panic(err)
	}

	// Create the user
	owner := models.User{
		Id:        ownerId,
		Role:      "Super Admin",
		FirstName: "Task",
		LastName:  "Mate",
		Password:  string(hashedPassword),
	}

	// Generate secure user history ID
	ownerHistoryId, err := models.GenerateSecureUserHistoryId()
	if err != nil {
		panic(err)
	}

	// Create user history
	ownerHistory := models.UserHistory{
		Id:             ownerHistoryId,
		UserId:         owner.Id,
		UhistoryAction: "Create",
		UhistoryTime:   time.Now(),
		UhistoryRemark: "Owner account created",
	}

	// Generate secure user email ID
	ownerEmailId, err := models.GenerateSecureUserEmailId()
	if err != nil {
		panic(err)
	}

	// Create user email
	ownerEmail := models.UserEmail{
		Id:            ownerEmailId,
		UserId:        owner.Id,
		UemailAddress: "syiarul45@gmail.com",
		UemailType:    "Work",
		IsPrimary:     true,
	}

	// Begin a transaction to ensure atomicity
	tx := db.Begin()

	if err := tx.Create(&owner).Error; err != nil {
		tx.Rollback()
		panic(err)
	}

	if err := tx.Create(&ownerHistory).Error; err != nil {
		tx.Rollback()
		panic(err)
	}

	if err := tx.Create(&ownerEmail).Error; err != nil {
		tx.Rollback()
		panic(err)
	}

	// Commit the transaction
	if err := tx.Commit().Error; err != nil {
		panic(err)
	}

	fmt.Println("Owner account created successfully")
}
