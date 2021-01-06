package model

// InstagramAccount represents response from instagram api.
type InstagramAccount struct {
	ID            uint   `gorm:"primary_key"`
	Username      string `json:"username"`
	FullName      string `json:"full_name"`
	Biography     string `json:"biography"`
	ProfilePicURL string `json:"profile_pic_url"`
	Email         string `json:"email"`
	MediaCount    int    `json:"media_count"`
	FollowerCount int    `json:"follower_count"`
	UsertagsCount int    `json:"usertags_count"`
}
