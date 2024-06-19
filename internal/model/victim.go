package model

type Victim struct {
	VictimID           int    `json:"victim_id" gorm:"id"`
	VictimType         string `json:"victim_type" gorm:"victim_type"`
	VictimName         string `json:"victim_name" gorm:"victim_name"`
	VictimAVGRating    int    `json:"victim_avg_rating" gorm:"victim_avg_rating"`
	VictimMedianRating int    `json:"victim_median_rating" gorm:"victim_median_rating"`
	VictimTagsId       string `json:"victim_tags_id" gorm:"victim_tags_id"`
	VictimLegendID     int    `json:"victim_legend_id" gorm:"victim_legend_id"`
	UserID             int    `json:"user_id" gorm:"user_id"`
}
