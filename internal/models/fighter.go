package models

import "time"

type Fighter struct {
	ID          int        `json:"id"`
	Name        string     `json:"name"`
	Nickname    *string    `json:"nickname"`
	Height      *string    `json:"height"`
	WeightClass *string    `json:"weight"`
	ReachIn     int        `json:"reach"`
	Wins        int        `json:"wins"`
	Losses      int        `json:"losses"`
	Draws       int        `json:"draws"`
	Stance      *string    `json:"stance"`
	DOB         *time.Time `json:"dob"`
	SLPM        float64    `json:"slpm"`
	StrAcc      int        `json:"str_acc"`
	SAPM        float64    `json:"sapm"`
	StrDef      int        `json:"str_def"`
	TdAvg       float64    `json:"td_avg"`
	TdAcc       int        `json:"td_acc"`
	TdDef       int        `json:"td_def"`
	SubAvg      float64    `json:"sub_avg"`
	URL         string     `json:"url"`
}
