package models

import "encoding/base64"

type Game struct {
	BaseModel
	Title              string
	Description        string
	Price              float64
	Banner             []byte
	Slideshows         []*GameSlideshow
	GameTags           []*GameTag `gorm:"many2many:game_tag_mappings;"`
	SystemRequirements string
}

func (g *Game) BannerBase64() string {
	return base64.StdEncoding.EncodeToString(g.Banner)
}
