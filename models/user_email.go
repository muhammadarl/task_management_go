package models

import (
	"fmt"

	"github.com/google/uuid"
)

type EmailType string

const (
	Work EmailType = "Work"
	Home EmailType = "Home"
)

type UserEmail struct {
	Id            string    `gorm:"type:varchar(30);not null;primaryKey" json:"id"`
	UserId        string    `gorm:"type:varchar(30);not null" json:"user_id"`
	UemailAddress string    `gorm:"type:varchar(255);not null" json:"uemail_address"`
	UemailType    EmailType `gorm:"type:enum('Work','Home');not null" json:"uemail_type"`
	IsPrimary     bool      `gorm:"type:boolean;not null" json:"is_primary"`
	User          User      `gorm:"foreignKey:UserId" json:"user,omitempty"`
}

func (u *UserEmail) GenerateSecureUserEmailId() (string, error) {
	uuid, err := uuid.NewRandom()
	if err != nil {
		return "", err
	}
	id := fmt.Sprintf("UE%s", uuid.String())
	return id, nil
}
