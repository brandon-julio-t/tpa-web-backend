package commands

import (
	"bytes"
	"fmt"
	"github.com/brandon-julio-t/tpa-web-backend/facades"
	"github.com/brandon-julio-t/tpa-web-backend/graph/models"
	"gorm.io/gorm"
	"html/template"
	"path/filepath"
	"time"
)

type GamePurchaseCommand struct {
	DB         *gorm.DB
	Games      []*models.Game
	User       models.User
	GrandTotal float64
}

func (g GamePurchaseCommand) Execute() error {
	header := &models.GamePurchaseTransactionHeader{
		GamePurchaseTransactionHeaderUserID: g.User.ID,
		GamePurchaseTransactionHeaderUser:   g.User,
		GrandTotal:                          g.GrandTotal,
	}

	if err := g.DB.Create(header).Error; err != nil {
		return err
	}

	var details []*models.GamePurchaseTransactionDetail
	for _, game := range g.Games {
		details = append(details, &models.GamePurchaseTransactionDetail{
			GamePurchaseTransactionHeaderID:     header.ID,
			GamePurchaseTransactionDetailGameID: game.ID,
			GamePurchaseTransactionDetailGame:   *game,
		})
	}

	if err := g.DB.Create(details).Error; err != nil {
		return err
	}

	templatePath := filepath.Join("email-templates", "game-purchase.html")
	html, err := template.ParseFiles(templatePath)
	if err != nil {
		return err
	}

	var receiptGame []struct {
		Title string
		Price string
	}

	for _, game := range g.Games {
		receiptGame = append(receiptGame, struct {
			Title string
			Price string
		}{
			Title: game.Title,
			Price: fmt.Sprintf("%.2f", game.Price),
		})
	}

	data := struct {
		User  string
		Games []struct {
			Title string
			Price string
		}
		GrandTotal  string
		AccountName string
		CreatedAt   string
	}{
		User:        g.User.DisplayName,
		Games:       receiptGame,
		GrandTotal:  fmt.Sprintf("%.2f", g.GrandTotal),
		AccountName: g.User.AccountName,
		CreatedAt:   header.CreatedAt.Format(time.RFC850),
	}

	var output bytes.Buffer
	if err := html.Execute(&output, data); err != nil {
		return err
	}

	if err := facades.UseMail().SendHTML(
		output.String(),
		"Game Purchase Receipt",
		"GamePurchaseReceipt",
		g.User.Email,
	); err != nil {
		return err
	}

	if err = g.DB.Model(&g.User).Association("UserCart").Clear(); err != nil {
		return err
	}

	g.User.Exp += 50
	return g.DB.Save(&g.User).Error
}
