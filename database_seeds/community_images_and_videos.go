package database_seeds

import (
	"github.com/brandon-julio-t/tpa-web-backend/facades"
	"github.com/brandon-julio-t/tpa-web-backend/graph/models"
	"io/ioutil"
	"path/filepath"
	"syreclabs.com/go/faker"
	"time"
)

func SeedCommunityImagesAndVideos() error {
	for i := 0; i < 20; i++ {
		isImage := faker.Number().NumberInt(1)%2 == 0
		file := make([]byte, 0)
		contentType := ""

		if isImage {
			path := filepath.Join("assets", "Background Zoom SLC.png")
			image, err := ioutil.ReadFile(path)
			if err != nil {
				return err
			}

			file = image
			contentType = "image/png"
		} else {
			path := filepath.Join("assets", "file_example_MP4_480_1_5MG.mp4")
			video, err := ioutil.ReadFile(path)
			if err != nil {
				return err
			}

			file = video
			contentType = "video/mp4"
		}

		user := new(models.User)
		if err := facades.UseDB().Order("random()").First(user).Error; err != nil {
			return err
		}

		now := time.Now()
		before := now.AddDate(0, 0, faker.Number().NumberInt(1)*-1)

		if err := facades.UseDB().Create(&models.CommunityImageAndVideo{
			CreatedAt:   faker.Time().Between(before, now),
			Description: faker.Lorem().Paragraph(faker.Number().NumberInt(1)),
			Name:        faker.Lorem().Sentence(faker.Number().NumberInt(1)),
			File_: models.AssetFile{
				File:        file,
				ContentType: contentType,
			},
			User_: *user,
		}).Error; err != nil {
			return err
		}
	}

	return nil
}
