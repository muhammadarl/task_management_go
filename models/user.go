package models

import (
	"fmt"

	"github.com/google/uuid"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type UserRole string

const (
	Admin      UserRole = "Admin"
	SuperAdmin UserRole = "Super Admin"
	Member     UserRole = "member"
)

type User struct {
	Id          string        `gorm:"type:varchar(30);not null;primaryKey" json:"id"`
	Role        UserRole      `gorm:"type:enum('Admin','Super Admin','member');not null" json:"role"`
	FirstName   string        `gorm:"type:varchar(255);not null" json:"first_name"`
	LastName    string        `gorm:"type:varchar(255);not null" json:"last_name"`
	Password    string        `gorm:"type:varchar(255);not null" json:"password"`
	Tasks       []Task        `gorm:"constraint:OnDelete:CASCADE" json:"tasks,omitempty"`
	UserHistory []UserHistory `gorm:"constraint:OnDelete:CASCADE" json:"user_history,omitempty"`
	UserEmails  []UserEmail   `gorm:"constraint:OnDelete:CASCADE" json:"user_emails,omitempty"`
}

func (u *User) AfterDelete(tx *gorm.DB) (err error) {
	tx.Clauses(clause.Returning{}).Where("user_id = ?", u.Id).Delete(&Task{})
	tx.Clauses(clause.Returning{}).Where("user_id = ?", u.Id).Delete(&UserHistory{})
	tx.Clauses(clause.Returning{}).Where("user_id = ?", u.Id).Delete(&UserEmail{})
	return
}
func (u *User) GenerateSecureUserId() (string, error) {
	uuid, err := uuid.NewRandom()
	if err != nil {
		return "", err
	}
	id := fmt.Sprintf("U%s", uuid.String())
	return id, nil
}
