// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package models

import (
	"github.com/99designs/gqlgen/graphql"
)

type AddMarketItemOffer struct {
	Category     string  `json:"category"`
	MarketItemID int64   `json:"marketItemId"`
	Price        float64 `json:"price"`
	Quantity     int64   `json:"quantity"`
}

type CommunityDiscussionCommentPagination struct {
	Data       []*CommunityDiscussionComment `json:"data"`
	TotalPages int64                         `json:"totalPages"`
}

type CommunityImageAndVideoCommentPagination struct {
	Data       []*CommunityImageAndVideoComment `json:"data"`
	TotalPages int64                            `json:"totalPages"`
}

type CreateCommunityImageAndVideo struct {
	Description string         `json:"description"`
	File        graphql.Upload `json:"file"`
	Name        string         `json:"name"`
}

type CreateGame struct {
	Title              string            `json:"title"`
	Description        string            `json:"description"`
	Price              float64           `json:"price"`
	Banner             graphql.Upload    `json:"banner"`
	Slideshows         []*graphql.Upload `json:"slideshows"`
	GameTags           []int64           `json:"gameTags"`
	GenreID            int64             `json:"genreId"`
	IsInappropriate    bool              `json:"isInappropriate"`
	SystemRequirements string            `json:"systemRequirements"`
}

type GamePagination struct {
	Data       []*Game `json:"data"`
	TotalPages int64   `json:"totalPages"`
}

type GameReviewCommentInput struct {
	ReviewID int64  `json:"reviewId"`
	Body     string `json:"body"`
}

type GameReviewCommentPagination struct {
	Data       []*GameReviewComment `json:"data"`
	TotalPages int64                `json:"totalPages"`
}

type Gift struct {
	UserID    int64  `json:"userId"`
	FirstName string `json:"firstName"`
	Message   string `json:"message"`
	Sentiment string `json:"sentiment"`
	Signature string `json:"signature"`
}

type MarketItemPagination struct {
	Data       []*MarketItem `json:"data"`
	TotalPages int64         `json:"totalPages"`
}

type MarketItemPrice struct {
	Price    float64 `json:"price"`
	Quantity int64   `json:"quantity"`
}

type PostCommunityDiscussion struct {
	GameID int64  `json:"gameId"`
	Body   string `json:"body"`
	Title  string `json:"title"`
}

type PostCommunityDiscussionComment struct {
	CommunityDiscussionID int64  `json:"communityDiscussionId"`
	Body                  string `json:"body"`
}

type PromoPagination struct {
	Data       []*Promo `json:"data"`
	TotalPages int64    `json:"totalPages"`
}

type UpdateGame struct {
	ID                 int64             `json:"id"`
	Title              string            `json:"title"`
	Description        string            `json:"description"`
	Price              float64           `json:"price"`
	Banner             *graphql.Upload   `json:"banner"`
	Slideshows         []*graphql.Upload `json:"slideshows"`
	GameTags           []int64           `json:"gameTags"`
	GenreID            int64             `json:"genreId"`
	IsInappropriate    bool              `json:"isInappropriate"`
	SystemRequirements string            `json:"systemRequirements"`
}

type UpdateUser struct {
	DisplayName  string          `json:"displayName"`
	RealName     string          `json:"realName"`
	CustomURL    string          `json:"customUrl"`
	Summary      string          `json:"summary"`
	CountryID    int64           `json:"countryId"`
	Avatar       *graphql.Upload `json:"avatar"`
	ProfileTheme string          `json:"profileTheme"`
}

type UserPagination struct {
	Data       []*User `json:"data"`
	TotalPages int64   `json:"totalPages"`
}
