package service

import (
	"co-viewing/internal/models"
	"context"
	"errors"
	"log/slog"

	"golang.org/x/crypto/bcrypt"
)

func hashPassword(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword(
		[]byte(password),
		bcrypt.DefaultCost,
	)
	if err != nil {
		slog.Error("error hash password", "error", err)
		return "", err
	}
	return string(hash), nil
}

func (s *Service) Create(ctx context.Context, user models.User) (int, error) {
	exists, err := s.store.FindByNickname(ctx, user.Nickname)
	if err != nil {
		return 0, err
	}
	if exists == true {
		slog.Info("user with this nickname already exists", "nickname", user.Nickname)
		return 0, errors.New("user with this nickname already exists")
	}

	hashPassword, err := hashPassword(user.Password)
	if err != nil {
		return 0, err
	}
	user.Password = hashPassword

	return s.store.Create(ctx, user)
}

func (s *Service) SignIn(ctx context.Context, user models.User) (models.User, error) {
	exists, err := s.store.FindByNickname(ctx, user.Nickname)
	if err != nil {
		return models.User{}, err
	}
	if exists == false {
		slog.Info("user with this nickname not exists", "nickname", user.Nickname)
		return models.User{}, errors.New("user with this nickname not exists")
	}

	UserWithPassword, err := s.store.GetPassword(ctx, user.Nickname)
	if err != nil {
		return models.User{}, err
	}
	err = bcrypt.CompareHashAndPassword(
		[]byte(UserWithPassword.Password),
		[]byte(user.Password),
	)
	if err != nil {
		return models.User{}, errors.New("password incorrect")
	}

	slog.Info("user is sign in")
	return UserWithPassword, nil
}

func (s *Service) GetAllUsers(ctx context.Context) ([]models.User, error) {
	return s.store.GetAllUsers(ctx)
}
