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

	r.RTCConnections[user.AccountName] = rtcConnection
	return rtcConnection, nil
}

func (r *mutationResolver) StopStreaming(ctx context.Context) (bool, error) {
	user, err := middlewares.UseAuth(ctx)
	if err != nil {
		return false, err
	}

	delete(r.RTCConnections, user.AccountName)
	return true, nil
}

func (r *mutationResolver) JoinStream(ctx context.Context, accountName string, rtcAnswer string) (string, error) {
	rtcConnection := r.RTCConnections[accountName]

	r.RTCJoinSockets[accountName] <- rtcAnswer

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

	r.RTCJoinSockets[user.AccountName] = events

	go func() {
		<-ctx.Done()
		delete(r.RTCJoinSockets, user.AccountName)
	}()

	return events, nil
}

func (r *subscriptionResolver) OnNewIceCandidate(ctx context.Context, accountName string) (<-chan string, error) {
	events := make(chan string, 1)
	r.Mutex.Lock()
	r.StreamingRooms[accountName] = append(r.StreamingRooms[accountName], events)
	r.Mutex.Unlock()
	return events, nil
}

func (r *userResolver) Stream(ctx context.Context, obj *models.User) (string, error) {
	return r.RTCConnections[obj.AccountName], nil
}
