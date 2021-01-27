package database_seeds

import (
	"fmt"
	"github.com/brandon-julio-t/tpa-web-backend/facades"
	"github.com/brandon-julio-t/tpa-web-backend/graph/models"
	"gorm.io/gorm"
	"io/ioutil"
	"path/filepath"
	"strconv"
	"syreclabs.com/go/faker"
	"time"
)

func SeedGames() error {
	imagePath := filepath.Join("assets", "Background Zoom SLC.png")
	backgroundSLC, err := ioutil.ReadFile(imagePath)
	if err != nil {
		return err
	}

	videoPath := filepath.Join("assets", "file_example_MP4_480_1_5MG.mp4")
	sampleVideo, err := ioutil.ReadFile(videoPath)
	if err != nil {
		return err
	}

	return facades.UseDB().Transaction(func(tx *gorm.DB) error {
		for i := 0; i < 50; i++ {
			hoursPlayed, err := strconv.ParseFloat(faker.Number().Decimal(5, 2), 64)
			if err != nil {
				return err
			}

			var tags []*models.GameTag
			counts := faker.Number().NumberInt(1)
			for i := 0; i < counts; i++ {
				id, err := strconv.ParseInt(faker.Number().Between(1, 424), 10, 64)
				if err != nil {
					return err
				}
				tags = append(tags, &models.GameTag{ID: id})
			}

			genreId, err := strconv.ParseInt(faker.Number().Between(1, 12), 10, 64)
			if err != nil {
				return err
			}

			discount := float64(0)

			if faker.Number().NumberInt(1) % 2 == 0 {
				discf, err := strconv.ParseFloat(faker.Number().Between(1, 80), 64)
				if err != nil {
					return err
				}

				discount = discf / float64(100)
			}

			if err := tx.Create(&models.Game{
				Banner:          models.AssetFile{File: backgroundSLC, ContentType: "image/png"},
				CreatedAt:       time.Time{},
				Description:     faker.Lorem().Paragraph(faker.Number().NumberInt(1)),
				Discount:        discount,
				GameTags:        tags,
				GenreID:         genreId,
				HoursPlayed:     hoursPlayed,
				IsInappropriate: faker.Number().NumberInt(1)%2 == 0,
				Price:           float64(faker.Commerce().Price()),
				GameGameReviews: []*models.GameReview{
					{
						GameReviewUserID: 2,
						Content:          faker.Lorem().Paragraph(faker.Number().NumberInt(1)),
						IsRecommended:    faker.Number().NumberInt(1)%2 == 0,
					},
					{
						GameReviewUserID: 2,
						Content:          faker.Lorem().Paragraph(faker.Number().NumberInt(1)),
						IsRecommended:    faker.Number().NumberInt(1)%2 == 0,
					},
					{
						GameReviewUserID: 2,
						Content:          faker.Lorem().Paragraph(faker.Number().NumberInt(1)),
						IsRecommended:    faker.Number().NumberInt(1)%2 == 0,
					},
					{
						GameReviewUserID: 2,
						Content:          faker.Lorem().Paragraph(faker.Number().NumberInt(1)),
						IsRecommended:    faker.Number().NumberInt(1)%2 == 0,
					},
					{
						GameReviewUserID: 2,
						Content:          faker.Lorem().Paragraph(faker.Number().NumberInt(1)),
						IsRecommended:    faker.Number().NumberInt(1)%2 == 0,
					},
				},
				GameSlideshows: []*models.GameSlideshow{
					{GameSlideshowFile: models.AssetFile{File: backgroundSLC, ContentType: "image/png"}},
					{GameSlideshowFile: models.AssetFile{File: backgroundSLC, ContentType: "image/png"}},
					{GameSlideshowFile: models.AssetFile{File: sampleVideo, ContentType: "video/mp4"}},
				},
				SystemRequirements: `
MINIMUM:
Requires a 64-bit processor and operating system
OS: Windows 7 or 10
Processor: Intel Core i5-3570K or AMD FX-8310
Memory: 8 GB RAM Graphics: NVIDIA GeForce GTX 780 or AMD Radeon RX 470
DirectX: Version 12
Storage: 70 GB available space
Additional Notes: In this game you will encounter a variety of visual effects that may provide seizures or loss of consciousness in a minority of people. If you or someone you know experiences any of the above symptoms while playing, stop and seek medical attention immediately.

RECOMMENDED:
Requires a 64-bit processor and operating system
OS: Windows 10
Processor: Intel Core i7-4790 or AMD Ryzen 3 3200G
Memory: 12 GB RAM
Graphics: GTX 1060 6GB / GTX 1660 Super or Radeon RX 590
DirectX: Version 12
Storage: 70 GB available space
Additional Notes: SSD recommended
`,
				Title: fmt.Sprintf("%v %v", faker.App().Name(), faker.Name().LastName()),
			}).Error; err != nil {
				return err
			}
		}

		return nil
	})
}
