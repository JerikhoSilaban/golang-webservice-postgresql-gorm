package models

import "time"

// Book represents the model of Author
type Author struct {
	ID        uint   `gorm:"primaryKey"`
	Name      string `gorm:"not null; unique; type:varchar(255)"`
	Books     []Book
	CreatedAt time.Time
	UpdatedAt time.Time
}
