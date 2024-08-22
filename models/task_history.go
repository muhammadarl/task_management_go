package models

type TaskHistory struct {
	Id             string `gorm:"type:varchar(30);not null;primaryKey" json:"id"`
	TaskId         string `gorm:"type:varchar(30);not null" json:"task_id"`
	ThistoryAction string `gorm:"type:varchar(255);not null" json:"thistory_action"`
	ThistoryTime   string `gorm:"type:datetime" json:"thistory_time"`
	ThistoryRemark string `gorm:"type:text;not null" json:"thistory_remark"`
	Task           Task   `gorm:"foreignKey:TaskId" json:"task,omitempty"`
}
