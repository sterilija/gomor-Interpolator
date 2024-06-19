package model

type Legend struct {
	LegendID    int    `json:"legend_id" gorm:"id"`
	Legend      string `json:"legend" gorm:"legend"`
	UserID      int    `json:"user_id" gorm:"user_id"`
	VictimCorps string `json:"victim_corps" gorm:"victim_corps"`
	VictimTeams string `json:"victim_teams" gorm:"victim_teams"`
	Proofs      string `json:"proofs" gorm:"proofs"`
}
