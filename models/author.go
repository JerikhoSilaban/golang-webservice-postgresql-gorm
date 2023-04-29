package models

import "time"

type Author struct {
	ID        uint   `gorm:"primaryKey"`
	Name      string `gorm:"not null; unique; type:varchar(255)"`
	Books     []Book
	CreatedAt time.Time
	UpdatedAt time.Time
}
