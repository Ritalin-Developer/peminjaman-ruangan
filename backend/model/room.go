package model

import "time"

type Room struct {
	Id          uint64    `gorm:"column:id;AUTO_INCREMENT;PRIMARY_KEY" json:"id"`
	IsAvailable bool      `gorm:"column:is_available" json:"is_available"`
	RoomNumber  string    `gorm:"column:room_number" json:"room_number"`
	Remark      string    `gorm:"column:remark" json:"remark"`
	CreatedAt   time.Time `gorm:"column:created_at;" json:"created_at"`
	UpdatedAt   time.Time `gorm:"column:updated_at;" json:"updated_at"`
}

func (Room) TableName() string {
	return "room"
}
