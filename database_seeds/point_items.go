package database_seeds

import (
	"github.com/brandon-julio-t/tpa-web-backend/facades"
	"github.com/brandon-julio-t/tpa-web-backend/graph/models"
	"io/ioutil"
	"path/filepath"
	"syreclabs.com/go/faker"
)

func SeedPointItems() error {
	animatedAvatars := []string{
		"animated_avatar_1.gif",
		"animated_avatar_2.gif",
		"animated_avatar_3.gif",
		"animated_avatar_4.gif",
		"animated_avatar_5.gif",
	}

	if err := saveAsset(animatedAvatars, "animated_avatar", true, "gif"); err != nil {
		return err
	}

	avatarBorders := []string{
		"avatar_border_1.png",
		"avatar_border_2.png",
		"avatar_border_3.png",
		"avatar_border_4.png",
		"avatar_border_5.png",
	}

	if err := saveAsset(avatarBorders, "avatar_border", true, "png"); err != nil {
		return err
	}

	miniProfileBackgrounds := []string{
		"mini_profile_background_1.webm",
		"mini_profile_background_2.webm",
		"mini_profile_background_3.webm",
		"mini_profile_background_4.webm",
		"mini_profile_background_5.webm",
	}

	if err := saveAsset(miniProfileBackgrounds, "mini_profile_background", false, "webm"); err != nil {
		return err
	}

	profileBackgrounds := []string{
		"profile_background_1.webm",
		"profile_background_2.webm",
		"profile_background_3.webm",
		"profile_background_4.webm",
		"profile_background_5.webm",
	}

	if err := saveAsset(profileBackgrounds, "profile_background", false, "webm"); err != nil {
		return err
	}

	stickers := []string{
		"sticker_1.png",
		"sticker_2.png",
		"sticker_3.png",
		"sticker_4.png",
		"sticker_5.png",
	}

	return saveAsset(stickers, "sticker", true, "png")
}

func saveAsset(files []string, itemType string, isImage bool, extension string) error {
	for _, file := range files {
		path := filepath.Join("assets", file)
		data, err := ioutil.ReadFile(path)
		if err != nil {
			return err
		}

		fileType := ""
		if isImage {
			fileType = "image"
		} else {
			fileType = "video"
		}

		if err := facades.UseDB().Create(&models.PointItem{
			Name:     faker.App().Name(),
			Category: itemType,
			Price:    2000,
			Image_: models.AssetFile{
				File:        data,
				ContentType: fileType + "/" + extension,
			},
		}).Error; err != nil {
			return err
		}
	}

	return nil
}
