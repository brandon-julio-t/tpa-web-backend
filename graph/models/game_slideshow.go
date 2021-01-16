package models

import "encoding/base64"

type GameSlideshow struct {
	BaseModel
	GameID int64
	Game   Game
	File   []byte
}

func (gs *GameSlideshow) FileBase64() string {
	return base64.StdEncoding.EncodeToString(gs.File)
}
