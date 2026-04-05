package postgres

import (
	"co-viewing/internal/models"
	"context"
	"log/slog"

	"github.com/jackc/pgx/v5/pgxpool"
)

type UserStore struct {
	conn *pgxpool.Pool
}

func NewUserStore(conn *pgxpool.Pool) *UserStore {
	return &UserStore{conn: conn}
}

func (s *UserStore) Create(ctx context.Context, user models.User) (int, error) {
	sqlQuery := `
		INSERT INTO "User" (nickname, password)
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

func (s *UserStore) FindByNickname(ctx context.Context, nickname string) (bool, error) {
	sqlQuery := `
		SELECT EXISTS(SELECT 1 FROM "User" WHERE nickname = $1)
	`

	var exists bool
	err := s.conn.QueryRow(ctx, sqlQuery, nickname).Scan(&exists)
	if err != nil {
		slog.Error("find user failed", "error", err)
		return false, err
	}

	return exists, nil
}

func (s *UserStore) GetPassword(ctx context.Context, nickname string) (string, error) {
	sqlQuery := `
		SELECT password FROM "User" WHERE nickname = $1
	`

	var password string
	err := s.conn.QueryRow(ctx, sqlQuery, nickname).Scan(&password)
	if err != nil {
		slog.Error("failed get password", "error", err)
		return "", err
	}

	return password, nil
}
