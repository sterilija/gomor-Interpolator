package model

type Tags struct {
	UserID    int    `json:"victim_tag_id" gorm:"user_id"`
	Id        int    `json:"victim_id" gorm:"id"`
	VictimTag string `json:"tag" gorm:"tag"`
}
