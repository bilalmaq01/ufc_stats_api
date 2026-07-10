package models

type Fighter struct {
	ID          int     `json:"id"`
	Name        string  `json:"name"`
	Nickname    *string `json:"nickname"`
	Height      *string `json:"height"`
	WeightClass *string `json:"weight_class"`
	ReachIn     int     `json:"reach_in"`
	Wins        int     `json:"wins"`
	Losses      int     `json:"losses"`
	Draws       int     `json:"draws"`
}
