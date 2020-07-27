package model

import (
	"time"
)

type User struct {
	ID        uint      `gorm:"column:id;auto_increment;unique;not null;primary_key" json:"id"`
	UserName  string    `gorm:"column:username;type:varchar(255);unique;not null" json:"name"`
	Password  string    `gorm:"column:password;type:varchar(255);not null" json:"password"`
	Status    bool      `gorm:"column:status;type:varchar(31);not null" json:"status"`
	Role      int8      `gorm:"column:role;type:varchar(31);not null" json:"role"`
	CreatedAt time.Time `gorm:"column:created_at;type:timestamp;not null" json:"created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at;type:timestamp;not null" json:"updated_at"`
}
