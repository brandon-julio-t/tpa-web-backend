package commands

import (
	"bytes"
	"fmt"
	"github.com/brandon-julio-t/tpa-web-backend/facades"
	"github.com/brandon-julio-t/tpa-web-backend/graph/models"
	"github.com/brandon-julio-t/tpa-web-backend/services/notification_service"
	"gorm.io/gorm"
	"html/template"
	"math"
	"path/filepath"
	"time"
)

type GameGiftCommand struct {
	DB    *gorm.DB
	User  models.User
	Games []*models.Game
	Input models.Gift
}

func (g GameGiftCommand) Execute() error {
	var friend models.User
	if err := g.DB.First(&friend, g.Input.UserID).Error; err != nil {
		return err
	}

	header := &models.GameGiftTransactionHeader{
		GameGiftTransactionHeaderUserID:   g.User.ID,
		GameGiftTransactionHeaderUser:     g.User,
		GameGiftTransactionHeaderFriendID: friend.ID,
		GameGiftTransactionHeaderFriend:   friend,
		Message:                           g.Input.Message,
		Sentiment:                         g.Input.Sentiment,
		Signature:                         g.Input.Signature,
	}

	if err := g.DB.Create(header).Error; err != nil {
		return err
	}

	var details []*models.GameGiftTransactionDetail
	for _, game := range g.Games {
		details = append(details, &models.GameGiftTransactionDetail{
			GameGiftTransactionHeaderID:     header.ID,
			GameGiftTransactionDetailGameID: game.ID,
			GameGiftTransactionDetailGame:   *game,
		})
	}

	if err := g.DB.Create(&details).Error; err != nil {
		return err
	}

	senderHtmlPath := filepath.Join("email-templates", "game-gift-sender.html")
	senderHtml, err := template.ParseFiles(senderHtmlPath)
	if err != nil {
		return err
	}

	var senderGames []mailSenderGamesData
	var grandTotal float64
	for _, game := range g.Games {
		senderGames = append(senderGames, mailSenderGamesData{
			Title: game.Title,
			Price: fmt.Sprintf("%.2f", game.Price),
		})
		grandTotal += game.Price
	}

	var senderHtmlOutput bytes.Buffer
	if err := senderHtml.Execute(&senderHtmlOutput, mailSenderData{
		User:        g.User.DisplayName,
		Friend:      friend.DisplayName,
		Games:       senderGames,
		GrandTotal:  fmt.Sprintf("%.2f", grandTotal),
		AccountName: g.User.AccountName,
		CreatedAt:   header.CreatedAt.Format(time.RFC850),
	}); err != nil {
		return err
	}

	receiverHtmlPath := filepath.Join("email-templates", "game-gift-receiver.html")
	receiverHtml, err := template.ParseFiles(receiverHtmlPath)
	if err != nil {
		return err
	}

	var receiverGames []mailReceiverGamesData
	for _, game := range g.Games {
		receiverGames = append(receiverGames, mailReceiverGamesData{Title: game.Title})
	}

	var receiverHtmlOutput bytes.Buffer
	if err := receiverHtml.Execute(&receiverHtmlOutput, mailReceiverData{
		Friend:    g.User.DisplayName,
		Games:     receiverGames,
		FirstName: g.Input.FirstName,
		Message:   g.Input.Message,
		Sentiment: g.Input.Sentiment,
	}); err != nil {
		return err
	}

	senderEmail := g.User.Email
	if err := facades.UseMail().SendHTML(senderHtmlOutput.String(), "STAEM Gift", "STAEMGiftSender", senderEmail); err != nil {
		return err
	}

	receiverEmail := friend.Email
	if err := facades.UseMail().SendHTML(receiverHtmlOutput.String(), "STAEM Gift", "STAEMGiftReceiver", receiverEmail); err != nil {
		return err
	}

	if err := notification_service.Notify(&friend, fmt.Sprintf("%v sent you a gift", g.User.DisplayName)); err != nil {
		return err
	}

	if err := g.DB.Model(&g.User).Association("UserCart").Clear(); err != nil {
		return err
	}


	g.User.Exp += 50
	g.User.Points += int64(math.Round(grandTotal / 106 * 15000))
	return g.DB.Save(&g.User).Error
}

type mailSenderGamesData struct {
	Title string
	Price string
}

type mailSenderData struct {
	User        string
	Friend      string
	Games       []mailSenderGamesData
	GrandTotal  string
	AccountName string
	CreatedAt   string
}

type mailReceiverGamesData struct {
	Title string
}

type mailReceiverData struct {
	Friend    string
	Games     []mailReceiverGamesData
	FirstName string
	Message   string
	Sentiment string
}
