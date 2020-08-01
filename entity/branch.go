package entity

import "time"

type Branch struct {
	ID          uint      `gorm:"column:id;auto_increment;unique;not null;primary_key" json:"id"`
	Name        string    `gorm:"column:name;type:varchar(31);not null" json:"password"`
	Description string    `gorm:"column:description;type:varchar(255);not null" json:"role"`
	CreatedAt   time.Time `gorm:"column:created_at;type:timestamp;not null" json:"created_at"`
	UpdatedAt   time.Time `gorm:"column:updated_at;type:timestamp;not null" json:"updated_at"`
}
