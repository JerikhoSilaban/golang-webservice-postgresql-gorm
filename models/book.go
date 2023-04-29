package models

import (
	"errors"
	"fmt"
	"time"

	"gorm.io/gorm"
)

type Book struct {
	ID        uint   `gorm:"primaryKey"`
	Title     string `gorm:"not null; unique; type:varchar(256)"`
	Publisher string `gorm:"not null; type:varchar(256)"`
	AuthorID  uint
	CreatedAt time.Time
	UpdatedAt time.Time
}

// Hook Function
func (b *Book) BeforeCreate(tx *gorm.DB) (err error) {
	fmt.Println("Book BeforeCreate()")

	if len(b.Title) < 2 {
		err = errors.New("book title is too short")
	}

	return
}
