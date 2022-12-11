package model

import "time"

type Submission struct {
	Id           uint64    `gorm:"column:id;AUTO_INCREMENT;PRIMARY_KEY" json:"id"`
	RoomNumber   string    `gorm:"column:room_number" json:"room_number"`
	Remark       string    `gorm:"column:remark" json:"remark"`
	StartUseDate string    `gorm:"column:start_use_date" json:"start_use_date"`
	EndUseDate   string    `gorm:"column:end_use_date" json:"end_use_date"`
	IsApproved   bool      `gorm:"column:is_approved" json:"is_approved"`
	RoomID       uint64    `gorm:"column:room_id" json:"room_id"`
	ApprovedBy   uint64    `gorm:"column:approved_by" json:"approved_by"`
	CreatedAt    time.Time `gorm:"column:created_at;" json:"created_at"`
	UpdatedAt    time.Time `gorm:"column:updated_at;" json:"updated_at"`
}

func (Submission) TableName() string {
	return "submission"
}
