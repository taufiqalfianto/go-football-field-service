package models

import "time"


type Role struct {
	ID int `gorm:"primarykey;autoIncrement"`
	Code string `gorm:"varchar(15);not null"`
	Name string `gorm:"varchar(15);not null"`
	CreatedAt *time.Time
	UpdatedAt *time.Time
} 