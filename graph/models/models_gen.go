// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package models

import (
	"github.com/99designs/gqlgen/graphql"
)

type CreateGame struct {
	Title              string            `json:"title"`
	Description        string            `json:"description"`
	Price              float64           `json:"price"`
	Banner             graphql.Upload    `json:"banner"`
	Slideshows         []*graphql.Upload `json:"slideshows"`
	GameTags           []int64           `json:"gameTags"`
	SystemRequirements string            `json:"systemRequirements"`
}

type UpdateGame struct {
	Title              string            `json:"title"`
	Description        string            `json:"description"`
	Price              float64           `json:"price"`
	Banner             *graphql.Upload   `json:"banner"`
	Slideshows         []*graphql.Upload `json:"slideshows"`
	GameTags           []int64           `json:"gameTags"`
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
