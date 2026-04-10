package postgres

import (
	"co-viewing/internal/models"
	"context"
	"log/slog"
)

func (s *Store) Create(ctx context.Context, user models.User) (int, error) {
	sqlQuery := `
		INSERT INTO "user" (nickname, password)
		VALUES ($1, $2)
		RETURNING id_pk;
	`

	var id int
	err := s.conn.QueryRow(ctx, sqlQuery,
		user.Nickname,
		user.Password,
	).Scan(&id)

	if err != nil {
		slog.Error("create user failed", "error", err)
		return 0, nil
	}
	slog.Info("user created", "id", id)
	return id, nil
}

func (s *Store) FindByNickname(ctx context.Context, nickname string) (bool, error) {
	sqlQuery := `
		SELECT EXISTS(SELECT 1 FROM "user" WHERE nickname = $1)
	`

	var exists bool
	err := s.conn.QueryRow(ctx, sqlQuery, nickname).Scan(&exists)
	if err != nil {
		slog.Error("find user failed", "error", err)
		return false, err
	}

	return exists, nil
}

func (s *Store) GetPassword(ctx context.Context, nickname string) (models.User, error) {
	sqlQuery := `
		SELECT password, id_pk FROM "user" WHERE nickname = $1
	`

	var user models.User
	err := s.conn.QueryRow(ctx, sqlQuery, nickname).Scan(&user.Password, &user.Id)
	if err != nil {
		slog.Error("failed get password", "error", err)
		return models.User{}, err
	}

	return user, nil
}
