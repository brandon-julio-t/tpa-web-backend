package resolvers

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/brandon-julio-t/tpa-web-backend/graph/models"
	"github.com/brandon-julio-t/tpa-web-backend/middlewares"
)

func (r *mutationResolver) StartStreaming(ctx context.Context, rtcConnection string) (string, error) {
	user, err := middlewares.UseAuth(ctx)
	if err != nil {
		return "", err
	}

	r.Mutex.Lock()
	r.RTCConnections[user.AccountName] = rtcConnection
	r.Mutex.Unlock()

	return rtcConnection, nil
}

func (r *mutationResolver) StopStreaming(ctx context.Context) (bool, error) {
	user, err := middlewares.UseAuth(ctx)
	if err != nil {
		return false, err
	}

	r.Mutex.Lock()
	delete(r.RTCConnections, user.AccountName)
	r.Mutex.Unlock()

	return true, nil
}

func (r *mutationResolver) JoinStream(ctx context.Context, accountName string, rtcAnswer string) (string, error) {
	rtcConnection := r.RTCConnections[accountName]

	r.Mutex.Lock()
	r.RTCJoinSockets[accountName] <- rtcAnswer
	r.Mutex.Unlock()

	return rtcConnection, nil
}

func (r *mutationResolver) NewIceCandidate(ctx context.Context, accountName string, candidate string) (string, error) {
	r.Mutex.Lock()
	for _, e := range r.StreamingRooms[accountName] {
		e <- candidate
	}
	r.Mutex.Unlock()

	return candidate, nil
}

func (r *queryResolver) Streams(ctx context.Context) ([]string, error) {
	var streamingUsers []string
	for accountName := range r.RTCConnections {
		streamingUsers = append(streamingUsers, accountName)
	}

	return streamingUsers, nil
}

func (r *subscriptionResolver) OnStreamJoin(ctx context.Context) (<-chan string, error) {
	user, err := middlewares.UseAuth(ctx)
	if err != nil {
		return nil, err
	}

	events := make(chan string, 1)

	r.Mutex.Lock()
	r.RTCJoinSockets[user.AccountName] = events
	r.Mutex.Unlock()

	go func() {
		<-ctx.Done()
		r.Mutex.Lock()
		delete(r.RTCJoinSockets, user.AccountName)
		r.Mutex.Unlock()
	}()

	return events, nil
}

func (r *subscriptionResolver) OnNewIceCandidate(ctx context.Context, accountName string) (<-chan string, error) {
	events := make(chan string, 1)
	r.Mutex.Lock()
	r.StreamingRooms[accountName] = append(r.StreamingRooms[accountName], events)
	r.Mutex.Unlock()

	go func() {
		<-ctx.Done()

		r.Mutex.Lock()
		var channels []chan string
		for _, channel := range r.StreamingRooms[accountName] {
			if channel != events {
				channels = append(channels, channel)
			}
		}

		r.StreamingRooms[accountName] = channels
		r.Mutex.Unlock()
	}()

	return events, nil
}

func (r *userResolver) Stream(ctx context.Context, obj *models.User) (string, error) {
	return r.RTCConnections[obj.AccountName], nil
}
