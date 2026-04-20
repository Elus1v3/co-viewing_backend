package postgres

import (
	"co-viewing/internal/models"
	"context"
	"log/slog"
)

func (s *Store) AddWatchedMovie(ctx context.Context, movie models.WatchedMovie) error {
	sqlQuery := `
		INSERT INTO watched_movies (user_id_fk, movie_id)
		VALUES ($1, $2)
		ON CONFLICT (user_id_fk, movie_id) DO NOTHING;
	`

	commandTag, err := s.conn.Exec(ctx, sqlQuery, movie.UserId, movie.MovieId)
	if err != nil {
		slog.Error("failed add wached movie", "error", err)
		return err
	}
	if commandTag.RowsAffected() == 1 {
		slog.Info("watched movie added")
	} else {
		slog.Info("watched movie already exists")
	}

	return nil
}

func (s *Store) GetWatchedMoviesById(ctx context.Context, userId int) ([]models.WatchedMovie, error) {
	sqlQuery := `
		SELECT user_id_fk, movie_id FROM watched_movies WHERE user_id_fk = $1; 
	`

	var watchedMovies []models.WatchedMovie
	rows, err := s.conn.Query(ctx, sqlQuery, userId)
	if err != nil {
		slog.Error("failed get watched movies", "error", err)
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var m models.WatchedMovie
		err := rows.Scan(&m.UserId, &m.MovieId)
		if err != nil {
			slog.Error("error", "failed scan watched movie", err)
			continue
		}
		watchedMovies = append(watchedMovies, m)
	}

	return watchedMovies, nil
}
