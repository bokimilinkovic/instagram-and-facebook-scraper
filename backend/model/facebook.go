package model

// FacebookResponse represents response from facebook api.
type FacebookResponse struct {
	ID        uint   `gorm:"primary_key"`
	Username  string `gorm:"not null"`
	Likes     string
	Followers string
}
