package model

import "time"

// Product represents model of application resources.
type Product struct {
	ID          int    `gorm:"primary_key"`
	Name        string `gorm:"not null" json:"name"`
	Description string
	Price       float32
	Sponsor     string `json:"sponsor"`
	CreatedAt   time.Time
}
