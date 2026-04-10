package postgres

import (
	"co-viewing/internal/models"
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

func (s *Store) GetFriendRequestsFromId(ctx context.Context, receivingUserId int) ([]models.User, error) {
	sqlQuery := `
		SELECT u.nickname, uf.user_id_fk
		FROM user_friends uf
		JOIN "user" u ON u.id_pk = uf.user_id_fk
		WHERE uf.friend_id_fk = $1
		AND uf.status = 'pending';
	`

	var sendingUsers []models.User
	rows, err := s.conn.Query(ctx, sqlQuery, receivingUserId)
	if err != nil {
		slog.Error("failed get friend requests", "error", err)
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var u models.User
		err := rows.Scan(&u.Nickname, &u.Id)
		if err != nil {
			slog.Error("failed to scan", "error", err)
		}
		sendingUsers = append(sendingUsers, u)
	}

	return sendingUsers, nil
}

func (s *Store) UpdateFriendRequest(ctx context.Context, request models.FriendRequest) error {
	sqlQuery := `
		UPDATE user_friends
		SET status = 'accepted'
		WHERE user_id_fk = $1 
		AND friend_id_fk = $2
	`

	commandTag, err := s.conn.Exec(ctx, sqlQuery, request.UserId, request.FriendId)
	if err != nil {
		slog.Error("failed update friend request", "error", err)
		return err
	}
	if commandTag.RowsAffected() != 1 {
		slog.Error("failed update friend request")
		return errors.New("failed update friend request")
	}

	return nil
}
