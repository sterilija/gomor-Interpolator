package model

type User struct {
	UserID       int    `json:"user_id" gorm:"id"`
	Username     string `json:"username" gorm:"username"`
	Role         string `json:"role" gorm:"role"`
	Active       bool   `json:"active" gorm:"active"`
	ShadowRating int    `json:"shadow_rating" gorm:"shadow_rating"`
}
