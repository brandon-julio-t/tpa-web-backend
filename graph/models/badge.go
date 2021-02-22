package models

type Badge struct {
	ID     int64
	Exp    int64
	GameID int64
	Game   Game
	Level  int64
	Name   string
}
