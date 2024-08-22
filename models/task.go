package models

import (
	"time"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type Task struct {
	Id          string        `gorm:"type:varchar(30);not null;primaryKey" json:"id"`
	UserId      string        `gorm:"type:varchar(30);not null" json:"user_id"`
	Title       string        `gorm:"type:varchar(255);not null" json:"title"`
	Description string        `gorm:"type:text;not null" json:"description"`
	Status      string        `gorm:"type:varchar(255);not null" json:"status"`
	Reason      string        `gorm:"type:text;not null" json:"reason"`
	Revision    int8          `gorm:"type:int;not null" json:"revision"`
	DueDate     time.Time     `gorm:"not null" json:"due_date"`
	User        User          `gorm:"foreignKey:UserId" json:"user,omitempty"`
	TaskHistory []TaskHistory `gorm:"constraint:OnDelete:CASCADE" json:"task_history,omitempty"`
}

func (t *Task) AfterDelete(tx *gorm.DB) (err error) {
	tx.Clauses(clause.Returning{}).Where("task_id = ?", t.Id).Delete(&TaskHistory{})
	return
}
