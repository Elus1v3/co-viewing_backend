package service

import (
	"co-viewing/internal/models"
	"co-viewing/internal/store/postgres"
	"context"
	"errors"
	"log/slog"
)

type UserService struct {
	store *postgres.UserStore
}

func NewUserService(store *postgres.UserStore) *UserService {
	return &UserService{store: store}
}

func (s *UserService) Create(ctx context.Context, user models.User) (int, error) {
	exists, err := s.store.FindByNickname(ctx, user.Nickname)
	if err != nil {
		return 0, err
	}
	if exists == true {
		slog.Info("user with this nickname already exists", "nickname", user.Nickname)
		return 0, errors.New("user with this nickname already exists")
	}

	return s.store.Create(ctx, user)
}

func (s *UserService) SignIn(ctx context.Context, user models.User) error {
	exists, err := s.store.FindByNickname(ctx, user.Nickname)
	if err != nil {
		return err
	}
	if exists == false {
		slog.Info("user with this nickname not exists", "nickname", user.Nickname)
		return errors.New("user with this nickname not exists")
	}

	password, err := s.store.GetPassword(ctx, user.Nickname)
	if err != nil {
		return err
	}
	if password != user.Password {
		slog.Info("password incorrect")
		return errors.New("password incorrect")
	}

	slog.Info("user is sign in")
	return nil
}
