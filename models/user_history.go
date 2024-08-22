package models

import (
	"fmt"

	"github.com/google/uuid"
)

type UserHistory struct {
	Id             string `gorm:"type:varchar(30);not null;primaryKey" json:"id"`
	UserId         string `gorm:"type:varchar(30);not null" json:"user_id"`
	UhistoryAction string `gorm:"type:varchar(255);not null" json:"uhistory_action"`
	UhistoryTime   string `gorm:"type:datetime" json:"uhistory_time"`
	UhistoryRemark string `gorm:"type:text;not null" json:"uhistory_remark"`
	User           User   `gorm:"foreignKey:UserId" json:"user,omitempty"`
}

func (u *UserHistory) GenerateSecureUserHistoryId() (string, error) {
	uuid, err := uuid.NewRandom()
	if err != nil {
		return "", err
	}
	id := fmt.Sprintf("UH%s", uuid.String())
	return id, nil
}
