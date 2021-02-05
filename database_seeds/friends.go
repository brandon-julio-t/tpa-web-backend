package database_seeds

import (
	"github.com/brandon-julio-t/tpa-web-backend/commands"
	"github.com/brandon-julio-t/tpa-web-backend/facades"
	"github.com/brandon-julio-t/tpa-web-backend/graph/models"
)

func SeedFriends() error {
	user := new(models.User)
	if err := facades.UseDB().First(user, 2).Error; err != nil {
		return err
	}

	friends := make([]*models.User, 0)
	if err := facades.UseDB().Find(&friends, "id not in (?, ?)", 1, 2).Error; err != nil {
		return err
	}

	for _, friend := range friends {
		if err := (commands.BefriendCommand{
			DB:     facades.UseDB(),
			User:   user,
			Friend: friend,
		}.Execute()); err != nil {
			return err
		}
	}

	return nil
}
