package model

import "time"

type School struct {
	ID           uint      `gorm:"column:id;auto_increment;not null;primary_key" json:"id"`
	Code         string    `gorm:"column:code;type:varchar(255);unique;not null" json:"code"`
	Name         string    `gorm:"column:name;type:varchar(255);" json:"name"`
	ShortName    string    `gorm:"column:short_name;type:varchar(31);" json:"short_name"`
	Introduction string    `gorm:"column:introduction;type:varchar(1023);" json:"introduction"`
	Address      string    `gorm:"column:address;type:varchar(255);" json:"address"`
	AvatarUrl    string    `gorm:"column:avatar_url;type:varchar(255);" json:"avatar_url"`
	Founding     time.Time `gorm:"column:founding_at;type:timestamp;" json:"founding_at"`
	UserID       uint      `gorm:"column:user_id;type:integer" json:"userId"`
	User         User      `gorm:"foreignkey:id;association_foreginkey:user_id" json:"user"`
	CreatedAt    time.Time `gorm:"column:created_at;type:timestamp;not null" json:"created_at"`
	UpdatedAt    time.Time `gorm:"column:updated_at;type:timestamp;not null" json:"updated_at"`
}
