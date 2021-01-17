package models

import "encoding/base64"

type GameSlideshow struct {
	ID          int64 `gorm:"primaryKey;autoIncrement:true"`
	GameID      int64
	Game        Game
	File        []byte
	ContentType string
}

func (gs *GameSlideshow) FileBase64() string {
	return base64.StdEncoding.EncodeToString(gs.File)
}
