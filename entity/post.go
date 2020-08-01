package entity

import (
	"time"
)

type Post struct {
	ID        uint      `gorm:"column:id;auto_increment;unique;not null;primary_key" json:"id"`
	Title     string    `gorm:"column:title;type:varchar(255);" json:"title"`
	Content   string    `gorm:"column:content;type:varchar(8191);" json:"content"`
	Status    bool      `gorm:"column:status;type:varchar(31);not null" json:"status"`
	UserID    uint      `gorm:"column:user_id;type:integer" json:"userId"`
	User      User      `gorm:"foreignkey:id;association_foreginkey:user_id" json:"user"`
	CreatedAt time.Time `gorm:"column:created_at;type:timestamp;not null" json:"created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at;type:timestamp;not null" json:"updated_at"`
}
