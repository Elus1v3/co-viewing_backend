package service

import (
	"co-viewing/internal/models"
	"context"
)

func (s *Service) CreateFriendRequest(ctx context.Context, sendingUserId int, receivingUserId int) error {
	return s.store.CreateFriendRequest(ctx, sendingUserId, receivingUserId)
}

func (s *Service) GetFriendRequestsFromId(ctx context.Context, receivingUserId int) ([]models.User, error) {
	return s.store.GetFriendRequestsFromId(ctx, receivingUserId)
}

func (s *Service) UpdateFriendRequest(ctx context.Context, request models.FriendRequest) error {
	return s.store.UpdateFriendRequest(ctx, request)
}
