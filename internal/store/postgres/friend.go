package postgres

import (
	"context"
	"errors"
	"log/slog"
)

func (s *Store) CreateFriendRequest(ctx context.Context, sendingUserId int, receivingUserId int) error {
	sqlQuery := `
		INSERT INTO user_friends (user_id_fk, friend_id_fk, status)
		VALUES ($1, $2, 'pending')
	`

	commandTag, err := s.conn.Exec(ctx, sqlQuery, sendingUserId, receivingUserId)
	if err != nil {
		slog.Error("failed create friend request", "error", err)
		return err
	}
	if commandTag.RowsAffected() != 1 {
		slog.Error("failed create friend request")
		return errors.New("failed create friend request")
	}
	return nil
}
