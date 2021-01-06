package dto

import "holycode-task/model"

// SocialMediaDto represents combined response from instragram and facebook api.
type SocialMediaDto struct {
	Username      string `json:"username"`
	FullName      string `json:"full_name"`
	Biography     string `json:"biography"`
	ProfilePicURL string `json:"profile_pic_url"`
	Email         string `json:"email"`
	MediaCount    int    `json:"media_count"`
	FollowerCount int    `json:"follower_count"`
	UsertagsCount int    `json:"usertags_count"`
	Likes         string `json:"likes"`
	Followers     string `json:"followers"`
}

func CreateSocialMediaDto(insta *model.InstagramAccount, fb *model.FacebookResponse) *SocialMediaDto {
	return &SocialMediaDto{
		Username:      insta.Username,
		FullName:      insta.FullName,
		Biography:     insta.Biography,
		ProfilePicURL: insta.ProfilePicURL,
		Email:         insta.Email,
		MediaCount:    insta.MediaCount,
		FollowerCount: insta.FollowerCount,
		UsertagsCount: insta.UsertagsCount,
		Likes:         fb.Likes,
		Followers:     fb.Followers,
	}
}
