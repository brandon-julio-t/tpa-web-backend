package models

type Wishlist struct {
	UserID int64
	GameID int64
	User_  User `foreignKey:"UserID"`
	Game_  Game `foreignKey:"GameID"`
}
