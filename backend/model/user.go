package model

import "time"

type User struct {
	Id        uint64    `gorm:"column:id;AUTO_INCREMENT;PRIMARY_KEY" json:"id"`
	Username  string    `gorm:"column:username" json:"username"`
	Password  string    `gorm:"column:password" json:"password"`
	RealName  string    `gorm:"column:real_name" json:"real_name"`
	RoleID    uint64    `gorm:"column:role_id" json:"role_id"`
	CreatedAt time.Time `gorm:"column:created_at;" json:"created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at;" json:"updated_at"`
}

func (User) TableName() string {
	return "user"
}
