package models

import (
	"time"
)

type Profile struct {
	ID          int       `gorm:"primary_key;auto_increment;" json:"id"`
	Name        string    `gorm:"size:255;not null;unique;" json:"name"`
	DateOfBirth string    `gorm:"null" json:"dateofbirth"`
	Status      string    `gorm:"default:'ACTIVE';" json:"status"`
	CreatedAt   time.Time `gorm:"default:CURRENT_TIMESTAMP;" json:"created_at"`
	UpdatedAt   time.Time `gorm:"default:CURRENT_TIMESTAMP;" json:"updated_at"`
}
