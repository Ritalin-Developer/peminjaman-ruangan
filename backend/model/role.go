package model

import "time"

type Role struct {
	Id        uint64    `gorm:"column:id;AUTO_INCREMENT;PRIMARY_KEY" json:"id"`
	RoleName  string    `gorm:"column:role_name" json:"role_name"`
	CreatedAt time.Time `gorm:"column:created_at;" json:"created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at;" json:"updated_at"`
}

func (Role) TableName() string {
	return "role"
}
