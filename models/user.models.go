package models

import (
	"time"
)

// User DB struct
type User struct {
	ID        uint `gorm:"primarykey"`
	Username string `gorm:"uniqueIndex;not null" json:"username"`
	Email    string `gorm:"uniqueIndex;not null" json:"email"`
	Password string `gorm:"not null" json:"password"`
	Name     string `gorm:"uniqueIndex;not null" json:"name"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

// User Data struct
type UserData struct {
	ID       uint   `json:"id"`
	Username string `json:"username"`
	Name     string `json:"name"`
	Email    string `json:"email"`
}
