package models

type Player struct {
	Id int32 `json:"player_id"`
	Name string `json:"player_name"`
	Password string `json:"player_password"`
}