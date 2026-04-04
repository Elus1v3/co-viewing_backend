package service

import (
	"co-viewing/internal/models"
	"co-viewing/internal/store/postgres"
	"context"
)

type UserService struct {
	store *postgres.UserStore
}

func NewUserService(store *postgres.UserStore) *UserService {
	return &UserService{store: store}
}

func (s *UserService) Create(ctx context.Context, user models.User) (int, error) {
	return s.store.Create(ctx, user)
}
