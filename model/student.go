package model

import "time"

type Student struct {
	ID             uint      `gorm:"column:id;auto_increment;not null;primary_key" json:"id"`
	Code           string    `gorm:"column:code;type:varchar(255);unique;not null" json:"code"`
	FistName       string    `gorm:"column:first_name;type:varchar(255);" json:"first_name"`
	LastName       string    `gorm:"column:last_name;type:varchar(255);" json:"last_name"`
	Address        string    `gorm:"column:address;type:varchar(255);" json:"address"`
	AvatarUrl      string    `gorm:"column:avatar_url;type:varchar(255);" json:"avatar_url"`
	DateOfBirth    time.Time `gorm:"column:date_of_birth;type:timestamp;" json:"date_of_birth"`
	StudentGroupID *uint     `gorm:"column:customer_group_id;type:integer" json:"student_group_id"`
	UserID         uint      `gorm:"column:user_id;type:integer" json:"userId"`
	User           User      `gorm:"foreignkey:id;association_foreginkey:user_id" json:"user"`
	CreatedAt      time.Time `gorm:"column:created_at;type:timestamp;not null" json:"created_at"`
	UpdatedAt      time.Time `gorm:"column:updated_at;type:timestamp;not null" json:"updated_at"`
}
