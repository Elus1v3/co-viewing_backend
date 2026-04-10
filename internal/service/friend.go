package service

import "context"

func (s Service) CreateFriendRequest(ctx context.Context, sendingUserId int, receivingUserId int) error {
	return s.store.CreateFriendRequest(ctx, sendingUserId, receivingUserId)
}
