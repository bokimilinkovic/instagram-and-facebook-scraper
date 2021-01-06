package model

import "time"

// User represents user model.
type User struct {
	ID        int    `gorm:"primary_key"`
	Username  string `gorm:"not null" json:"username"`
	Password  string
	CreatedAt time.Time
}
