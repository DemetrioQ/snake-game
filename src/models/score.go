package models

type Score struct{
	Id int32 `json:"score_id"`
	PlayerId int32 `json:"player_id"`
	Points int32 `json:"score"`
}
