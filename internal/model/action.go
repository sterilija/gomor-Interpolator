package model

type Action struct {
	ActionID   int    `json:"action_id" gorm:"id"`
	ActionType string `json:"action_type" gorm:"action_type"`
	Approved   bool   `json:"approved" gorm:"approved"`
	UserID     int    `json:"user_id" gorm:"user_id"`
}
