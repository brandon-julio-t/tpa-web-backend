package commands

import (
	"github.com/brandon-julio-t/tpa-web-backend/graph/models"
	"gorm.io/gorm"
)

type BefriendCommand struct {
	DB     *gorm.DB
	User   *models.User
	Friend *models.User
}

func (c BefriendCommand) Execute() error {
	if err := c.DB.Create(&models.Friendship{
		User:   *c.User,
		Friend: *c.Friend,
	}).Error; err != nil {
		return err
	}

	if err := c.DB.Create(&models.Friendship{
		User:   *c.Friend,
		Friend: *c.User,
	}).Error; err != nil {
		return err
	}

	if err := c.DB.
		Where("user_id = ? and friend_id = ?", c.User.ID, c.Friend.ID).
		Or("user_id = ? and friend_id = ?", c.Friend.ID, c.User.ID).
		Unscoped().
		Delete(new(models.FriendRequest)).
		Error; err != nil {
		return err
	}

	return nil
}
