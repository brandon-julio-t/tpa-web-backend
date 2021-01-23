//go:generate go run github.com/99designs/gqlgen

package resolvers

import (
	"github.com/brandon-julio-t/tpa-web-backend/graph/models"
	"sync"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	PrivateChatSockets map[int64]chan *models.PrivateMessage
	RTCConnections     map[string]string
	RTCJoinSockets     map[string]chan string
	StreamingRooms     map[string][]chan string
	Mutex              sync.Mutex
}

func NewResolver() *Resolver {
	return &Resolver{
		PrivateChatSockets: map[int64]chan *models.PrivateMessage{},
		RTCConnections:     map[string]string{},
		RTCJoinSockets:     map[string]chan string{},
		StreamingRooms:     map[string][]chan string{},
	}
}
